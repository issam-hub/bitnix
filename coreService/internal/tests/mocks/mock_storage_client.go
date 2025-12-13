package mocks

import (
	"context"
	"io"

	"github.com/stretchr/testify/mock"
)

type MockStorageClient struct {
	mock.Mock
}

func (m *MockStorageClient) Upload(ctx context.Context, filename, path string, content io.Reader, contentType string) (string, error) {
	args := m.Called(ctx, filename, path, content, contentType)
	return args.String(0), args.Error(1)
}

func (m *MockStorageClient) Download(ctx context.Context, path string) (io.ReadCloser, error) {
	args := m.Called(ctx, path)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}

func (m *MockStorageClient) Delete(ctx context.Context, path string) error {
	args := m.Called(ctx, path)
	return args.Error(0)
}

func (m *MockStorageClient) Exists(ctx context.Context, path string) (bool, error) {
	args := m.Called(ctx, path)
	return args.Bool(0), args.Error(1)
}

func (m *MockStorageClient) GetURL(ctx context.Context, path string, expiration int64) (string, error) {
	args := m.Called(ctx, path, expiration)
	return args.String(0), args.Error(1)
}
