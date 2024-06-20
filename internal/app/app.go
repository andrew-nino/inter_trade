package app

import (
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

	// Services dependencies
	log.Info("Initializing services...")

	repos := repo.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	// HTTP server
	srv := new(httpserver.Server)

	go func() {
		if err := srv.Run(cfg.HTTP.Port, handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}

		log.Info("Starting http server...")
	}()

	log.Print(cfg.App.Name + " Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print(cfg.App.Name + " Shutting Down")

	//TODO Корректный выход

	//
	// log.Debugf("Server port: %s", cfg.HTTP.Port)
	// httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// // Waiting signal
	// log.Info("Configuring graceful shutdown...")
	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// select {
	// case s := <-interrupt:
	// 	log.Info("app - Run - signal: " + s.String())
	// case err = <-httpServer.Notify():
	// 	log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	// }

	// // Graceful shutdown
	// log.Info("Shutting down...")
	// err = httpServer.Shutdown()
	// if err != nil {
	// 	log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	// }
}
