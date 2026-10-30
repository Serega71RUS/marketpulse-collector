package dto

type BinanceTradeEventInput struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	ID     string   `json:"id"`
}

type BinanceTradeEventOutput struct {
	EventType    string `json:"e"`
	EventTime    int64  `json:"E"`
	Symbol       string `json:"s"`
	TradeID      int64  `json:"t"`
	Price        string `json:"p"`
	Quantity     string `json:"q"`
	TradeTime    int64  `json:"T"`
	BuyerIsMaker bool   `json:"m"`
	Ignore       bool   `json:"M"`
}
