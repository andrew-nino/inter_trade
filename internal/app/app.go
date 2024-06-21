package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	log "github.com/sirupsen/logrus"

	"international_trade/config"
	handler "international_trade/internal/controller/http/v1"
	repo "international_trade/internal/repo/pgdb"
	"international_trade/internal/service"

	httpserver "international_trade/pkg/httpserver"
	"international_trade/pkg/postgres"
	"international_trade/pkg/redis"
)

func Run(configPath string) {

	// Configuration
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logger
	SetLogrus(cfg.Log.Level)

	// Repositories
	log.Info("Initializing postgres...")
	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	// Migrates running
	log.Info("Migrates running...")
	m := NewMigration(cfg)
	m.Steps(1)

	// Starting Redis
	log.Info("Initializing Redis...")
	redis.ConnectRedis(cfg)

	// Services dependencies
	log.Info("Initializing services...")

	repos := repo.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	// HTTP server
	log.Info("Starting http server...")

	srv := new(httpserver.Server)

	go func() {
		if err := srv.Run(cfg.HTTP.Port, handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Print(cfg.App.Name + " Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print(cfg.App.Name + " Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Errorf("error occured on db connection close: %s", err.Error())
	}
}
