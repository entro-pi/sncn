package main

import (

  "fmt"
  "context"
  "time"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func insertBroadcasts(broadside []Broadcast) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	collection := client.Database("broadcasts").Collection("general")
  for i := 0;i < len(broadside);i++ {
    update := bson.M{"broadcast":bson.M{"event":broadside[i].Event,"ref":broadside[i].Ref,"payload":bson.M{"channel":broadside[i].Payload.Channel,"message":broadside[i].Payload.Message,"game":broadside[i].Payload.Game, "name":broadside[i].Payload.Name, "row":broadside[i].Payload.Row, "col":broadside[i].Payload.Col, "selected":false,"bigmessage":broadside[i].Payload.BigMessage}}}

  	result, err := collection.InsertOne(context.Background(), update)
  	if err != nil {
  		panic(err)
  	}
    fmt.Println("\033[38:2:255:0:0m", result, "\033[0m")

  }
}

func updateRoom(play Player, populated []Space) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"vnum": bson.M{"$eq":play.CurrentRoom.Vnum}}
	collection := client.Database("zones").Collection("Spaces")
	update := bson.M{"$set": bson.M{"vnums":populated[play.CurrentRoom.Vnum].Vnums,
		"zone":populated[play.CurrentRoom.Vnum].Zone,"vnum":populated[play.CurrentRoom.Vnum].Vnum,
		 "desc":populated[play.CurrentRoom.Vnum].Desc,"exits": populated[play.CurrentRoom.Vnum].Exits,
			"mobiles": populated[play.CurrentRoom.Vnum].Mobiles, "items": populated[play.CurrentRoom.Vnum].Items,
			 "altered": true,"zonepos":populated[play.CurrentRoom.Vnum].ZonePos, "zonemap": populated[play.CurrentRoom.Vnum].ZoneMap }}

	result, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		panic(err)
	}
	fmt.Println("\033[38:2:255:0:0m", result, "\033[0m")
}

func updateZoneMap(play Player, populated []Space) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"zone": bson.M{"$eq":play.CurrentRoom.Zone}}
	collection := client.Database("zones").Collection("Spaces")
	findOptions := options.Find()
  findOptions.SetLimit(1000)
	result, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		panic(err)
	}
  fmt.Println("\033[38:2:150:0:0m",play.CurrentRoom.Zone)
	defer result.Close(context.Background())
	for result.Next(context.Background()) {
		var current Space
		err := result.Decode(&current)
		if err != nil {
			panic(err)
		}
		filter = bson.M{"vnum": bson.M{"$eq":current.Vnum}}
//		update := bson.M{"$set": bson.M{"zonepos":populated[current.Vnum].ZonePos, "zonemap": populated[play.CurrentRoom.Vnum].ZoneMap }}
    update := bson.M{"$set": bson.M{"zonepos":populated[current.Vnum].ZonePos, "zonemap": play.CurrentRoom.ZoneMap }}

		result, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
		if err != nil {
			panic(err)
		}
		fmt.Println("\033[38:2:255:0:0m", result, "\033[0m")
	}

}
