package httpserver

import (
	"context"
	"log/slog"
	"os"
	"time"
	"user-service/config"
	"user-service/internal/application/services"
	mongodb "user-service/internal/infrastructure/db"
	"user-service/internal/infrastructure/events/kafka/publisher"
	"user-service/internal/interface/api/rest"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	app    *echo.Echo
	config *config.Config
}

func StartHTTPServer(db *mongo.Client, cfg *config.Config) (*Server, error) {

	app := echo.New()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app.Use(CustomLogger(logger))
	app.Use(middleware.Recover())

	app.GET("/swagger/*", echoSwagger.WrapHandler)

	accountsRepo := mongodb.NewMongoUserRepository(db)

	accountsService := services.NewAccountsService(accountsRepo)

	rest.NewUserController(app, accountsService)

	publisher.StartPoller(db.Database("user-service"), cfg.Kafka.Broker, cfg.Kafka.Topic, 5*time.Second)

	return &Server{}, nil
}

func (s *Server) Start() error {
	return s.app.Start(":" + s.config.HTTP.Port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.Shutdown(ctx)
}
