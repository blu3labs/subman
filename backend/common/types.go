package types

type SubPayment struct {
	SubPlanID    int64  `json:"sub_plan_id" bson:"sub_plan_id"`
	Subscriber   string `json:"subscriber" bson:"subscriber"`
	StartTime    int64  `json:"start_time" bson:"start_time"`
	EndTime      int64  `json:"end_time" bson:"end_time"`
	Duration     int64  `json:"duration" bson:"duration"`
	PaymentToken string `json:"payment_token" bson:"payment_token"`
	Price        string `json:"price" bson:"price"`
}
