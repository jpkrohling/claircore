package scanner

import (
	"context"

	"github.com/quay/claircore"
)

// Store is an interface for dealing with objects libscan needs to persist.
// Stores may be implemented per storage backend.
type Store interface {
	// ManifestScanned returns whether the given manifest was scanned by the provided scanners
	ManifestScanned(ctx context.Context, hash string, scnrs VersionedScanners) (bool, error)
	// LayerScanned returns whether the given layer was scanned by the provided scanner.
	LayerScanned(ctx context.Context, hash string, scnr VersionedScanner) (bool, error)
	// IndexPackages indexes a package into the persistence layer.
	IndexPackages(ctx context.Context, pkgs []*claircore.Package, layer *claircore.Layer, scnr VersionedScanner) error
	// IndexDistributions indexes distributions into the persistence layer
	IndexDistributions(ctx context.Context, dists []*claircore.Distribution, layer *claircore.Layer, scnr VersionedScanner) error
	// IndexRepositories indexes repositories into the persistence layer
	IndexRepositories(ctx context.Context, repos []*claircore.Repository, layer *claircore.Layer, scnr VersionedScanner) error
	// PackagesByLayer gets all the packages found in a layer limited by the provided scanners
	PackagesByLayer(ctx context.Context, hash string, scnrs VersionedScanners) ([]*claircore.Package, error)
	// DistributionsByLayer gets all the distributions found in a layer limited by the provided scanners
	DistributionsByLayer(ctx context.Context, hash string, scnrs VersionedScanners) ([]*claircore.Distribution, error)
	// RepositoriesByLayer gets all the repositories found in a layer limited by the provided scanners
	RepositoriesByLayer(ctx context.Context, hash string, scnrs VersionedScanners) ([]*claircore.Repository, error)
	// RegisterPackageScanners registers the provided scanners with the persistence layer
	RegisterScanners(ctx context.Context, scnrs VersionedScanners) error
	// ScanReport attempts to retrieve a persisted ScanReport.
	ScanReport(ctx context.Context, hash string) (*claircore.ScanReport, bool, error)
	// SetScanReport persists the current state of the ScanReport. ScanReports may
	// be in intermediate states to provide feedback for clients. this method should be
	// used to communicate scanning state updates. to signal the scan has completely successfully
	// see SetScanFinished
	SetScanReport(context.Context, *claircore.ScanReport) error
	// SetScanFinished marks a scan successfully completed. an association between
	// the provided manifest hash within the ScanReport and the list of VersionedScanners
	// should be made in such a way that ManifestScanned() correctly identifies if the manifest
	// was previously scanned by the given scnrs. the ScanResult should be pushed to the persistence
	// store.
	SetScanFinished(ctx context.Context, sr *claircore.ScanReport, scnrs VersionedScanners) error
}
