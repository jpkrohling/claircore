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
	selectDistsByArtifactJoin = `SELECT
	dist.id,
	dist.name,
	dist.did,
	dist.version,
	dist.version_code_name,
	dist.version_id,
	dist.arch,
	dist.cpe,
	dist.pretty_name
FROM
	dist_scanartifact
	LEFT JOIN dist ON dist_scanartifact.dist_id = dist.id
WHERE
	dist_scanartifact.layer_hash = '%s' AND dist_scanartifact.scanner_id IN (?);`
)

func distributionsByLayer(ctx context.Context, db *sqlx.DB, hash string, scnrs scanner.VersionedScanners) ([]*claircore.Distribution, error) {
	// TODO Use passed-in Context.
	// get scanner ids
	scannerIDs := []int{}
	for _, scnr := range scnrs {
		var scannerID int
		err := db.Get(&scannerID, scannerIDByNameVersionKind, scnr.Name(), scnr.Version(), scnr.Kind())
		if err != nil {
			return nil, fmt.Errorf("store:distributionseByLayer failed to retrieve scanner ids for scnr %v: %v", scnr, err)
		}
		scannerIDs = append(scannerIDs, scannerID)
	}

	res := []*claircore.Distribution{}

	// rebind see: https://jmoiron.github.io/sqlx/ "in queries" section
	// we need to format this query since an IN query can only have one bindvar. TODO: confirm this
	withHash := fmt.Sprintf(selectDistsByArtifactJoin, hash)
	inQuery, args, err := sqlx.In(withHash, scannerIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to bind scannerIDs to query: %v", err)
	}
	inQuery = db.Rebind(inQuery)

	rows, err := db.Queryx(inQuery, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("store:distributionsByLayer no distribution found for hash %v and scnrs %v", hash, scnrs)
		}
		return nil, fmt.Errorf("store:distributionsByLayer failed to retrieve package rows for hash %v and scanners %v: %v", hash, scnrs, err)
	}
	defer rows.Close()

	for rows.Next() {
		var dist claircore.Distribution

		err := rows.Scan(
			&dist.ID,
			&dist.Name,
			&dist.DID,
			&dist.Version,
			&dist.VersionCodeName,
			&dist.VersionID,
			&dist.Arch,
			&dist.CPE,
			&dist.PrettyName,
		)
		if err != nil {
			return nil, fmt.Errorf("store:distributionsByLayer failed to scan distribution: %v", err)
		}

		res = append(res, &dist)
	}

	return res, nil
}
