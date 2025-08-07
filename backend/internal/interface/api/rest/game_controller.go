package rest

import (
	"bitnix-backend/internal/application/interfaces"
	"bitnix-backend/internal/interface/api/rest/dto/mapper"
	"bitnix-backend/internal/interface/api/rest/dto/request"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type GameController struct {
	service interfaces.GameService
}

func NewgameController(e *echo.Echo, service interfaces.GameService) *GameController {
	controller := &GameController{
		service: service,
	}

	e.POST("/api/v1/game", controller.CreateGameController)

	return controller
}

// @Summary Create a new game
// @Description Create a new game with the given details
// @Tags Game Service
// @Accept json
// @Produce json
// @Param request body request.CreateGameRequest true "Game creation request"
// @Success 201 {object} response.CreateGameResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/game [post]
func (gc *GameController) CreateGameController(c echo.Context) error {
	var CreateGameRequest request.CreateGameRequest

	if err := c.Bind(&CreateGameRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "failed to parse request body",
		})
	}

	createGameCommand, err := CreateGameRequest.ToCreateGameCommand()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := gc.service.CreateGame(ctx, createGameCommand)
	if err != nil {
		return err
	}

	response := mapper.ToCreateGameResponse(result.Result)

	return c.JSON(http.StatusCreated, response)
}
