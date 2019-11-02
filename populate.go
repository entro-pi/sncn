package main

import (

  "context"
  "time"
  "strings"
  "strconv"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func PopulateAreaMobiles() []Mobile {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	var Mobiles []Mobile
	collection := client.Database("npcs").Collection("mobiles")
	results, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	for results.Next(context.Background()) {

			var Mobile Mobile
			err := results.Decode(&Mobile)
			if err != nil {
				panic(err)
			}
			Mobiles = append(Mobiles, Mobile)

//			fmt.Println(Spaces.Vnum)
	}
	return Mobiles
}

func PopulateAreaBuild(rangeVnums string) []Space {

  beginString := strings.Split(rangeVnums, "-")[0]

  endString := strings.Split(rangeVnums, "-")[1]

  begin, err := strconv.Atoi(beginString)
  if err != nil {
    panic(err)
  }
  end, err := strconv.Atoi(endString)
  if err != nil {
    panic(err)
  }
  length := end - begin
	areas := make([]Space, length)
	return areas
}

func PopulateAreas() []Space {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	var Spaces []Space
	collection := client.Database("zones").Collection("Spaces")
	results, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	for results.Next(context.Background()) {

			var Space Space
			err := results.Decode(&Space)
			if err != nil {
				panic(err)
			}
			Spaces = append(Spaces, Space)

//			fmt.Println(Spaces.Vnum)
	}
	return Spaces
}
