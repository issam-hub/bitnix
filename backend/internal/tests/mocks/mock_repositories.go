package mocks

import (
	"bitnix-backend/internal/domain/apperrors"
	"bitnix-backend/internal/domain/entities"
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/google/uuid"
)

type InMemoryGameRepository struct {
	Listings []entities.Game
}

func (mg *InMemoryGameRepository) Create(ctx context.Context, game entities.Game) error {
	mg.Listings = append(mg.Listings, game)
	return nil
}

func (mg *InMemoryGameRepository) Get(ctx context.Context, id uuid.UUID) (*entities.Game, error) {
	for _, game := range mg.Listings {
		if game.ID == id {
			return &game, nil
		}
	}
	return nil, apperrors.ErrGameNotFound
}

type InMemoryAssetRepository struct {
	AssetsListings []entities.Asset
}

func (ma *InMemoryAssetRepository) CreateAll(ctx context.Context, assets []entities.Asset) error {
	ma.AssetsListings = slices.Concat(ma.AssetsListings, assets)
	return nil
}
func (ma *InMemoryAssetRepository) GetAllByGame(ctx context.Context, gameID uuid.UUID) ([]*entities.Asset, error) {
	var assets []*entities.Asset
	for _, asset := range ma.AssetsListings {
		if asset.GameID == gameID {
			assets = append(assets, &asset)
		}
	}
	if len(assets) > 0 {
		return assets, nil
	} else {
		return nil, errors.New("assets not found")
	}
}

type FailingInMemoryGameRepository struct{}

func (fmg *FailingInMemoryGameRepository) Create(ctx context.Context, game entities.Game) error {
	return fmt.Errorf("failed to create game")
}

func (fmg *FailingInMemoryGameRepository) Get(ctx context.Context, id uuid.UUID) (*entities.Game, error) {
	return nil, apperrors.ErrGameNotFound
}

type InMemoryCatalogRepository struct {
	Items []entities.CatalogItem
}

func (mc *InMemoryCatalogRepository) GetAll(ctx context.Context) ([]*entities.CatalogItem, error) {
	var result []*entities.CatalogItem
	for _, item := range mc.Items {
		result = append(result, &item)
	}
	return result, nil
}

type FailingInMemoryCatalogRepository struct{}

func (fmc *FailingInMemoryCatalogRepository) GetAll(ctx context.Context) ([]*entities.CatalogItem, error) {
	return nil, errors.New("error happening in catalog repository")
}
