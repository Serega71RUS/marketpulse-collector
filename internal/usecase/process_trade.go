package usecase

import (
	"context"

	"github.com/Serega71RUS/marketpulse-collector/internal/domain"
	"github.com/rs/zerolog/log"
)

func (u *UseCase) ProcessTrade(ctx context.Context, trade domain.Trade) error {
	log.Info().
		Str("exchange", trade.Exchange).
		Str("symbol", trade.Symbol).
		Str("price", trade.Price).
		Str("quantity", trade.Quantity).
		Msg("trade processed")

	return nil
}
