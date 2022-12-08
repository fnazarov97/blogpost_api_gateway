package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	App         string
	AppVersion  string
	Environment string // development, staging, production

	HTTPPort string

	DefaultOffset string
	DefaultLimit  string

	AuthorServiceGrpcHost string
	AuthorServiceGrpcPort string

	ArticleServiceGrpcHost string
	ArticleServiceGrpcPort string

	AuthorizationServiceGrpcHost string
	AuthorizationServiceGrpcPort string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}
	config.App = cast.ToString(getOrReturnDefaultValue("APP", "blockpost_rest_API"))
	config.AppVersion = cast.ToString(getOrReturnDefaultValue("APP_VERSION", "1.0.1"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "development"))

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":7070"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.AuthorServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("AUTHOR_SERVICE_GRPC_HOST", "localhost"))
	config.AuthorServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("AUTHOR_SERVICE_GRPC_PORT", ":9000"))

	config.ArticleServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("ARTICLE_SERVICE_GRPC_HOST", "localhost"))
	config.ArticleServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("ARTICLE_SERVICE_GRPC_PORT", ":9000"))

	config.AuthorizationServiceGrpcHost = cast.ToString(getOrReturnDefaultValue("AUTHORIZATION_SERVICE_GRPC_HOST", "localhost"))
	config.AuthorizationServiceGrpcPort = cast.ToString(getOrReturnDefaultValue("AUTHORIZATION_SERVICE_GRPC_PORT", ":9002"))
	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
