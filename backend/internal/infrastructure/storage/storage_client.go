package storage

import (
	"context"
	"errors"
	"io"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type StorageClient struct {
	CLD *cloudinary.Cloudinary
}

func NewStorageClient(cld *cloudinary.Cloudinary) *StorageClient {
	return &StorageClient{
		CLD: cld,
	}
}

func (sc *StorageClient) Upload(ctx context.Context, filename, path string, content io.Reader, contentType string) (string, error) {
	res, err := sc.CLD.Upload.Upload(ctx, content, uploader.UploadParams{
		AssetFolder:  path,
		DisplayName:  filename,
		ResourceType: contentType,
	})
	if err != nil {
		return "", err
	}

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}

	return res.SecureURL, nil
}

func (sc *StorageClient) Download(ctx context.Context, path string) (io.ReadCloser, error) {
	return nil, nil
}

func (sc *StorageClient) Delete(ctx context.Context, path string) error {
	return nil
}

func (sc *StorageClient) Exists(ctx context.Context, path string) (bool, error) {
	return false, nil
}

func (sc *StorageClient) GetURL(ctx context.Context, path string, expiration int64) (string, error) {
	return "", nil
}
