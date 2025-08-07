package entities

import (
	"bitnix-backend/internal/validator"
	"reflect"
	"testing"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

func TestGame(t *testing.T) {
	cases := []struct {
		description string
		game        Game
		want        validator.ValidationErrors
	}{
		{
			description: "invalid game",
			game:        *NewGame("", "", *money.NewFromFloat(0.0, "USD"), User{}, time.Now(), []Asset{}),
			want: validator.ValidationErrors{
				{
					Field:   "title",
					Message: "title is required and cannot be empty",
				},
				{
					Field:   "description",
					Message: "description is required and cannot be empty",
				},
				{
					Field:   "price",
					Message: "price cannot be zero or negative",
				},
				{
					Field:   "developer",
					Message: "developer is required",
				},
				{
					Field:   "assets",
					Message: "at least one asset is required",
				},
			},
		},
		{
			description: "valid game",
			game: *NewGame("valid title",
				"valid description",
				*money.NewFromFloat(15.0, "USD"),
				User{
					ID:    uuid.New(),
					Name:  "valid developer",
					Email: "valid.developer@gmail.com",
					Role:  "both",
				},
				time.Now(),
				[]Asset{
					{
						ID:       uuid.New(),
						Type:     DownloadFile,
						GameID:   uuid.New(),
						URL:      "https://downloadMe.com",
						Filename: "downloadFile",
					},
				}),
			want: validator.ValidationErrors(nil),
		},
		{
			description: "valid game with one invalid asset",
			game: *NewGame("valid title",
				"valid description",
				*money.NewFromFloat(15.0, "USD"),
				User{
					ID:    uuid.New(),
					Name:  "valid developer",
					Email: "valid.developer@gmail.com",
					Role:  "developer",
				},
				time.Now(),
				[]Asset{
					{
						ID:       uuid.New(),
						Type:     "",
						GameID:   uuid.Nil,
						URL:      "",
						Filename: "",
					},
				}),
			want: validator.ValidationErrors{
				{
					Field:   "assets",
					Message: "invalid assets: #1",
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			got := tc.game.Validate()
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("game validation error, got %#v, want %#v", got, tc.want)
			}
		})
	}
}
