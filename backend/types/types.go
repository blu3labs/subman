package types

type SubPayment struct {
	SubPlanID    uint64 `json:"subPlanId" bson:"subPlanId"`
	Subscriber   string `json:"subscriber" bson:"subscriber"`
	StartTime    uint64 `json:"startTime" bson:"startTime"`
	EndTime      uint64 `json:"endTime" bson:"endTime"`
	Duration     uint64 `json:"duration" bson:"duration"`
	PaymentToken string `json:"paymentToken" bson:"paymentToken"`
	Price        string `json:"price" bson:"price"`
	// SubDeadline  uint64 `json:"subDeadline" bson:"subDeadline"`
	// PlanActive   bool   `json:"deactivated" bson:"deactivated"`
	// Cancelled    bool   `json:"cancelled" bson:"cancelled"`
}

type Subscription struct {
	SubPayment  SubPayment `json:"subPayment" bson:"subPayment"`
	ChainID     uint64     `json:"chainId" bson:"chainId"`
	SubDeadline uint64     `json:"subDeadline" bson:"subDeadline"`
	PlanActive  bool       `json:"planActive" bson:"planActive"`
	Cancelled   bool       `json:"cancelled" bson:"cancelled"`
	Signature   string     `json:"signature" bson:"signature"`
}

type LastIndexedBlock struct {
	ChainID   uint64 `json:"chainId" bson:"chainId"`
	LastBlock uint64 `json:"lastBlock" bson:"lastBlock"`
}

type Chain struct {
	ChainID  uint64 `json:"chainId" bson:"chainId"`
	Name     string `json:"name" bson:"name"`
	RPC      string `json:"rpc" bson:"rpc"`
	Address  string `json:"address" bson:"address"`
	Interval uint64 `json:"interval" bson:"interval"`
	Confirms uint64 `json:"confirms" bson:"confirms"`
}
