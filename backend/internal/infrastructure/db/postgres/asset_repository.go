package postgres

import (
	"bitnix-backend/internal/domain/entities"
	"context"
	"database/sql"

	"github.com/google/uuid"
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

func (par PostgresAssetRepository) GetAllByGame(ctx context.Context, gameID uuid.UUID) ([]*entities.Asset, error) {
	query := `SELECT id, game_id, type, url, filename, created_at FROM assets WHERE game_id = $1`

	rows, err := par.DB.QueryContext(ctx, query, gameID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dbAssets []Asset
	for rows.Next() {
		var dbAsset Asset

		err := rows.Scan(
			&dbAsset.ID,
			&dbAsset.GameID,
			&dbAsset.Type,
			&dbAsset.URL,
			&dbAsset.Filename,
			&dbAsset.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		dbAssets = append(dbAssets, dbAsset)
	}

	var assets []*entities.Asset

	for _, dbAsset := range dbAssets {
		assets = append(assets, fromDBAsset(&dbAsset))
	}

	return assets, nil
}
