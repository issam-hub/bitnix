package storage

import (
	"context"
	"io"
	"strings"

	infraif "bitnix-backend/internal/infrastructure/interfaces"
)

type MultiStorageClient struct {
	cloud infraif.StorageClient
	supa  infraif.StorageClient
}

func NewMultiStorageClient(cloud, supa infraif.StorageClient) infraif.StorageClient {
	return &MultiStorageClient{cloud: cloud, supa: supa}
}

func (m *MultiStorageClient) Upload(ctx context.Context, filename, path string, content io.Reader, contentType string) (string, error) {
	if strings.HasPrefix(contentType, "image/") || strings.HasPrefix(contentType, "video/") {
		if m.cloud != nil {
			return m.cloud.Upload(ctx, filename, path, content, contentType)
		}
	}
	if m.supa != nil {
		return m.supa.Upload(ctx, filename, path, content, contentType)
	}
	return "", ErrNoStorageProviders
}

func (m *MultiStorageClient) Download(ctx context.Context, path string) (io.ReadCloser, error) {
	return nil, nil
}

func (m *MultiStorageClient) Delete(ctx context.Context, path string) error {
	return nil
}

func (m *MultiStorageClient) Exists(ctx context.Context, path string) (bool, error) {
	return false, nil
}

func (m *MultiStorageClient) GetURL(ctx context.Context, path string, expiration int64) (string, error) {
	return "", nil
}
