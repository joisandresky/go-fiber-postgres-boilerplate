package main

import (
	"os"

	"go-fiber-postgres-boilerplate/configs"
	"go-fiber-postgres-boilerplate/internal/infrastructure"
	"go-fiber-postgres-boilerplate/pkg/database"

	"go.uber.org/zap"
)

func main() {
	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync() // flushes buffer, if any

	logger := zapLogger.Sugar()

	cfg, err := configs.LoadConfig()
	if err != nil {
		logger.Errorf("failed to initialize configs", err)
		os.Exit(1)
	}

	// postgres db
	postgresDb := database.InitGormPostgres(cfg.App.Environment, &cfg.Database, logger)
	logger.Info("Successfully connect into Database!")

	appService := infrastructure.NewServer(cfg, logger, postgresDb)
	if err := appService.Run(); err != nil {
		logger.Errorf("failed to kickstart the service!: %v", err)
		os.Exit(1)
	}
}
