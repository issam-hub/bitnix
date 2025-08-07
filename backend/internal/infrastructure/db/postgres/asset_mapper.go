package postgres

import "bitnix-backend/internal/domain/entities"

func toDBAsset(asset *entities.Asset) *Asset {
	return &Asset{
		ID:       asset.ID,
		GameID:   asset.GameID,
		Type:     asset.Type,
		URL:      asset.URL,
		Filename: asset.Filename,
	}
}

func fromDBAsset(asset *Asset) *entities.Asset {
	return &entities.Asset{
		ID:       asset.ID,
		GameID:   asset.GameID,
		Type:     asset.Type,
		URL:      asset.URL,
		Filename: asset.Filename,
	}
}
