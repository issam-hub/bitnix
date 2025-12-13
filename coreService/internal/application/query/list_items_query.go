package query

import "bitnix-backend/internal/application/common"

type ListItemsQueryResult struct {
	Result []*common.CatalogResult
}
