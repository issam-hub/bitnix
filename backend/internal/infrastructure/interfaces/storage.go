package interfaces

import (
	"context"
	"io"
)

type StorageClient interface {
	// Upload stores a file in the storage
	Upload(ctx context.Context, path string, content io.Reader, contentType string) (string, error)

	// Download retrieves a file from the storage
	Download(ctx context.Context, path string) (io.ReadCloser, error)

	// Delete removes a file from the storage
	Delete(ctx context.Context, path string) error

	// Exists checks if a file exists in the storage
	Exists(ctx context.Context, path string) (bool, error)

	// GetURL returns a URL for accessing the file
	GetURL(ctx context.Context, path string, expiration int64) (string, error)
}
