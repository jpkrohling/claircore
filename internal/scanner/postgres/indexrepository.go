package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jmoiron/sqlx"
	"github.com/quay/claircore"
	"github.com/quay/claircore/internal/scanner"
	"github.com/quay/claircore/pkg/microbatch"
)

const (
	insertRepository = `
    INSERT INTO repo 
		(name, key, uri)
	VALUES
		($1, $2, $3)
	ON CONFLICT (name, key, uri) DO NOTHING;
	`

	insertRepoScanArtifactWith = `
    WITH repositories AS (
		SELECT id AS repo_id FROM repo WHERE
			name = $1 AND
			key = $2 AND
			uri = $3
			),
		
		scanner AS (
		SELECT id AS scanner_id FROM scanner WHERE
		name = $4 AND version = $5 AND kind = $6
		)
	INSERT INTO repo_scanartifact (layer_hash, repo_id, scanner_id) VALUES 
			  ($7, 
			  (SELECT repo_id FROM repositories),
			  (SELECT scanner_id FROM scanner))
			  ON CONFLICT DO NOTHING;
	`
)

func indexRepositories(ctx context.Context, db *sqlx.DB, pool *pgxpool.Pool, repos []*claircore.Repository, l *claircore.Layer, scnr scanner.VersionedScanner) error {
	// obtain a transaction scopped batch
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("store:indexRepositories failed to create transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	insertRepoStmt, err := tx.Prepare(ctx, "insertRepoStmt", insertRepository)
	if err != nil {
		return fmt.Errorf("failed to create insert repo statement: %v", err)
	}
	insertRepoScanArtifactWithStmt, err := tx.Prepare(ctx, "insertRepoScanArtifactWith", insertRepoScanArtifactWith)
	if err != nil {
		return fmt.Errorf("failed to create insert repo scanartifact statement: %v", err)
	}
	if err != nil {
		return fmt.Errorf("failed to create statement: %v", err)
	}

	mBatcher := microbatch.NewInsert(tx, 500, time.Minute)
	for _, repo := range repos {
		err := mBatcher.Queue(
			ctx,
			insertRepoStmt.SQL,
			repo.Name,
			repo.Key,
			repo.URI,
		)
		if err != nil {
			return fmt.Errorf("batch insert failed for repo %v: %v", repo, err)
		}
	}
	err = mBatcher.Done(ctx)
	if err != nil {
		return fmt.Errorf("final batch insert failed for repo: %v", err)
	}

	// make repo scan artifacts
	mBatcher = microbatch.NewInsert(tx, 500, time.Minute)
	for _, repo := range repos {
		err := mBatcher.Queue(
			ctx,
			insertRepoScanArtifactWithStmt.SQL,
			repo.Name,
			repo.Key,
			repo.URI,
			scnr.Name(),
			scnr.Version(),
			scnr.Kind(),
			l.Hash,
		)
		if err != nil {
			return fmt.Errorf("batch insert failed for repo_scanartifact %v: %v", repo, err)
		}
	}
	err = mBatcher.Done(ctx)
	if err != nil {
		return fmt.Errorf("final batch insert failed for repo_scanartifact: %v", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("store:indexRepositories failed to commit tx: %v", err)
	}
	return nil
}
