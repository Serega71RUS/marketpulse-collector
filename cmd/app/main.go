package main

import (
	"context"
	"fmt"

	"github.com/Serega71RUS/marketpulse-collector/internal/controller/binance"
)

func main() {
	err := binance.ReadTrades(context.Background())
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
