package query

type BlockByNumber struct {
	Result Result `json:"result"`
}

type Result struct {
	Number       string         `json:"number"`
	Transactions []Transactions `json:"transactions,omitempty"`
}

type Transactions struct {
	BlockNumber string `json:"blockNumber"`
	Value       string `json:"value"`
}
