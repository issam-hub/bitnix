package services

import (
	"bitnix-backend/internal/application/command"
	"bitnix-backend/internal/application/common"
	"bitnix-backend/internal/application/mapper"
	"bitnix-backend/internal/application/query"
	"bitnix-backend/internal/domain/entities"
	"bitnix-backend/internal/domain/repositories"
	"bitnix-backend/internal/infrastructure/interfaces"
	"context"

	"github.com/google/uuid"
)

type GameService struct {
	gameRepository  repositories.GameRepository
	assetRepository repositories.AssetRepository
	storageClient   interfaces.StorageClient
}

func NewGameService(
	gameRepository repositories.GameRepository,
	assetRepository repositories.AssetRepository,
	storageClient interfaces.StorageClient,
) GameService {
	return GameService{
		gameRepository:  gameRepository,
		assetRepository: assetRepository,
		storageClient:   storageClient,
	}
}

func (s GameService) CreateGame(ctx context.Context, gameCommand *command.CreateGameCommand) (*command.CreateGameCommandResult, error) {
	tempUser := entities.User{
		ID:    gameCommand.DeveloperID,
		Name:  "temp user",
		Email: "temp@example.com",
		Role:  "developer",
	}
	game := entities.NewGame(
		gameCommand.Title,
		gameCommand.Description,
		gameCommand.Price,
		tempUser,
		gameCommand.ReleaseDate,
		[]entities.Asset{},
	)

	var assets []entities.Asset
	for _, assetDetail := range gameCommand.Assets {
		assets = append(assets, entities.Asset{
			ID:       uuid.New(),
			Type:     assetDetail.Type,
			GameID:   game.ID,
			URL:      assetDetail.URL,
			Filename: assetDetail.Filename,
		})
	}
	game.Assets = assets

	if err := s.gameRepository.Create(ctx, *game); err != nil {
		return nil, err
	}

	if err := s.assetRepository.CreateAll(ctx, game.Assets); err != nil {
		return nil, err
	}

	commandResult := mapper.NewGameResultFromEntity(game)

	return &command.CreateGameCommandResult{
		Result: commandResult,
	}, nil
}

func (s GameService) GetGame(ctx context.Context, id uuid.UUID) (*query.GameQueryResult, error) {
	game, err := s.gameRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	assetPointers, err := s.assetRepository.GetAllByGame(ctx, id)
	if err != nil {
		return nil, err
	}

	assets := make([]entities.Asset, len(assetPointers))
	for i, asset := range assetPointers {
		assets[i] = *asset
	}
	game.Assets = assets

	queryResult := mapper.NewGameResultFromEntity(game)

	return &query.GameQueryResult{
		Result: queryResult,
	}, nil
}

func (s GameService) UploadAssets(ctx context.Context, assetsCommand *command.UploadAssetsCommand) (*command.UploadAssetsCommandResult, error) {
	var uploadResults []*common.AssetResult
	for _, asset := range assetsCommand.Assets {
		url, err := s.storageClient.Upload(ctx, "uploads/files/"+asset.Filename, asset.Content, asset.ContentType)
		if err != nil {
			return nil, err
		}

		entityAsset := entities.NewAsset(
			asset.Type,
			url,
			asset.Filename,
		)

		uploadResults = append(uploadResults, mapper.NewUploadAssetResultFromEntity(entityAsset))
	}

	return &command.UploadAssetsCommandResult{
		Result: uploadResults,
	}, nil
}
