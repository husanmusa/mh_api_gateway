package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode string = "debug"
	// TestMode indicates service mode is test.
	TestMode string = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode string = "release"
)

type Config struct {
	ServiceName string
	Environment string // debug, test, release
	Version     string

	HTTPPort   string
	HTTPScheme string

	DefaultOffset string
	DefaultLimit  string

	SecretKey string

	AppointmentServiceHost string
	AppointmentGRPCPort    string
}

// Load ...
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "mh_api_gateway"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))
	config.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.SecretKey = cast.ToString(getOrReturnDefaultValue("SECRET_KEY", "Here$houldBe$ome$ecretKey"))

	config.AppointmentServiceHost = cast.ToString(getOrReturnDefaultValue("APPOINTMENT_SERVICE_HOST", "0.0.0.0"))
	config.AppointmentGRPCPort = cast.ToString(getOrReturnDefaultValue("APPOINTMENT_GRPC_PORT", ":9101"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
