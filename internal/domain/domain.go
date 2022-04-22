package domain

type BlockInfo struct {
	BlockNumber  string  `json:"blockNumber" bson:"_id"`
	Transactions int     `json:"transactions" bson:"transactions"`
	Count        float64 `json:"count" bson:"count"`
}
