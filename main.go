package main

import (
	"context"
	"time"
	"fmt"
	"strconv"
	"strings"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

type Room struct{
	Vnum int
	Desc string
	Mobiles []int
	Items []int
}

func InitZoneRooms(roomRange string, zoneName string) {
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
		var mobiles []int
		var items []int
		mobiles = append(mobiles, 0)
		items = append(items, 0)
		_, err = collection.InsertOne(context.Background(), bson.M{"vnum":i, "desc":"The absence of light is blinding.",
							"mobiles": mobiles, "items": items })
	}
	if err != nil {
		panic(err)
	}
}

func main() {
	InitZoneRooms("0-100", "The Void")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	var rooms Room
	collection := client.Database("zone").Collection("rooms")
	results, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	for results.Next(context.Background()) {
			err := results.Decode(&rooms)
			if err != nil {
				panic(err)
			}
//			fmt.Println(rooms.Vnum)
	}
//	res, err := collection.InsertOne(context.Background(), bson.M{"Noun":"x"})
//	res, err = collection.InsertOne(context.Background(), bson.M{"Verb":"+"})
//	res, err = collection.InsertOne(context.Background(), bson.M{"ProperNoun":"y"})

}
