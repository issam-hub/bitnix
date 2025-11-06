package httpserver

import (
	"context"
	"log/slog"
	"os"
	"user-service/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	app    *echo.Echo
	config *config.Config
}

func NewHTTPServer(db *mongo.Client, cfg *config.Config) (*Server, error) {

	app := echo.New()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app.Use(CustomLogger(logger))
	app.Use(middleware.Recover())

	app.GET("/swagger/*", echoSwagger.WrapHandler)

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
