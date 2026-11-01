package config

import (
	"fmt"

	"github.com/Serega71RUS/marketpulse-collector/internal/controller/binance"
	"github.com/Serega71RUS/marketpulse-collector/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Logger  logger.Config
	Binance binance.Config
}

func New() (Config, error) {
	var config Config

	err := godotenv.Load(".env")
	if err != nil {
		return config, fmt.Errorf("godotenv.Load: %w", err)
	}

	err = envconfig.Process("", &config)
	if err != nil {
		return config, fmt.Errorf("envconfig.Process: %w", err)
	}

	return config, nil
}
