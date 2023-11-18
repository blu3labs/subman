package database

import (
	"context"
	"os"
	"subman/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var db *mongo.Database
var index *mongo.Collection
var subPayments *mongo.Collection
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
	subPayments = db.Collection("subPayments")
}

func SetUp() {
	baseStartBlock := uint64(0)
	scrollStartBlock := uint64(0)
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

func BulkWrite(models []mongo.WriteModel) error {
	_, err := subPayments.BulkWrite(ctx, models)
	if err != nil {
		return err
	}
	return nil
}

func NewDeadlineModel(subPay types.SubPayment, newDeadline uint64) mongo.WriteModel {
	filter := filterAll(subPay)
	return mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(bson.M{"$set": bson.M{"subDeadline": newDeadline}})
}

func NewCanceledModel(subPay types.SubPayment) mongo.WriteModel {
	filter := filterAll(subPay)
	return mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(bson.M{"$set": bson.M{"cancelled": true}})
}

func NewActivatedModel(subPlanId uint64) mongo.WriteModel {
	filter := bson.M{"subscrioption.subPayment.subPlanId": subPlanId}
	return mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(bson.M{"$set": bson.M{"planActive": true}})
}

func NewDeactivatedModel(subPlanId uint64) mongo.WriteModel {
	filter := bson.M{"subscrioption.subPayment.subPlanId": subPlanId}
	return mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(bson.M{"$set": bson.M{"planActive": false}})
}

func filterAll(subPay types.SubPayment) bson.M {
	filter := bson.M{
		"subscrioption.subPayment.subPlanId":    subPay.SubPlanID,
		"subscrioption.subPayment.subscriber":   subPay.Subscriber,
		"subscrioption.subPayment.startTime":    subPay.StartTime,
		"subscrioption.subPayment.endTime":      subPay.EndTime,
		"subscrioption.subPayment.duration":     subPay.Duration,
		"subscrioption.subPayment.paymentToken": subPay.PaymentToken,
		"subscrioption.subPayment.price":        subPay.Price,
	}
	return filter
}
