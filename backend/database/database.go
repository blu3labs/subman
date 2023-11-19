package database

import (
	"context"
	"log"
	"os"
	"subman/types"
	"time"

	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
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
		"subDeadline":        bson.M{"$lt": now},
		"subPayment.endTime": bson.M{"$gt": now},
		"planActive":         true, "canceled": false,
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
	filter := bson.M{"subPayment.subPlanId": subPlanId, "subscription.chainId": chainId}
	return mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(bson.M{"$set": bson.M{"planActive": true}})
}

func NewDeactivatedModel(subPlanId uint64, chainId uint64) mongo.WriteModel {
	filter := bson.M{"subPayment.subPlanId": subPlanId, "subscription.chainId": chainId}
	return mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(bson.M{"$set": bson.M{"planActive": false}})
}

func filterPayment(subPay types.SubPayment) bson.M {
	filter := bson.M{
		"subPayment.subPlanId":    subPay.SubPlanID,
		"subPayment.subscriber":   subPay.Subscriber,
		"subPayment.startTime":    subPay.StartTime,
		"subPayment.endTime":      subPay.EndTime,
		"subPayment.duration":     subPay.Duration,
		"subPayment.paymentToken": subPay.PaymentToken,
		"subPayment.price":        subPay.Price,
	}
	return filter
}

func filterSubscription(s types.Subscription) bson.M {
	filter := bson.M{
		"subPayment":  s.SubPayment,
		"chainId":     s.ChainID,
		"subDeadline": s.SubDeadline,
		"planActive":  s.PlanActive,
		"canceled":    s.Canceled,
		"signature":   s.Signature,
	}
	return filter
}
