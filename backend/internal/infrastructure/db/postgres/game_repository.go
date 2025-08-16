package postgres

import (
	"bitnix-backend/internal/domain/apperrors"
	"bitnix-backend/internal/domain/entities"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type PostgresGameRepository struct {
	DB *sql.DB
}

func (pgr PostgresGameRepository) Create(ctx context.Context, game entities.Game) error {
	query := `
	INSERT INTO games
	(id, title, description, price, developer_id, release_date)
	VALUES
	($1, $2, $3, $4, $5, $6)
	RETURNING created_at, version
	`

	dbGame := toDBGame(&game)

	args := []any{
		dbGame.ID,
		dbGame.Title,
		dbGame.Description,
		dbGame.Price,
		dbGame.DeveloperID,
		dbGame.ReleaseDate,
	}

	err := pgr.DB.QueryRowContext(ctx, query, args...).Scan(&dbGame.CreatedAt, &dbGame.Version)
	if err != nil {
		return err
	}
	fmt.Println(dbGame.CreatedAt)
	return nil
}

func (fmg PostgresGameRepository) Get(ctx context.Context, id uuid.UUID) (*entities.Game, error) {
	query := `SELECT id, title, description, price, developer_id, release_date, created_at FROM games WHERE id = $1`

	var dbGame Game

	err := fmg.DB.QueryRowContext(ctx, query, id).Scan(
		&dbGame.ID,
		&dbGame.Title,
		&dbGame.Description,
		&dbGame.Price,
		&dbGame.DeveloperID,
		&dbGame.ReleaseDate,
		&dbGame.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, apperrors.ErrGameNotFound
		default:
			return nil, err
		}
	}

	game := fromDBGame(&dbGame)

	return game, nil
}
