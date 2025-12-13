package response

import (
	"github.com/google/uuid"
)

type AssetResponse struct {
	ID       uuid.UUID `json:"id"`
	Type     string    `json:"type"`
	URL      string    `json:"url"`
	Filename string    `json:"filename"`
}

type GetGameResponse struct {
	ID          uuid.UUID       `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Price       float64         `json:"price"`
	ReleaseDate string          `json:"release_date"`
	DeveloperID string          `json:"developer_id"`
	Assets      []AssetResponse `json:"assets"`
}
