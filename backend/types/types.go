package types

type SubPayment struct {
	SubPlanID    uint64 `json:"sub_plan_id" bson:"sub_plan_id"`
	Subscriber   string `json:"subscriber" bson:"subscriber"`
	StartTime    uint64 `json:"start_time" bson:"start_time"`
	EndTime      uint64 `json:"end_time" bson:"end_time"`
	Duration     uint64 `json:"duration" bson:"duration"`
	PaymentToken string `json:"payment_token" bson:"payment_token"`
	Price        string `json:"price" bson:"price"`
	SubDeadline  uint64 `json:"sub_deadline" bson:"sub_deadline"`
}

type LastIndexedBlock struct {
	ChainID   uint64 `json:"chain_id" bson:"chain_id"`
	LastBlock uint64 `json:"last_indexed_block" bson:"last_indexed_block"`
}

type Chain struct {
	ChainID  uint64 `json:"chainId" bson:"chainId"`
	Name     string `json:"name" bson:"name"`
	RPC      string `json:"rpc" bson:"rpc"`
	Address  string `json:"address" bson:"address"`
	Interval uint64 `json:"interval" bson:"interval"`
	Confirms uint64 `json:"confirms" bson:"confirms"`
}
