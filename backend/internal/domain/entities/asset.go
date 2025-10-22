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
	Poster       AssetType = "poster"
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

func NewAsset(fileType AssetType, url string, filename string) *Asset {
	return &Asset{
		ID:       uuid.New(),
		Type:     fileType,
		URL:      url,
		Filename: filename,
	}
}

var urlRegex = regexp.MustCompile(`^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&//=]*)$`)

func ParseAssetType(assetType string) (AssetType, error) {
	switch strings.ToLower(strings.TrimSpace(assetType)) {
	case string(CoverImage):
		return CoverImage, nil
	case string(TrailerVideo):
		return TrailerVideo, nil
	case string(DownloadFile):
		return DownloadFile, nil
	case string(Screeshot):
		return Screeshot, nil
	default:
		return "", validator.ValidationErrors{validator.ValidationError{
			Field:   "type",
			Message: "invalid asset type",
		}}
	}
}
