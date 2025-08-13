package services

import (
	"bitnix-backend/internal/application/command"
	"bitnix-backend/internal/application/mapper"
	"bitnix-backend/internal/domain/entities"
	"bitnix-backend/internal/domain/repositories"
	"context"

	"github.com/google/uuid"
)

type GameService struct {
	gameRepository  repositories.GameRepository
	assetRepository repositories.AssetRepository
}

func NewGameService(
	gameRepository repositories.GameRepository,
	assetRepository repositories.AssetRepository,
) GameService {
	return GameService{
		gameRepository:  gameRepository,
		assetRepository: assetRepository,
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
