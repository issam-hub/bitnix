package storage

import (
	"bitnix-backend/config"
	infraif "bitnix-backend/internal/infrastructure/interfaces"
)

func BuildStorageClients(cfg *config.Config) (cloud infraif.StorageClient, supa infraif.StorageClient, err error) {
	if cfg.StorageClient.CloudinaryURL != "" {
		cld, err := NewCloudinaryClient(cfg.StorageClient.CloudinaryURL)
		if err != nil {
			return nil, nil, err
		}
		cloud = cld
	}

	// Supabase
	if cfg.StorageClient.SupabaseProjectRef != "" && cfg.StorageClient.SupabaseAPIKey != "" && cfg.StorageClient.SupabaseBucket != "" {
		s, err := NewSupabaseClient(cfg.StorageClient.SupabaseProjectRef, cfg.StorageClient.SupabaseAPIKey, cfg.StorageClient.SupabaseBucket)
		if err != nil {
			return nil, nil, err
		}
		supa = s
	}

	return cloud, supa, nil
}
