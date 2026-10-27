package domain

import "time"

type Trade struct {
	Exchange     string
	Symbol       string
	ID           string
	Price        string
	Quantity     string
	Timestamp    time.Time
	BuyerIsMaker bool
}
