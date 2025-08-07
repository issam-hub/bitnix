package postgres

import (
	"bitnix-backend/internal/domain/entities"
	"context"
	"database/sql"
	"fmt"
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
