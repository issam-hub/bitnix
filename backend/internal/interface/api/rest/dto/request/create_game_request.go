package request

import (
	"bitnix-backend/internal/application/command"
	"bitnix-backend/internal/validator"
	"errors"
	"fmt"
	"strings"
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

func (req *CreateGameRequest) Validate() validator.ValidationErrors {
	var errors validator.ValidationErrors
	if strings.TrimSpace(req.Title) == "" {
		errors = append(errors, validator.ValidationError{
			Field:   "title",
			Message: "title is required and cannot be empty",
		})
	}

	if strings.TrimSpace(req.Description) == "" {
		errors = append(errors, validator.ValidationError{
			Field:   "description",
			Message: "description is required and cannot be empty",
		})
	}

	if req.Price <= 0 {
		errors = append(errors, validator.ValidationError{
			Field:   "price",
			Message: "price cannot be zero or negative",
		})
	}

	if req.DeveloperID == "" {
		errors = append(errors, validator.ValidationError{
			Field:   "developer_id",
			Message: "developer ID is required",
		})
	}

	if len(req.Assets) == 0 {
		errors = append(errors, validator.ValidationError{
			Field:   "assets",
			Message: "it should be at least one asset",
		})
	}
	return errors
}

func (req *CreateGameRequest) ToCreateGameCommand() (*command.CreateGameCommand, error) {
	validationErrors := req.Validate()
	if validationErrors.HasErrors() {
		return nil, validationErrors
	}
	devID, err := uuid.Parse(req.DeveloperID)
	if err != nil {
		return nil, append(validationErrors, validator.ValidationError{
			Field:   "developer_id",
			Message: "invalid developer ID format",
		})
	}

	releaseDate, err := time.Parse("2006-01-02", req.ReleaseDate)
	if err != nil {
		return nil, append(validationErrors, validator.ValidationError{
			Field:   "release_date",
			Message: "invalid release date format",
		})
	}

	var parsedAssetsIDs []uuid.UUID
	for i, assetID := range req.Assets {
		parsedAssetID, err := uuid.Parse(assetID)
		if err != nil {
			return nil, append(validationErrors, validator.ValidationError{
				Field:   "assets",
				Message: fmt.Sprintf("invalid asset ID #%d format", i+1),
			})
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
