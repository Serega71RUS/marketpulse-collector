package binance

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Serega71RUS/marketpulse-collector/internal/dto"
	"github.com/gorilla/websocket"
)

type Config struct {
	Host string `envconfig:"HTTP_BINANCE_HOST"`
	Port string `envconfgi:"HTTP_BINANCE_PORT"`
}

type BinanceHttpClient struct {
	client http.Client
	host   string
}

//func New(cfg Config) *BinanceHttpClient {
//	client := &BinanceHttpClient{
//		client: http.Client{},
//		host:   net.JoinHostPort(cfg.Host, cfg.Port),
//	}
//}

func ReadTrades(ctx context.Context) error {
	const endpoint = "wss://stream.binance.com:443/ws/btcusdt@trade"

	conn, _, err := websocket.DefaultDialer.DialContext(ctx, endpoint, nil)
	if err != nil {
		return fmt.Errorf("connect Binance: %w", err)
	}
	defer conn.Close()

	for range 10 {
		var event dto.BinanceTradeEventOutput
		if err := conn.ReadJSON(&event); err != nil {
			return fmt.Errorf("read trade: %w", err)
		}

		trade := dto.ToDomainTrade(event)
		fmt.Printf("%#v\n", trade)
	}
	return nil
}

//func (c *BinanceHttpClient) TradeStream(symbol string, ctx context.Context) {
//	var tradeStream = "ws/" + symbol + "@trade"
//
//	path := fmt.Sprintf("wss://%s/%s", c.host, tradeStream)
//
//	req, err := http.NewRequestWithContext(ctx, http.M)
//
//	c.client.Do()
//
//	input := dto.BinanceTradeEventInput{
//		Method: "SUBSCRIBE",
//		Params: []string{"bnbusdt@trade"},
//		ID:     "e9d6b4349871b40611412680b3445fac",
//	}
//
//	c
//}
