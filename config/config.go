package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AppName     string `default:"CMS"`
	AppEnv      string `default:"development"`
	AppHost     string `default:"localhost"`
	AppPort     int    `default:"8000"`
	LogLevel    string `default:"debug"`
	DatabaseURL string `default:"mysql://root:password@localhost:3306/cms"`
}

var Cfg *AppConfig

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using default values.")
	}

	Cfg = &AppConfig{}
	err = envconfig.Process("", Cfg)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err == nil {
		Cfg.AppPort = port
	}

	// Wait for MySQL to start
	time.Sleep(10 * time.Second)
}
