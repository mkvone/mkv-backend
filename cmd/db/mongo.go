package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CoinPrice struct {
	ID        string    `bson:"_id"`
	CoinID    string    `bson:"id"`
	Name      string    `bson:"name"`
	Denom     string    `bson:"denom"`
	USD       float64   `bson:"usd"`
	KRW       float64   `bson:"krw"`
	AddedDate time.Time `bson:"addedDate"`
}

var client *mongo.Client

func ConnectToMongoDB(url string) {
	// Log the URI for debugging purposes (remove sensitive info if logging in production)
	var err error

	ctx := context.Background() // 시간 제한 없음

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		l(err)
	}
	// Optional: Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		l(err)
	}

}
func FindLatestCoinPrice(denom string) *CoinPrice {
	if client == nil {
		log.Println("MongoDB client is not initialized")
		return nil
	}

	collection := client.Database("Coin").Collection("coinprice")

	filter := bson.M{"denom": denom}
	opts := options.FindOne().SetSort(bson.D{{"addedDate", -1}}) // Sort by addedDate in descending order

	var result CoinPrice
	err := collection.FindOne(context.Background(), filter, opts).Decode(&result)
	if err != nil {
		l(err)
		return nil
	}

	return &result
}
