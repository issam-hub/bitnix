package mapper

import (
	"bitnix-backend/internal/application/common"
	"bitnix-backend/internal/domain/entities"
)

func NewUploadAssetResultFromEntity(asset *entities.Asset) *common.AssetResult {
	return &common.AssetResult{
		ID:       asset.ID,
		Type:     asset.Type,
		URL:      asset.URL,
		Filename: asset.Filename,
	}
}
