package rest

import (
	"bitnix-backend/internal/application/interfaces"
	"bitnix-backend/internal/domain/apperrors"
	"bitnix-backend/internal/interface/api/rest/dto/mapper"
	"bitnix-backend/internal/interface/api/rest/dto/request"
	"bitnix-backend/internal/validator"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type GameController struct {
	service interfaces.GameService
}

func NewgameController(e *echo.Echo, service interfaces.GameService) *GameController {
	controller := &GameController{
		service: service,
	}

	router := e.Group("/api/v1")

	router.POST("/game", controller.CreateGameController)
	router.GET("/game/:id", controller.GetGameController)
	router.POST("/assets/upload", controller.UploadAssetsController)

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
// @Failure 500 {object} resterror.ErrInternal
// @Router /game [post]
func (gc *GameController) CreateGameController(c echo.Context) error {
	var CreateGameRequest request.CreateGameRequest

	if err := c.Bind(&CreateGameRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	createGameCommand, err := CreateGameRequest.ToCreateGameCommand()
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			return echo.NewHTTPError(http.StatusBadRequest, errors.ToMap())
		}
		return echo.ErrInternalServerError
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := gc.service.CreateGame(ctx, createGameCommand)
	if err != nil {
		return echo.ErrInternalServerError
	}

	response := mapper.ToCreateGameResponse(result.Result)

	return c.JSON(http.StatusCreated, response)
}

// @Summary get a game
// @Description Get a game using its identifier
// @Tags Game Service
// @Accept json
// @Produce json
// @Param id path string true "Game ID (UUID)"
// @Success 200 {object} response.GetGameResponse
// @Failure 400 {object} resterror.ErrGetGameBadRequest
// @Failure 404 {object} resterror.ErrGameNotFoundResponse
// @Failure 500 {object} resterror.ErrInternal
// @Router /game/{id} [get]
func (gc *GameController) GetGameController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid game ID format")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := gc.service.GetGame(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrGameNotFound):
			return echo.NewHTTPError(http.StatusNotFound, apperrors.ErrGameNotFound.Error())
		default:
			return echo.ErrInternalServerError
		}
	}

	response := mapper.ToGetGameResponse(result.Result)

	return c.JSON(http.StatusOK, response)
}

// @Summary Upload assets
// @Description upload a set of assets to a claude-based storage client
// @Tags Game Service
// @Accept json
// @Produce json
// @Param request body request.UploadAssetsRequest true "Assets upload request"
// @Success 201 {object} response.UploadAssetsResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} resterror.ErrInternal
// @Router /assets/upload [post]
func (gc *GameController) UploadAssetsController(c echo.Context) error {
	var uploadAssetsRequest request.UploadAssetsRequest

	if err := c.Bind(&uploadAssetsRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	uploadAssetsCommand, err := uploadAssetsRequest.ToUploadAssetsCommand()
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			return echo.NewHTTPError(http.StatusBadRequest, errors.ToMap())
		}
		return echo.ErrInternalServerError
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := gc.service.UploadAssets(ctx, uploadAssetsCommand)
	if err != nil {
		fmt.Println("rah sra error ya kho: ", err.Error())
		return echo.ErrInternalServerError
	}

	response := mapper.ToUploadAssetsResponse(result.Result)

	return c.JSON(http.StatusCreated, response)
}
