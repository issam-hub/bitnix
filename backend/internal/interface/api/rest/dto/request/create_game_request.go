package request

import (
	"bitnix-backend/internal/application/command"
	"bitnix-backend/internal/domain/entities"
	"bitnix-backend/internal/validator"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

var urlRegex = regexp.MustCompile(`^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&//=]*)$`)

type AssetRequest struct {
	Type     string `json:"type"`
	URL      string `json:"url"`
	Filename string `json:"filename"`
}

type CreateGameRequest struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	ReleaseDate string         `json:"release_date"`
	DeveloperID string         `json:"developer_id"`
	Assets      []AssetRequest `json:"assets"`
}

func (req *CreateGameRequest) Validate() validator.ValidationErrors {
	var errors validator.ValidationErrors
	if strings.TrimSpace(req.Title) == "" {
		errors = append(errors, validator.ValidationError{
			Field:   "title",
			Message: "title is required",
		})
	}

	if strings.TrimSpace(req.Description) == "" {
		errors = append(errors, validator.ValidationError{
			Field:   "description",
			Message: "description is required",
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

	// Validate each asset
	for i, asset := range req.Assets {
		if strings.TrimSpace(asset.Type) == "" {
			errors = append(errors, validator.ValidationError{
				Field:   fmt.Sprintf("assets[%d].type", i),
				Message: "asset type is required",
			})
		}

		if strings.TrimSpace(asset.URL) == "" {
			errors = append(errors, validator.ValidationError{
				Field:   fmt.Sprintf("assets[%d].url", i),
				Message: "URL is required",
			})
		} else if !urlRegex.MatchString(asset.URL) {
			errors = append(errors, validator.ValidationError{
				Field:   fmt.Sprintf("assets[%d].url", i),
				Message: "URL format is invalid",
			})
		}

		if strings.TrimSpace(asset.Filename) == "" {
			errors = append(errors, validator.ValidationError{
				Field:   fmt.Sprintf("assets[%d].filename", i),
				Message: "filename is required",
			})
		}
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

	var assetDetails []command.AssetDetail
	for _, asset := range req.Assets {
		assetType, err := entities.ParseAssetType(asset.Type)
		if err != nil {
			return nil, append(validationErrors, validator.ValidationError{
				Field:   "assets.type",
				Message: err.Error(),
			})
		}

		assetDetails = append(assetDetails, command.AssetDetail{
			Type:     assetType,
			URL:      asset.URL,
			Filename: asset.Filename,
		})
	}

	return &command.CreateGameCommand{
		Title:       req.Title,
		Description: req.Description,
		Price:       *money.NewFromFloat(req.Price, "USD"),
		DeveloperID: devID,
		ReleaseDate: releaseDate,
		Assets:      assetDetails,
	}, nil
}
