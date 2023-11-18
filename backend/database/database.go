package database

import (
	"context"
	"os"
	converter "subman/converter"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LastFetchedBlock struct {
	ChainID              string `json:"chainId" bson:"chainId"`
	LastInitializedBlock uint64 `json:"lastInitializedBlock" bson:"lastInitializedBlock"`
	LastCompletedBlock   uint64 `json:"lastCompletedBlock" bson:"lastCompletedBlock"`
}

var client *mongo.Client
var db *mongo.Database
var subPayments *mongo.Collection
var ctx = context.Background()

func init() {
	Connect()
	SetUp()
}

func Connect() {
	mongoURI := os.Getenv("MONGO_URI")
	options := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, options)

	if err != nil {
		panic(err)
	}
	db = client.Database("subman")
	subPayments = db.Collection("subPayments")
}

func SetUp() {
	//todo add initial blocks
	chains.UpdateOne(ctx, bson.M{"chainId": "97"}, bson.M{"$set": bson.M{"lastInitializedBlock": 0}}, options.Update().SetUpsert(true))
	chains.UpdateOne(ctx, bson.M{"chainId": "80001"}, bson.M{"$set": bson.M{"lastInitializedBlock": 0}}, options.Update().SetUpsert(true))
	chains.UpdateOne(ctx, bson.M{"chainId": "11155111"}, bson.M{"$set": bson.M{"lastInitializedBlock": 0}}, options.Update().SetUpsert(true))
}

func BulkInitialize(items []converter.Item) error {
	var models []mongo.WriteModel
	for _, item := range items {
		count, err := requests.CountDocuments(ctx, bson.M{"hash": item.Hash})
		if err != nil {
			return err
		}
		if count == 0 {
			models = append(models, mongo.NewInsertOneModel().SetDocument(item))
		}
	}
	_, err := requests.BulkWrite(ctx, models)
	return err
}

func BulkComplete(hash []string) error {
	var models []mongo.WriteModel
	for _, h := range hash {
		models = append(models, mongo.NewUpdateOneModel().SetFilter(bson.M{"hash": h}).SetUpdate(bson.M{"$set": bson.M{"completed": true}}))
	}
	_, err := requests.BulkWrite(ctx, models)
	return err
}

func GetUserRequests(user string) ([]converter.Item, error) {
	filter := bson.M{"request.user": user}
	cursor, err := requests.Find(ctx, filter, options.Find().SetSort(bson.M{"createdAt": -1}))
	if err != nil {
		return nil, err
	}
	var items []converter.Item
	if err = cursor.All(ctx, &items); err != nil { //todo first time using this, maybe it's not the best way
		return nil, err
	}
	return items, nil
}

func GetNotSignedByValidator(validator string) ([]converter.Item, error) {
	filter := bson.M{"signatures": bson.M{"$not": bson.M{"$elemMatch": bson.M{"validator": validator}}}}
	cursor, err := requests.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var items []converter.Item
	if err = cursor.All(ctx, &items); err != nil { //todo first time using this, maybe it's not the best way
		return nil, err
	}
	return items, nil
}

func SignRequest(hash string, validator string, signature string) error {
	filter := bson.M{"hash": hash}
	var item converter.Item
	err := requests.FindOne(ctx, filter).Decode(&item)
	if err != nil {
		return err
	}
	_, err = verifier.VerifySignature(item.Request, hash, validator, signature)
	if err != nil {
		return err
	}
	update := bson.M{"$push": bson.M{"signatures": bson.M{"validator": validator, "signature": signature}}}
	_, err = requests.UpdateOne(ctx, filter, update)
	return err
}

func GetAutoSendAndNotCompleted() ([]converter.Item, error) {
	filter := bson.M{"request.autoSend": true, "completed": false}
	cursor, err := requests.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var items []converter.Item
	if err = cursor.All(ctx, &items); err != nil { //todo first time using this, maybe it's not the best way
		return nil, err
	}
	return items, nil
}

func InsertChain(chainId string) error {
	filter := bson.M{"chainId": chainId}
	count, err := chains.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	if count == 0 {
		chain := LastFetchedBlock{
			ChainID:              chainId,
			LastInitializedBlock: 0,
			LastCompletedBlock:   0,
		}
		_, err := chains.InsertOne(ctx, chain)
		return err
	}
	return err
}

func UpsertLastInitializedBlock(chainId string, lastInitializedBlock uint64) error {
	filter := bson.M{"chainId": chainId}
	update := bson.M{"$set": bson.M{"lastInitializedBlock": lastInitializedBlock}}
	_, err := chains.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return err
}

func UpsertLastCompletedBlock(chainId string, lastCompletedBlock uint64) error {
	filter := bson.M{"chainId": chainId}
	update := bson.M{"$set": bson.M{"lastCompletedBlock": lastCompletedBlock}}
	_, err := chains.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return err
}

func GetLastFetchedBlock(chainId string) (LastFetchedBlock, error) {
	filter := bson.M{"chainId": chainId}
	var lastFetchedBlock LastFetchedBlock
	err := chains.FindOne(ctx, filter).Decode(&lastFetchedBlock)
	return lastFetchedBlock, err
}
