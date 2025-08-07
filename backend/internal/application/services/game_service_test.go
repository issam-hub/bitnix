package services

import (
	"bitnix-backend/internal/application/command"
	"bitnix-backend/internal/domain/entities"
	"context"
	"slices"
	"testing"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type inMemoryGameRepository struct {
	listings []entities.Game
}

func (mg *inMemoryGameRepository) Create(ctx context.Context, game entities.Game) error {
	mg.listings = append(mg.listings, game)
	return nil
}

type inMemoryAssetRepository struct {
	assetsListings []entities.Asset
}

func (ma *inMemoryAssetRepository) CreateAll(ctx context.Context, assets []entities.Asset) error {
	ma.assetsListings = slices.Concat(ma.assetsListings, assets)
	return nil
}

func TestGameService(t *testing.T) {
	gameRepo := &inMemoryGameRepository{}
	assetRepo := &inMemoryAssetRepository{}
	svc := NewGameService(gameRepo, assetRepo)
	cmd := command.NewCreateGameCommand(
		"hollow knight",
		"hollow knight game",
		*money.NewFromFloat(20.99, "USD"),
		uuid.New(),
		time.Date(2025, 9, 2, 0, 0, 0, 0, time.UTC),
		[]uuid.UUID{
			uuid.New(),
		},
	)

	ctx := context.Background()

	if _, err := svc.CreateGame(ctx, cmd); err != nil {
		t.Errorf("error while creating the game: %v", err)
	}

	if len(gameRepo.listings) == 0 {
		t.Errorf("error while creating the game: want %d, got %d", 1, len(gameRepo.listings))
	}

	if len(assetRepo.assetsListings) == 0 {
		t.Errorf("error while creating the game (assets didn't get created): want %d, got %d", 1, len(assetRepo.assetsListings))
	}
}
