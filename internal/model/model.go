package model

type Message struct {
	EventClass string     `json:"e"`
	EventTime int64       `json:"E"`
	Content string        `json:"s"`
	TradeID int           `json:"a"`
	Price string          `json:"p"`
	Qty string            `json:"q"`
	FrontID int           `json:"f"`
	LastID int            `json:"l"`
	TradeTime int64       `json:"T"`
	IsMaker bool          `json:"m"`
	IsIgnore bool         `json:"M"`
}

type StreamMsg struct {
	Stream string  `json:"stream"`
	Data Message   `json:"data"`
}
