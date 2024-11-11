package main

import (
	"books-api/config"
	"books-api/internal/delivery/grpc"
	api "books-api/internal/delivery/http/middleware"
	"books-api/internal/delivery/http/route"
	"books-api/internal/entity"
	"books-api/internal/gateway/externalapi"
	"books-api/internal/repository"
	services "books-api/internal/services"
	"books-api/pkg/database"
	"books-api/pkg/logger"
	"books-api/pkg/server"
	"books-api/pkg/signature"
	"books-api/pkg/xvalidator"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var (
	sqlClientRepo *database.Database
)

func main() {
	validate, _ := xvalidator.NewValidator()
	conf := config.InitAppConfig(validate)
	logger.SetupLogger(&logger.Config{
		AppENV:  conf.AppEnvConfig.AppEnv,
		LogPath: conf.AppEnvConfig.LogFilePath,
		Debug:   conf.AppEnvConfig.AppDebug,
	})
	initInfrastructure(conf)
	ginServer := server.NewGinServer(&server.GinConfig{
		HttpPort:     conf.AppEnvConfig.BookHttpPort,
		AllowOrigins: conf.AppEnvConfig.AllowOrigins,
		AllowMethods: conf.AppEnvConfig.AllowMethods,
		AllowHeaders: conf.AppEnvConfig.AllowHeaders,
	})
	// external
	signaturer := signature.NewSignature(conf.AuthConfig.JwtSecretAccessToken)
	authorExternal := externalapi.NewAuthorExternalImpl(conf)
	categoryExternal := externalapi.NewCategoryExternalImpl(conf)
	// repository
	booksRepository := repository.NewBookSQLRepository()

	// service
	booksService := services.NewBookService(sqlClientRepo.GetDB(), booksRepository, authorExternal, categoryExternal, validate)

	router := route.Router{
		App:            ginServer.App,
		BookHandler:    grpc.NewBookGRPCHandler(booksService),
		AuthMiddleware: api.NewAuthMiddleware(signaturer),
	}
	router.BookSetup(conf.AppEnvConfig.BookGRPCPort)
	echan := make(chan error)
	go func() {
		echan <- ginServer.Start()
	}()

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)

	select {
	case <-term:
		slog.Info("signal terminated detected")
	case err := <-echan:
		slog.Error("Failed to start http server", err)
	}
}

func initInfrastructure(config *config.Config) {
	//initPostgreSQL()
	sqlClientRepo = initSQL(config)
}

func initSQL(conf *config.Config) *database.Database {
	db := database.NewDatabase(conf.DatabaseConfig.Dbservice, &database.Config{
		DbHost:   conf.DatabaseConfig.Dbhost,
		DbUser:   conf.DatabaseConfig.Dbuser,
		DbPass:   conf.DatabaseConfig.Dbpassword,
		DbName:   conf.DatabaseConfig.BookDbname,
		DbPort:   strconv.Itoa(conf.DatabaseConfig.Dbport),
		DbPrefix: conf.DatabaseConfig.DbPrefix,
	})
	if conf.IsStaging() {
		db.MigrateDB(&entity.Book{})
	}
	return db
}
