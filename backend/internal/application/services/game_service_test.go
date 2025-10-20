package services

import (
	"bitnix-backend/internal/application/command"
	"bitnix-backend/internal/application/common"
	"bitnix-backend/internal/application/query"
	"bitnix-backend/internal/domain/apperrors"
	"bitnix-backend/internal/domain/entities"
	"bitnix-backend/internal/tests/mocks"
	"bytes"
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGameService(t *testing.T) {
	gameRepo := new(mocks.InMemoryGameRepository)
	assetRepo := new(mocks.InMemoryAssetRepository)
	storageClient := new(mocks.MockStorageClient)
	failingGameRepo := new(mocks.FailingInMemoryGameRepository)

	cmd := command.NewCreateGameCommand(
		"hollow knight",
		"hollow knight game",
		*money.NewFromFloat(20.99, "USD"),
		uuid.New(),
		time.Date(2025, 9, 2, 0, 0, 0, 0, time.UTC),
		[]command.AssetDetail{
			{
				Type:     entities.DownloadFile,
				URL:      "https://downloadMe.com",
				Filename: "downloadFile",
			},
		},
	)

	t.Run("create game - happy case", func(t *testing.T) {
		svc := NewGameService(gameRepo, assetRepo, storageClient)

		ctx := context.Background()

		if _, err := svc.CreateGame(ctx, cmd); err != nil {
			t.Errorf("error while creating the game: %v", err)
		}

		if len(gameRepo.Listings) == 0 {
			t.Errorf("error while creating the game: want %d, got %d", 1, len(gameRepo.Listings))
		}

		if len(assetRepo.AssetsListings) == 0 {
			t.Errorf("error while creating the game (assets didn't get created): want %d, got %d", 1, len(assetRepo.AssetsListings))
		}
	})

	t.Run("create game - sad case", func(t *testing.T) {
		svc := NewGameService(failingGameRepo, assetRepo, storageClient)

		ctx := context.Background()

		if _, err := svc.CreateGame(ctx, cmd); err == nil || err.Error() != "failed to create game" {
			t.Errorf("expected nil or 'failed to create game' error, got: %v", err)
		}

		if len(gameRepo.Listings) == 0 {
			t.Errorf("error while creating the game: want %d, got %d", 1, len(gameRepo.Listings))
		}

		if len(assetRepo.AssetsListings) == 0 {
			t.Errorf("error while creating the game (assets didn't get created): want %d, got %d", 1, len(assetRepo.AssetsListings))
		}
	})

	t.Run("get game - happy case", func(t *testing.T) {
		svc := NewGameService(gameRepo, assetRepo, storageClient)

		ctx := context.Background()

		result, err := svc.CreateGame(ctx, cmd)
		if err != nil {
			t.Errorf("error while creating the game: %v", err)
		}

		game, err := svc.GetGame(ctx, result.Result.ID)
		if err != nil {
			t.Errorf("error while getting the game: %v", err)
		}

		gameQueryResult := &query.GameQueryResult{
			Result: &common.GameResult{
				ID:          result.Result.ID,
				Title:       result.Result.Title,
				Description: result.Result.Description,
				Price:       result.Result.Price,
				DeveloperID: result.Result.DeveloperID,
				ReleaseDate: result.Result.ReleaseDate,
				Assets:      result.Result.Assets,
				CreatedAt:   result.Result.CreatedAt,
			},
		}

		if !reflect.DeepEqual(game, gameQueryResult) {
			t.Errorf("expected %#v, got %#v", gameQueryResult, game)
		}
	})

	t.Run("get game - sad case", func(t *testing.T) {
		svc := NewGameService(gameRepo, assetRepo, storageClient)

		ctx := context.Background()

		_, err := svc.GetGame(ctx, uuid.New())
		if !errors.Is(err, apperrors.ErrGameNotFound) {
			t.Errorf("expected error %v, got %v", apperrors.ErrGameNotFound, err)
		}
	})

	assetsCommand := command.NewUploadAssetsCommand([]command.AssetDetail{
		{
			Type:        entities.DownloadFile,
			Filename:    "download.exe",
			Content:     bytes.NewReader([]byte("fake download file")),
			ContentType: "application/octet-stream",
		},
		{
			Type:        entities.Screeshot,
			Filename:    "screenshot.png",
			Content:     bytes.NewReader([]byte("fake screenshot file")),
			ContentType: "image/png",
		},
	})

	t.Run("upload assets - happy case", func(t *testing.T) {
		svc := NewGameService(gameRepo, assetRepo, storageClient)

		ctx := context.Background()

		storageClient.On("Upload", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything, "application/octet-stream").Return("https://storage.example.com/games/download.exe", nil).Once()
		storageClient.On("Upload", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything, "image/png").Return("https://storage.example.com/games/screenshot.png", nil).Once()

		result, err := svc.UploadAssets(ctx, assetsCommand)

		assert.NoError(t, err)

		assert.Len(t, result.Result, 2)

		assert.Equal(t, entities.DownloadFile, result.Result[0].Type)
		assert.Equal(t, "download.exe", result.Result[0].Filename)
		assert.Equal(t, "https://storage.example.com/games/download.exe", result.Result[0].URL)

		assert.Equal(t, entities.Screeshot, result.Result[1].Type)
		assert.Equal(t, "screenshot.png", result.Result[1].Filename)
		assert.Equal(t, "https://storage.example.com/games/screenshot.png", result.Result[1].URL)

		storageClient.AssertExpectations(t)
	})

	t.Run("upload assets - sad case", func(t *testing.T) {
		svc := NewGameService(gameRepo, assetRepo, storageClient)

		ctx := context.Background()

		storageClient.On("Upload", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything, "application/octet-stream").Return("", errors.New("error while uploading download.exe")).Once()

		result, err := svc.UploadAssets(ctx, assetsCommand)

		assert.Error(t, err)

		assert.Equal(t, err.Error(), "error while uploading download.exe")

		assert.Nil(t, result)

		storageClient.AssertExpectations(t)
	})

}
