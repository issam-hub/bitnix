package entities

import (
	"bitnix-backend/internal/validator"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type Game struct {
	ID          uuid.UUID
	Title       string
	Description string
	Price       money.Money
	Developer   User
	ReleaseDate time.Time
	Assets      []Asset
}

func NewGame(title string, description string, price money.Money, developer User, releaseDate time.Time, assets []Asset) *Game {
	return &Game{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Price:       price,
		Developer:   developer,
		ReleaseDate: releaseDate,
		Assets:      assets,
	}
}

func (g Game) Validate() validator.ValidationErrors {
	var errors validator.ValidationErrors

	if strings.TrimSpace(g.Title) == "" {
		errors = append(errors, validator.ValidationError{
			Field:   "title",
			Message: "title is required and cannot be empty",
		})
	}

	if strings.TrimSpace(g.Description) == "" {
		errors = append(errors, validator.ValidationError{
			Field:   "description",
			Message: "description is required and cannot be empty",
		})
	}

	if g.Price.IsZero() || g.Price.IsNegative() {
		errors = append(errors, validator.ValidationError{
			Field:   "price",
			Message: "price cannot be zero or negative",
		})
	}

	if reflect.DeepEqual(g.Developer, User{}) {
		errors = append(errors, validator.ValidationError{
			Field:   "developer",
			Message: "developer is required",
		})
	}

	if assetErrors := g.validateAssets(); len(assetErrors) > 0 {
		errors = append(errors, assetErrors...)
	}

	return errors
}

func (g Game) validateAssets() validator.ValidationErrors {
	var errors validator.ValidationErrors

	if len(g.Assets) == 0 {
		errors = append(errors, validator.ValidationError{
			Field:   "assets",
			Message: "at least one asset is required",
		})
		return errors
	}

	var invalidAssets []string
	for i, asset := range g.Assets {
		if assetErrors := asset.Validate(); assetErrors.HasErrors() {
			invalidAssets = append(invalidAssets, fmt.Sprintf("#%d", i+1))
		}
	}

	if len(invalidAssets) > 0 {
		errors = append(errors, validator.ValidationError{
			Field:   "assets",
			Message: fmt.Sprintf("invalid assets: %s", strings.Join(invalidAssets, ", ")),
		})
	}

	return errors
}
