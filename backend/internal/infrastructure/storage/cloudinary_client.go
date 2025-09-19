package storage

import (
	"context"
	"errors"
	"io"
	"strings"

	infraif "bitnix-backend/internal/infrastructure/interfaces"
	storageclient "bitnix-backend/pkg/storage_client"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryClient struct {
	CLD *cloudinary.Cloudinary
}

func NewCloudinaryClient(cloudinaryURL string) (infraif.StorageClient, error) {
	cld, err := storageclient.NewCloudinaryInstance(cloudinaryURL)
	if err != nil {
		return nil, err
	}
	return &CloudinaryClient{CLD: cld}, nil
}

func (sc *CloudinaryClient) Upload(ctx context.Context, filename, path string, content io.Reader, contentType string) (string, error) {
	resourceType := "raw"
	if strings.HasPrefix(contentType, "image/") {
		resourceType = "image"
	} else if strings.HasPrefix(contentType, "video/") {
		resourceType = "video"
	}

	res, err := sc.CLD.Upload.Upload(ctx, content, uploader.UploadParams{
		AssetFolder:  path,
		DisplayName:  filename,
		ResourceType: resourceType,
	})
	if err != nil {
		return "", err
	}

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}

	return res.SecureURL, nil
}

func (sc *CloudinaryClient) Download(ctx context.Context, path string) (io.ReadCloser, error) {
	return nil, errors.New("cloudinary download not implemented")
}

func (sc *CloudinaryClient) Delete(ctx context.Context, path string) error {
	return errors.New("cloudinary delete not implemented")
}

func (sc *CloudinaryClient) Exists(ctx context.Context, path string) (bool, error) {
	return false, errors.New("cloudinary exists not implemented")
}

func (sc *CloudinaryClient) GetURL(ctx context.Context, path string, expiration int64) (string, error) {
	return "", errors.New("cloudinary getURL not implemented")
}
