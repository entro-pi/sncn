package main

import (
	"bufio"
	"os"
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
	Vnums string
	Zone string
	Vnum int
	Desc string
	Mobiles []int
	Items []int
}
type Player struct {
	Name string
	Title string
	Inventory []int
	Equipment []int

	Str int
	Int int
	Dex int
	Wis int
	Con int
	Cha int
}

func DescribePlayer(play Player) {
	fmt.Println("======================")
	fmt.Println("\033[38:2:0:200:0mStrength     :\033[0m", play.Str)
	fmt.Println("\033[38:2:0:200:0mIntelligence :\033[0m", play.Int)
	fmt.Println("\033[38:2:0:200:0mDexterity    :\033[0m", play.Dex)
	fmt.Println("\033[38:2:0:200:0mWisdom       :\033[0m", play.Wis)
	fmt.Println("\033[38:2:0:200:0mConstitution :\033[0m", play.Con)
	fmt.Println("\033[38:2:0:200:0mCharisma     :\033[0m", play.Cha)
	fmt.Println("======================")
}



func InitPlayer(name string) Player {
	var play Player
	var inv []int
	var equ []int
	inv = append(inv, 1)
	equ = append(equ, 1)
	play.Name = name
	play.Title = "The Unknown"
	play.Inventory = inv
	play.Equipment = equ
	play.Str = 1
	play.Int = 1
	play.Dex = 1
	play.Wis = 1
	play.Con = 1
	play.Cha = 1
	return play
}

func InitZoneRooms(roomRange string, zoneName string, desc string) {
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

	vnums := strings.Split(roomRange, "-")
	vnumStart, err := strconv.Atoi(vnums[0])
	if err != nil {
		panic(err)
	}

	vnumEnd, err := strconv.Atoi(vnums[1])
	if err != nil {
		panic(err)
	}
	for i := vnumStart;i < vnumEnd;i++ {
		var mobiles []int
		var items []int
		mobiles = append(mobiles, 0)
		items = append(items, 0)
		_, err = collection.InsertOne(context.Background(), bson.M{"vnums":roomRange,"zone":zoneName,"vnum":i, "desc":desc,
							"mobiles": mobiles, "items": items })
	}
	if err != nil {
		panic(err)
	}
}

func PopulateAreas() []Room {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	var rooms []Room
	collection := client.Database("zone").Collection("rooms")
	results, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	for results.Next(context.Background()) {

			var room Room
			err := results.Decode(&room)
			if err != nil {
				panic(err)
			}
			rooms = append(rooms, room)

//			fmt.Println(rooms.Vnum)
	}
	return rooms
}
func DescribeRoom(vnum int, rooms []Room) {
	for i := 0; i < len(rooms);i++ {
		if rooms[i].Vnum == vnum {
			fmt.Println(rooms[i].Zone)
			fmt.Println(rooms[i].Desc)
		}
	}
}

func main() {
	//TODO Get the rooms that are already loaded in the database and skip
	//if vnum is taken
	//Get the flags passed in
	var populated []Room
	var play Player
	if len(os.Args) > 1 {
		if os.Args[1] == "--init" {
			//TODO testing suite - one test will be randomly generating 10,000 rooms
			//and seeing if the system can take it
			InitZoneRooms("0-100", "The Void", "The absence of light is blinding.")
			InitZoneRooms("100-150", "Midgaard", "I wonder what day is recycling day.")
			populated = PopulateAreas()
			play = InitPlayer("FSM")
		} else {
			os.Exit(1)
		}
	} else {
		os.Exit(1)
	}


	//Game loop
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		input := scanner.Text()

		if strings.HasPrefix(input, "look at") {
			splitCommand := strings.Split(input, "at")
			stripped := strings.TrimSpace(splitCommand[1])
			vnumLook, _ := strconv.Atoi(stripped)
			DescribeRoom(vnumLook, populated)
		}
		if strings.HasPrefix(input, "go to") {
			splitCommand := strings.Split(input, "to")
			stripped := strings.TrimSpace(splitCommand[1])
			inp, err := strconv.Atoi(stripped)
			if err != nil {
				fmt.Println("Error converting a stripped string")
			}
			fmt.Println(play.Name, play.Inventory, play.Equipment)
			fmt.Println(populated[inp].Vnum, populated[inp].Vnums, populated[inp].Zone)

		}
		if input == "score" {
			DescribePlayer(play)
		}
	}
//	res, err := collection.InsertOne(context.Background(), bson.M{"Noun":"x"})
//	res, err = collection.InsertOne(context.Background(), bson.M{"Verb":"+"})
//	res, err = collection.InsertOne(context.Background(), bson.M{"ProperNoun":"y"})

}
