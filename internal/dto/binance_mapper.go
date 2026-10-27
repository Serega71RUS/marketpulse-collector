package dto

import (
	"strconv"
	"time"

	"github.com/Serega71RUS/marketpulse-collector/internal/domain"
)

func ToDomainTrade(event BinanceTradeEvent) domain.Trade {
	return domain.Trade{
		Exchange:     "binance",
		Symbol:       event.Symbol,
		ID:           strconv.FormatInt(event.TradeID, 10),
		Price:        event.Price,
		Quantity:     event.Quantity,
		Timestamp:    time.UnixMilli(event.TradeTime).UTC(),
		BuyerIsMaker: event.BuyerIsMaker,
	}
}
