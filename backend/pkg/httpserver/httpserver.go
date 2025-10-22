package httpserver

import (
	"bitnix-backend/config"
	"bitnix-backend/internal/application/services"
	pg_infra "bitnix-backend/internal/infrastructure/db/postgres"
	"bitnix-backend/internal/infrastructure/storage"
	"bitnix-backend/internal/interface/api/rest"
	"context"
	"database/sql"
	"log/slog"
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

	cloudClient, supaClient, err := storage.BuildStorageClients(cfg)
	if err != nil {
		return nil, err
	}

	storageClientIface := storage.NewMultiStorageClient(cloudClient, supaClient)

	app := echo.New()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app.Use(CustomLogger(logger))
	app.Use(middleware.Recover())

	app.GET("/swagger/*", echoSwagger.WrapHandler)

	gameRepo := pg_infra.NewPostgresGameRepository(db)
	assetRepo := pg_infra.NewPostgresAssetRepository(db)
	gameService := services.NewGameService(gameRepo, assetRepo, storageClientIface)
	rest.NewgameController(app, gameService)

	catalogRepo := pg_infra.NewPostgresCatalogRepository(db)
	catalogService := services.NewCatalogService(catalogRepo)
	rest.NewCatalogController(app, catalogService)

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
