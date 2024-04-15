package app

import (
	"awesomeProject/internal/config"
	"awesomeProject/internal/handler"
	"awesomeProject/internal/repository"
	"awesomeProject/internal/server"
	"awesomeProject/internal/service"
	"awesomeProject/migrations"
	"awesomeProject/pkg/logger"
	"database/sql"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"log"
)

func Run() {
	Logger, err := logger.SetupLogger()

	if err != nil {
		log.Fatal("cant initialize logger")
	}

	Logger.Info("initialized logger")

	cfg, err := config.SetupConfig()

	if err != nil {
		Logger.Fatal("cant initialize config", zap.Error(err))
	}

	db, err := sql.Open("postgres", cfg.Storage)

	if err != nil {
		Logger.Fatal("cant initialize db", zap.Error(err))
	}

	defer db.Close()

	migrations.RunMigrations(db)

	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{Repos: repos})
	handlers := handler.NewHandler(services)

	srv := server.NewServer(cfg, handlers.Init())

	err = srv.Start()

	if err != nil {
		Logger.Error(err.Error(), zap.Error(err))
	}
}
