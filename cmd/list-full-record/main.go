package main

import (
	"context"
	"fmt"
	"github.com/michael-bc/list-full-record-backend/internal/brand/brand_database"
	"github.com/michael-bc/list-full-record-backend/internal/brand/brand_domain"
	"github.com/michael-bc/list-full-record-backend/internal/brand/brand_http"
	"github.com/michael-bc/list-full-record-backend/internal/infrastructure/configs"
	"github.com/michael-bc/list-full-record-backend/internal/infrastructure/database/postgres"
	"github.com/michael-bc/list-full-record-backend/internal/infrastructure/http_server"
	"github.com/michael-bc/list-full-record-backend/internal/infrastructure/logger/logrus"
	"os"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		panic(fmt.Errorf("error on init configuration: %s", err))
	}

	log, err := logrus.NewLogger(cfg.LogLevel)
	if err != nil {
		panic(fmt.Errorf("error on init logger: %s", err))
	}

	fileVersion, err := os.ReadFile("VERSION")
	if err != nil {
		log.Panicf("error on read VERSION file: %s", err)
	}

	version := string(fileVersion)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbConn, err := postgres.NewPostgres(ctx, cfg.DbURI)
	if err != nil {
		log.Fatal(fmt.Errorf("error on connection to postgres: %w", err))
	}

	defer dbConn.Close()

	brandRepo := brand_database.NewBrandRepoImpl(dbConn)
	brandService := brand_domain.NewBrandServiceImpl(brandRepo)
	criterionController := brand_http.NewCriterionController(brandService)

	server := http_server.NewHTTPServer(
		cfg,
		log,
		version,
		criterionController,
	)
	if err = server.Start(ctx); err != nil {
		log.Fatal(fmt.Errorf("error on start http server: %w", err))
	}

	log.Info("Server stopped gracefully")
}
