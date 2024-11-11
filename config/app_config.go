package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	AppEnv           string
	AppDebug         bool
	AppVersion       string   `validate:"required,startswith=v,alphanum" name:"APP_VERSION"`
	AppName          string   `validate:"required" name:"APP_NAME"`
	UserHttpPort     string   `validate:"number" name:"USER_HTTP_PORT"`
	UserGRPCPort     string   `validate:"number" name:"USER_GRPC_PORT"`
	AuthorHTTPPort   string   `validate:"number" name:"AUTHOR_HTTP_PORT"`
	AuthorGRPCPort   string   `validate:"number" name:"AUTHOR_GRPC_PORT"`
	CategoryHTTPPort string   `validate:"number" name:"CATEGORY_HTTP_PORT"`
	CategoryGRPCPort string   `validate:"number" name:"CATEGORY_GRPC_PORT"`
	BookHttpPort     string   `validate:"number" name:"BOOK_HTTP_PORT"`
	BookGRPCPort     string   `validate:"number" name:"BOOK_GRPC_PORT"`
	AllowOrigins     []string `name:"HTTP_ALLOW_ORIGINS"`
	AllowMethods     []string `name:"HTTP_ALLOW_METHODS"`
	AllowHeaders     []string `name:"HTTP_ALLOW_HEADERS"`
	LogFilePath      string   `validate:"required" name:"LOG_PATH"`
}

func AppConfigInit() *AppConfig {
	return &AppConfig{
		AppEnv:           viper.GetString("APP_ENV"),
		AppDebug:         viper.GetBool("APP_DEBUG"),
		AppVersion:       viper.GetString("APP_VERSION"),
		AppName:          viper.GetString("APP_NAME"),
		UserHttpPort:     viper.GetString("USER_HTTP_PORT"),
		UserGRPCPort:     viper.GetString("USER_GRPC_PORT"),
		AuthorHTTPPort:   viper.GetString("AUTHOR_HTTP_PORT"),
		AuthorGRPCPort:   viper.GetString("AUTHOR_GRPC_PORT"),
		CategoryHTTPPort: viper.GetString("CATEGORY_HTTP_PORT"),
		CategoryGRPCPort: viper.GetString("CATEGORY_GRPC_PORT"),
		BookHttpPort:     viper.GetString("BOOK_HTTP_PORT"),
		BookGRPCPort:     viper.GetString("BOOK_GRPC_PORT"),
		LogFilePath:      viper.GetString("LOG_PATH"),
		AllowOrigins:     viper.GetStringSlice("ALLOW_ORIGINS"),
		AllowMethods:     viper.GetStringSlice("ALLOW_METHODS"),
		AllowHeaders:     viper.GetStringSlice("ALLOW_HEADERS"),
	}
}
