package main

import (
	"context"
	"time"
	"strconv"
	"strings"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func InitZone(roomRange string, zoneName string) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	collection := client.Database("zone").Collection("rooms")
	_, err = collection.InsertOne(context.Background(), bson.M{"zone":zoneName})
	if err != nil {
		panic(err)
	}
	vnums := strings.Split(roomRange, "-")
	vnumStart, err := strconv.Atoi(vnums[0])
	if err != nil {
		panic(err)
	}

	vnumEnd, err := strconv.Atoi(vnums[1])
	if err != nil {
		panic(err)
	}
	_, err = collection.InsertOne(context.Background(), bson.M{"vnums":roomRange})
	for i := vnumStart;i < vnumEnd;i++ {
		vnum := strconv.Itoa(i)
		_, err = collection.InsertOne(context.Background(), bson.M{"vnum":vnum})
	}
	if err != nil {
		panic(err)
	}
}

func main() {
	InitZone("0-100", "The Void")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	_ = client.Database("zone").Collection("rooms")
//	res, err := collection.InsertOne(context.Background(), bson.M{"Noun":"x"})
//	res, err = collection.InsertOne(context.Background(), bson.M{"Verb":"+"})
//	res, err = collection.InsertOne(context.Background(), bson.M{"ProperNoun":"y"})

}
