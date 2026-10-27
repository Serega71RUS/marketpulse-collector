package dto

import (
	"testing"
	"time"

	"github.com/Serega71RUS/marketpulse-collector/internal/domain"
	"github.com/stretchr/testify/require"
)

func Test_ToDomainTrade_Success(t *testing.T) {
	input := BinanceTradeEvent{
		Symbol:       "BNBUSDT",
		TradeID:      1608510363,
		Price:        "743.71000000",
		Quantity:     "0.79500000",
		TradeTime:    1793112313696,
		BuyerIsMaker: false,
	}

	output := domain.Trade{
		Exchange:     "binance",
		Symbol:       "BNBUSDT",
		ID:           "1608510363",
		Price:        "743.71000000",
		Quantity:     "0.79500000",
		Timestamp:    time.Date(2026, time.October, 27, 14, 45, 13, 696000000, time.UTC),
		BuyerIsMaker: false,
	}

	actual := ToDomainTrade(input)

	require.Equal(t, output, actual)
}
