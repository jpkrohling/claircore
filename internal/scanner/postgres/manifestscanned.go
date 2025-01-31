package postgres

import (
	"context"
	"fmt"

	"github.com/quay/claircore/internal/scanner"

	"github.com/jmoiron/sqlx"
)

const (
	selectScannerIDs              = `SELECT id FROM scanner WHERE name = $1 AND version = $2 AND kind = $3`
	selectScannerIDsByScannerList = `SELECT scanner_id FROM scannerlist WHERE manifest_hash = $1`
)

// manifestScanned determines if a manifest has been scanned by ALL the provided
// scnrs.
func manifestScanned(ctx context.Context, db *sqlx.DB, hash string, scnrs scanner.VersionedScanners) (bool, error) {
	// TODO Use passed-in Context.
	// get the ids of the scanners we are testing for.
	var expectedIDs []int
	for _, scnr := range scnrs {
		var id int
		row := db.QueryRowx(selectScannerIDs, scnr.Name(), scnr.Version(), scnr.Kind())
		err := row.Scan(&id)
		if err != nil {
			return false, fmt.Errorf("store:manifestScanned failed to retrieve expected scanner id for scnr %v: %v", scnr, err)
		}
		expectedIDs = append(expectedIDs, id)
	}

	// get a map of the found ids which have scanned this package
	var temp = []int{}
	var foundIDs = map[int]struct{}{}
	err := db.Select(&temp, selectScannerIDsByScannerList, hash)
	if err != nil {
		return false, fmt.Errorf("store:manifestScanned failed to select scanner IDs for manifest: %v", err)
	}

	// if we are unable to find any scanner ids for this manifest hash, we have
	// never scanned this manifest.
	if len(temp) == 0 {
		return false, nil
	}

	// create foundIDs map from temporary array
	for _, id := range temp {
		foundIDs[id] = struct{}{}
	}

	// compare the expectedIDs array with our foundIDs. if we get a lookup
	// miss we can say the manifest has not been scanned by all the layers provided
	for _, id := range expectedIDs {
		if _, ok := foundIDs[id]; !ok {
			return false, nil
		}
	}

	return true, nil
}
