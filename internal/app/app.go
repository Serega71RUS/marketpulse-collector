package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Serega71RUS/marketpulse-collector/config"
	"github.com/Serega71RUS/marketpulse-collector/internal/controller/binance"
	"github.com/Serega71RUS/marketpulse-collector/internal/usecase"
	"github.com/rs/zerolog/log"
)

func Run(ctx context.Context, c config.Config) error {
	uc := usecase.New()

	binanceWSClient := binance.New(c.Binance, uc)

	log.Info().Msg("app: started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig

	log.Info().Msg("app: got signal to stop")

	binanceWSClient.Close()

	log.Info().Msg("app: stopped")

	return nil
}
