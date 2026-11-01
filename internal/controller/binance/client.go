package binance

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/Serega71RUS/marketpulse-collector/internal/dto"
	"github.com/Serega71RUS/marketpulse-collector/internal/usecase"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Host string `envconfig:"HTTP_BINANCE_HOST" required:"true"`
	Port string `envconfig:"HTTP_BINANCE_PORT" required:"true"`
}

type Client struct {
	websocket *websocket.Dialer
	host      string
	usecase   *usecase.UseCase
	stop      context.CancelFunc
	done      chan struct{}
}

func New(cfg Config, uc *usecase.UseCase) *Client {
	ctx, stop := context.WithCancel(context.Background())

	client := &Client{
		websocket: websocket.DefaultDialer,
		host:      net.JoinHostPort(cfg.Host, cfg.Port),
		usecase:   uc,
		stop:      stop,
		done:      make(chan struct{}),
	}

	go func() {
		defer close(client.done)
		client.run(ctx)
	}()

	return client
}

func (b *Client) ReadTrades(ctx context.Context) error {
	endpoint := "wss://" + b.host + "/ws/btcusdt@trade"

	conn, _, err := b.websocket.DialContext(ctx, endpoint, nil)
	if err != nil {
		return fmt.Errorf("connect Binance: %w", err)
	}
	defer conn.Close()

	stopClose := context.AfterFunc(ctx, func() {
		conn.Close()
	})
	defer stopClose()

	for {
		var event dto.BinanceTradeEventOutput

		err := conn.ReadJSON(&event)

		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			return fmt.Errorf("read trade: %w", err)
		}

		if event.EventType != "trade" {
			log.Debug().Str("event_type", event.EventType).Msg("event skipped")
			continue
		}

		trade := dto.ToDomainTrade(event)

		err = b.usecase.ProcessTrade(ctx, trade)
		if err != nil {
			return fmt.Errorf("process trade: %w", err)
		}

	}
}

func (b *Client) run(ctx context.Context) {
	for {
		err := b.ReadTrades(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			log.Error().Err(err).Msg("binance reconnect, next attempt in 3 seconds")
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(3 * time.Second):
		}
	}

}

func (b *Client) Close() {
	log.Info().Msg("binance client: closing")
	b.stop()

	<-b.done
	log.Info().Msg("binance client: closed")
}
