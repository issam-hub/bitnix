package postgres

import (
	"bitnix-backend/internal/domain/entities"
	"context"
	"database/sql"
)

type PostgresCatalogRepository struct {
	DB *sql.DB
}

func NewPostgresCatalogRepository(db *sql.DB) *PostgresCatalogRepository {
	return &PostgresCatalogRepository{
		DB: db,
	}
}

func (pcr PostgresCatalogRepository) GetAll(ctx context.Context) ([]*entities.CatalogItem, error) {
	query := `SELECT g.id, g.title, g.price, a.url AS thumbnail FROM games g LEFT JOIN assets a ON g.id = a.game_id AND a.type = 'poster'`

	var dbCatalogItems []CatalogItem

	rows, err := pcr.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dbCatalogItem CatalogItem
		var nullyThumbnail sql.NullString
		err := rows.Scan(
			&dbCatalogItem.GameID,
			&dbCatalogItem.Title,
			&dbCatalogItem.Price,
			&nullyThumbnail,
		)
		if err != nil {
			return nil, err
		}
		dbCatalogItem.Thumbnail = nullyThumbnail.String
		dbCatalogItems = append(dbCatalogItems, dbCatalogItem)
	}

	var catalogItems []*entities.CatalogItem
	for _, item := range dbCatalogItems {
		catalogItems = append(catalogItems, fromDbCatalogItem(&item))
	}

	return catalogItems, nil
}
