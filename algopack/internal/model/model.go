package model

type TradeData struct {
	Data struct {
		Data [][]interface{} `json:"data"`
	} `json:"data"`
}

type TicketPredict struct {
	Title   string  `json:"ticker"`
	Predict float64 `json:"target"`
}

type TicketData struct {
	Ticket []interface{} `json:"ticket"`
}
