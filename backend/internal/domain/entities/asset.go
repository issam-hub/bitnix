package entities

import (
	"bitnix-backend/internal/validator"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type AssetType string

const (
	CoverImage   AssetType = "cover"
	TrailerVideo AssetType = "trailer"
	DownloadFile AssetType = "download"
	Screeshot    AssetType = "screenshot"
)

type Asset struct {
	ID       uuid.UUID
	GameID   uuid.UUID
	Type     AssetType
	URL      string
	Filename string
}

var urlRegex = regexp.MustCompile(`^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&//=]*)$`)

func (a Asset) Validate() validator.ValidationErrors {
	var errors validator.ValidationErrors

	if strings.TrimSpace(a.Filename) == "" {
		errors = append(errors, validator.ValidationError{
			Field:   "filename",
			Message: "filename is required",
		})
	}

	if a.GameID == uuid.Nil {
		errors = append(errors, validator.ValidationError{
			Field:   "gameID",
			Message: "game ID is required",
		})
	}

	if a.Type == "" {
		errors = append(errors, validator.ValidationError{
			Field:   "type",
			Message: "asset type is required",
		})
	}

	if strings.TrimSpace(a.URL) == "" {
		errors = append(errors, validator.ValidationError{
			Field:   "url",
			Message: "URL is required",
		})
	} else if !urlRegex.MatchString(a.URL) {
		errors = append(errors, validator.ValidationError{
			Field:   "url",
			Message: "URL format is invalid",
		})
	}

	return errors
}
