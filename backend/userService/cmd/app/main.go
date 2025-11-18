package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
	"user-service/config"

	_ "user-service/docs"
	"user-service/pkg/httpserver"
	"user-service/pkg/mongo"
)

// @title           user service REST API
// @version         1.0
// @description     user accounts & reviews service REST API documentation
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:7070
// @BasePath  /api/v1
// @schemes   http

func main() {
	config := config.NewConfig()

	dbClient, err := mongo.OpenDB(*config)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer func() {
		if err := dbClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	httpServer, err := httpserver.NewHTTPServer(dbClient, config)
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
