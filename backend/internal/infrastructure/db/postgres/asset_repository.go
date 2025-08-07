package postgres

import (
	"bitnix-backend/internal/domain/entities"
	"context"
	"database/sql"
)

type PostgresAssetRepository struct {
	DB *sql.DB
}

func (par PostgresAssetRepository) CreateAll(ctx context.Context, assets []entities.Asset) error {
	query := `INSERT INTO assets
	(id, game_id, type, url, filename)
	VALUES 
	($1, $2, $3, $4, $5)
	RETURNING version`

	tx, err := par.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	for _, asset := range assets {
		dbAsset := toDBAsset(&asset)
		args := []any{
			dbAsset.ID,
			dbAsset.GameID,
			dbAsset.Type,
			dbAsset.URL,
			dbAsset.Filename,
		}
		err := tx.QueryRowContext(ctx, query, args...).Scan(&dbAsset.Version)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
