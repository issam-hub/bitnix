package main

import (
	"bitnix-backend/config"
	_ "bitnix-backend/docs"
	"bitnix-backend/pkg/httpserver"
	"bitnix-backend/pkg/postgres"
	"context"
	"log"
	"os"
	"os/signal"
	"time"
)

// @title           games catalog service REST API
// @version         1.0
// @description     games catalog service REST API documentation
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
// @schemes   http

func main() {
	config := config.NewConfig()

	db, err := postgres.OpenDB(*config)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer db.Close()

	httpServer, err := httpserver.StartHTTPServer(db, config)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := httpServer.Start(); err != nil {
			log.Fatalln("shutting down the server...")
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

}
