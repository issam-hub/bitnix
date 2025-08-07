package httpserver

import (
	"bitnix-backend/config"
	"bitnix-backend/internal/application/services"
	pg_infra "bitnix-backend/internal/infrastructure/db/postgres"
	"bitnix-backend/internal/interface/api/rest"
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Server struct {
	app    *echo.Echo
	config *config.Config
}

func NewHTTPServer(db *sql.DB, cfg *config.Config) (*Server, error) {
	gameRepo := pg_infra.PostgresGameRepository{DB: db}
	assetRepo := pg_infra.PostgresAssetRepository{DB: db}

	gameService := services.NewGameService(&gameRepo, &assetRepo)

	app := echo.New()
	app.HTTPErrorHandler = customHTTPErrorHandler

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app.Use(CustomLogger(logger))
	app.Use(middleware.Recover())

	app.GET("/swagger/*", echoSwagger.WrapHandler)

	rest.NewgameController(app, gameService)

	return &Server{
		app:    app,
		config: cfg,
	}, nil
}

func (s *Server) Start() error {
	return s.app.Start(":" + s.config.HTTP.Port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.Shutdown(ctx)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	var status int
	var message interface{}

	switch e := err.(type) {
	case *echo.HTTPError:
		status = e.Code
		message = e.Message
	default:
		status = http.StatusInternalServerError
		message = "the server encountered a problem and could not process your request"
	}

	if !c.Response().Committed {
		c.JSON(status, echo.Map{"error": message})
	}
}
