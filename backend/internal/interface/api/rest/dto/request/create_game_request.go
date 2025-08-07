package request

import (
	"bitnix-backend/internal/application/command"
	"errors"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

var (
	ErrInvalidDevID       = errors.New("invalid developer ID format")
	ErrInvalidReleaseDate = errors.New("invalid release date format")
	ErrInvalidAssetID     = errors.New("invalid asset ID format")
)

type CreateGameRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	ReleaseDate string   `json:"release_date"`
	DeveloperID string   `json:"developer_id"`
	Assets      []string `json:"assets"`
}

func (req *CreateGameRequest) ToCreateGameCommand() (*command.CreateGameCommand, error) {
	devID, err := uuid.Parse(req.DeveloperID)
	if err != nil {
		return nil, ErrInvalidDevID
	}

	releaseDate, err := time.Parse("2006-01-02", req.ReleaseDate)
	if err != nil {
		return nil, ErrInvalidReleaseDate
	}

	var parsedAssetsIDs []uuid.UUID
	for _, assetID := range req.Assets {
		parsedAssetID, err := uuid.Parse(assetID)
		if err != nil {
			return nil, ErrInvalidAssetID
		}
		parsedAssetsIDs = append(parsedAssetsIDs, parsedAssetID)
	}

	return &command.CreateGameCommand{
		Title:       req.Title,
		Description: req.Description,
		Price:       *money.NewFromFloat(req.Price, "USD"),
		DeveloperID: devID,
		ReleaseDate: releaseDate,
		Assets:      parsedAssetsIDs,
	}, nil
}
