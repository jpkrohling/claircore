package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/quay/claircore"
	"github.com/quay/claircore/internal/scanner"
)

const (
	selectReposByArtifactJoin = `SELECT
	repo.id,
	repo.name,
	repo.key,
	repo.uri
FROM
	repo_scanartifact
	LEFT JOIN repo ON repo_scanartifact.repo_id = repo.id
WHERE
	repo_scanartifact.layer_hash = '%s' AND repo_scanartifact.scanner_id IN (?);`
)

func repositoriesByLayer(ctx context.Context, db *sqlx.DB, hash string, scnrs scanner.VersionedScanners) ([]*claircore.Repository, error) {
	// TODO Use passed-in Context.
	// get scanner ids
	scannerIDs := []int{}
	for _, scnr := range scnrs {
		var scannerID int
		err := db.Get(&scannerID, scannerIDByNameVersionKind, scnr.Name(), scnr.Version(), scnr.Kind())
		if err != nil {
			return nil, fmt.Errorf("store:repositoriesByLayer failed to retrieve scanner ids for scnr %v: %v", scnr, err)
		}
		scannerIDs = append(scannerIDs, scannerID)
	}

	res := []*claircore.Repository{}

	// rebind see: https://jmoiron.github.io/sqlx/ "in queries" section
	// we need to format this query since an IN query can only have one bindvar. TODO: confirm this
	withHash := fmt.Sprintf(selectReposByArtifactJoin, hash)
	inQuery, args, err := sqlx.In(withHash, scannerIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to bind scannerIDs to query: %v", err)
	}
	inQuery = db.Rebind(inQuery)

	rows, err := db.Queryx(inQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("store:repositoriesByLayer no repositories found for hash %v and scnrs %v", hash, scnrs)
		}
		return nil, fmt.Errorf("store:repositoriesByLayer failed to retrieve package rows for hash %v and scanners %v: %v", hash, scnrs, err)
	}
	defer rows.Close()

	for rows.Next() {
		var repo claircore.Repository

		err := rows.Scan(
			&repo.ID,
			&repo.Name,
			&repo.Key,
			&repo.URI,
		)
		if err != nil {
			return nil, fmt.Errorf("store:repositoriesByLayer failed to scan repositories: %v", err)
		}

		res = append(res, &repo)
	}

	return res, nil
}
