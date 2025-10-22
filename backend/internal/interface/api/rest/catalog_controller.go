package rest

import (
	"bitnix-backend/internal/application/interfaces"
	"bitnix-backend/internal/interface/api/rest/dto/mapper"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type CatalogController struct {
	service interfaces.CatalogService
}

func NewCatalogController(e *echo.Echo, service interfaces.CatalogService) *CatalogController {
	controller := &CatalogController{
		service: service,
	}

	router := e.Group("/api/v1")

	router.GET("/catalog", controller.ListItemsController)

	return controller
}

// @Summary List Items
// @Description List games in a catalog
// @Tags Catalog Service
// @Accept json
// @Produce json
// @Success 200 {object} response.ListItemsResponse
// @Failure 500 {object} resterror.ErrInternal
// @Router /catalog [get]
func (cc *CatalogController) ListItemsController(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := cc.service.ListItems(ctx)
	if err != nil {
		return echo.ErrInternalServerError
	}

	response := mapper.ToListItemsResponse(result.Result)

	return c.JSON(http.StatusOK, response)

}
