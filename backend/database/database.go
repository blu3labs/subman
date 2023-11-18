package database

import (
	"context"
	"os"
	"subman/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database
var index *mongo.Collection
var subscriptions *mongo.Collection
var ctx = context.Background()

func init() {
	Connect()
	SetUp()
}

func Connect() {
	mongoURI, success := os.LookupEnv("MONGO_URI")
	if !success {
		panic("MONGO_URI not found")
	}
	options := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		panic(err)
	}
	db = client.Database("subman")
	index = db.Collection("index")
	subscriptions = db.Collection("subscriptions")
}

func SetUp() {
	baseStartBlock := uint64(12570487)
	scrollStartBlock := uint64(2305901)
	index.UpdateOne(ctx, bson.M{"chainId": 84531}, bson.M{"$set": bson.M{"lastBlock": baseStartBlock}}, options.Update().SetUpsert(true))
	index.UpdateOne(ctx, bson.M{"chainId": 534351}, bson.M{"$set": bson.M{"lastBlock": scrollStartBlock}}, options.Update().SetUpsert(true))
}

func GetLastIndexedBlock(chainId uint64) (uint64, error) {
	var lastIndexedBlock types.LastIndexedBlock
	err := index.FindOne(ctx, bson.M{"chainId": chainId}).Decode(&lastIndexedBlock)
	if err != nil {
		return 0, err
	}
	return lastIndexedBlock.LastBlock, nil
}

func GetExecuteableSubscriptions() ([]types.Subscription, error) {
	var executeableSubscriptions []types.Subscription
	now := time.Now().Unix()
	filter := bson.M{
		"subscription.subDeadline":        bson.M{"$lt": now},
		"subscription.subPayment.endTime": bson.M{"$gt": now},
		"subscription.planActive":         true, "subscription.canceled": false,
	}
	cursor, err := subscriptions.Find(ctx, filter)
	if err != nil {
		return executeableSubscriptions, err
	}
	err = cursor.All(ctx, &executeableSubscriptions)
	if err != nil {
		return executeableSubscriptions, err
	}
	return executeableSubscriptions, nil
}

func BulkWrite(models []mongo.WriteModel) error {
	_, err := subscriptions.BulkWrite(ctx, models)
	if err != nil {
		return err
	}
	return nil
}

func InsertSubscription(sub types.Subscription) error {
	_, err := subscriptions.UpdateOne(ctx, filterSubscription(sub), bson.M{"$set": sub}, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

func NewDeadlineModel(subPay types.SubPayment, newDeadline uint64) mongo.WriteModel {
	filter := filterPayment(subPay)
	return mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(bson.M{"$set": bson.M{"subDeadline": newDeadline}})
}

func NewCanceledModel(subPay types.SubPayment) mongo.WriteModel {
	filter := filterPayment(subPay)
	return mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(bson.M{"$set": bson.M{"canceled": true}})
}

func NewActivatedModel(subPlanId uint64, chainId uint64) mongo.WriteModel {
	filter := bson.M{"subscription.subPayment.subPlanId": subPlanId, "subscription.chainId": chainId}
	return mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(bson.M{"$set": bson.M{"planActive": true}})
}

func NewDeactivatedModel(subPlanId uint64, chainId uint64) mongo.WriteModel {
	filter := bson.M{"subscription.subPayment.subPlanId": subPlanId, "subscription.chainId": chainId}
	return mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(bson.M{"$set": bson.M{"planActive": false}})
}

func filterPayment(subPay types.SubPayment) bson.M {
	filter := bson.M{
		"subscription.subPayment.subPlanId":    subPay.SubPlanID,
		"subscription.subPayment.subscriber":   subPay.Subscriber,
		"subscription.subPayment.startTime":    subPay.StartTime,
		"subscription.subPayment.endTime":      subPay.EndTime,
		"subscription.subPayment.duration":     subPay.Duration,
		"subscription.subPayment.paymentToken": subPay.PaymentToken,
		"subscription.subPayment.price":        subPay.Price,
	}
	return filter
}

func filterSubscription(s types.Subscription) bson.M {
	filter := bson.M{
		"subscription.subPayment":  s.SubPayment,
		"subscription.chainId":     s.ChainID,
		"subscription.subDeadline": s.SubDeadline,
		"subscription.planActive":  s.PlanActive,
		"subscription.canceled":    s.Canceled,
		"subscription.signature":   s.Signature,
	}
	return filter
}
