package main

import (
	"context"

	"github.com/Serega71RUS/marketpulse-collector/config"
	"github.com/Serega71RUS/marketpulse-collector/internal/app"
	"github.com/Serega71RUS/marketpulse-collector/pkg/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("config error")
	}

	logger.Init(cfg.Logger)

	ctx := context.Background()

	err = app.Run(ctx, cfg)
	if err != nil {
		log.Error().Err(err).Msg("app error")
	}
}
