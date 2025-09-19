package storage

import (
	"context"
	"errors"
	"io"

	infraif "bitnix-backend/internal/infrastructure/interfaces"
	storageclient "bitnix-backend/pkg/storage_client"

	supabase_str "github.com/supabase-community/storage-go"
)

type SupabaseClient struct {
	client *supabase_str.Client
	bucket string
}

func NewSupabaseClient(projectRefID, apiKey, bucket string) (infraif.StorageClient, error) {
	cl := storageclient.NewSupaBaseInstance(projectRefID, apiKey)
	if cl == nil {
		return nil, errors.New("failed to initialize supabase client")
	}
	return &SupabaseClient{client: cl, bucket: bucket}, nil
}

func (s *SupabaseClient) Upload(ctx context.Context, filename, path string, content io.Reader, contentType string) (string, error) {
	opts := supabase_str.FileOptions{ContentType: &contentType, CacheControl: strPtr("3600"), Upsert: boolPtr(true)}
	_, err := s.client.UploadFile(s.bucket, path, content, opts)
	if err != nil {
		return "", err
	}
	pub := s.client.GetPublicUrl(s.bucket, path)
	return pub.SignedURL, nil
}

func (s *SupabaseClient) Download(ctx context.Context, path string) (io.ReadCloser, error) {
	return nil, nil
}

func (s *SupabaseClient) Delete(ctx context.Context, path string) error {
	return nil
}

func (s *SupabaseClient) Exists(ctx context.Context, path string) (bool, error) {
	return false, nil
}

func (s *SupabaseClient) GetURL(ctx context.Context, path string, expiration int64) (string, error) {
	return "", nil
}

func strPtr(s string) *string { return &s }
func boolPtr(b bool) *bool    { return &b }
