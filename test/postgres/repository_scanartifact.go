package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/quay/claircore"
	"github.com/quay/claircore/internal/scanner"
)

func InsertRepoScanArtifact(db *sqlx.DB, layerHash string, repos []*claircore.Repository, scnrs scanner.VersionedScanners) error {
	n := len(scnrs)
	for i, repo := range repos {
		nn := i % n
		_, err := db.Exec(`INSERT INTO repo_scanartifact
			(layer_hash, repo_id, scanner_id)
		VALUES
			($1, $2, $3)`,
			&layerHash, &repo.ID, &nn)
		if err != nil {
			return fmt.Errorf("failed to insert repo scan artifact: %v", err)
		}
	}

	return nil
}
