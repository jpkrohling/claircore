package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/quay/claircore/internal/scanner"

	"github.com/jmoiron/sqlx"
)

const (
	insertScanner = `INSERT INTO scanner (name, version, kind) VALUES ($1, $2, $3) ON CONFLICT (name, version, kind) DO NOTHING;`
	selectScanner = `SELECT id FROM scanner WHERE name = $1 AND version = $2 AND kind = $3;`
)

func registerScanners(ctx context.Context, db *sqlx.DB, scnrs scanner.VersionedScanners) error {
	// TODO Use passed-in Context.
	// check if all scanners scanners exist
	ids := make([]sql.NullInt64, len(scnrs))
	for i, scnr := range scnrs {
		err := db.Get(&ids[i], selectScanner, scnr.Name(), scnr.Version(), scnr.Kind())
		if err != nil {
			fmt.Errorf("failed to get scanner id for scnr %v: %v", scnr, err)
		}
	}

	// register scanners not found
	for i, id := range ids {
		if !id.Valid {
			s := scnrs[i]
			_, err := db.Exec(insertScanner, s.Name(), s.Version(), s.Kind())
			if err != nil {
				return fmt.Errorf("failed to insert scanner %v: %v", s, err)
			}
		}
	}

	return nil
}
