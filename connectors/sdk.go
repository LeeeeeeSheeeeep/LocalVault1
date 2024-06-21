package connectors

import (
	"context"
	"time"
	"localvault/storage"
)

// SyncState represents the state of a sync operation
type SyncState struct {
	LastSync time.Time
	Cursor   string
}

// Connector is the interface that all data sources must implement
type Connector interface {
	// ID returns a unique identifier for this connector (e.g., "github", "notion")
	ID() string
	
	// Name returns the human-readable name of the connector
	Name() string

	// Auth performs authentication (OAuth, PAT, etc.)
	Auth(ctx context.Context, config map[string]string) error

	// Sync pulls data incrementally or fully and sends it to the docChan
	Sync(ctx context.Context, state SyncState, docChan chan<- *storage.Document) (SyncState, error)
}
