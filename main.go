package main

import (
	"bufio"
	"os"
	"context"
	"time"
	"fmt"
	"strconv"
	"strings"
	"math/rand"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/SolarLune/dngn"
)
type Chat struct {
	Message string
	Time time.Time
}
type Space struct{
	Room dngn.Room
	Vnums string
	Zone string
	Vnum int
	Desc string
	Mobiles []int
	Items []int
	CoreBoard string
}
type Player struct {
	Name string
	Title string
	Inventory []int
	Equipment []int
	CoreBoard string
	CurrentRoom Space

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

func InitZoneSpaces(SpaceRange string, zoneName string, desc string) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	collection := client.Database("zone").Collection("Spaces")

	vnums := strings.Split(SpaceRange, "-")
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
		_, err = collection.InsertOne(context.Background(), bson.M{"vnums":SpaceRange,"zone":zoneName,"vnum":i, "desc":desc,
							"mobiles": mobiles, "items": items })
	}
	if err != nil {
		panic(err)
	}
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
	collection := client.Database("zone").Collection("Spaces")
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
func DescribeSpace(vnum int, Spaces []Space) {
	for i := 0; i < len(Spaces);i++ {
		if Spaces[i].Vnum == vnum {
			fmt.Println(Spaces[i].Zone)
			fmt.Println(Spaces[i].Desc)
		}
	}
}
func addPfile(play Player) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	collection := client.Database("pfiles").Collection("Players")
	_, err = collection.InsertOne(context.Background(), bson.M{"name":play.Name,"title":play.Title,"inventory":play.Inventory, "equipment":play.Equipment,
						"coreboard": play.CoreBoard, "str": play.Str, "int": play.Int, "dex": play.Dex, "wis": play.Wis, "con":play.Con, "cha":play.Cha })
}
func savePfile(play Player) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	collection := client.Database("pfiles").Collection("Players")
	_, err = collection.UpdateOne(context.Background(), options.Update().SetUpsert(true), bson.M{"name":play.Name,"title":play.Title,"inventory":play.Inventory, "equipment":play.Equipment,
						"coreboard": play.CoreBoard, "str": play.Str, "int": play.Int, "dex": play.Dex, "wis": play.Wis, "con":play.Con, "cha":play.Cha })
}

//TODO make this modular
func createMobiles(name string) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	collection := client.Database("mobiles").Collection("lvl1")
	_, err = collection.InsertOne(context.Background(), bson.M{"name":name,
						"str": 1, "int": 1, "dex": 1, "wis": 1, "con":1, "cha":1, "challengedice":1 })
}

//TODO make this modular
func createChat(message string, play Player) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	//process the strings
	if len(message) >= 180 {
		message = message[:180]
	}

	collection := client.Database("chat").Collection("log")
	_, err = collection.InsertOne(context.Background(), bson.M{"name":play.Name,
						"message":message, "time":time.Now() })
	if err != nil {
		panic(err)
	}
}


func genMap(play Player, populated []Space) (Player, []Space) {
	//Create a room map
	Room := dngn.NewRoom(50, 30)
	splits := rand.Intn(75)
	Room.GenerateBSP('%', 'D', splits)
//	_, err = collection.InsertOne(context.Background(), bson.M{"room":Room})
//	if err != nil {
//		panic(err)
//	}

	populated[0].CoreBoard = ""
	play.CoreBoard = ""
	fmt.Println("Generating and populating map")
	for i := 0;i < len(Room.Data);i++ {
		if i == 0 {
			continue
		}
	//				fmt.Println(populated[0].Room.Data[populated[0].Room.Width-1][i])
			value := string(Room.Data[i])
			newValue := ""
			for s := 0;s < len(value);s++ {
				if s == 0 {
					continue
				}
				if string(value[s]) == " " {
					ChanceTreasure := "\033[48:2:200:150:0mT\033[0m"
					if rand.Intn(100) > 98 {
							newValue += ChanceTreasure
							continue
					}
					if rand.Intn(100) > 95 {
						ChanceMonster := "\033[48:2:200:50:50mM\033[0m"
						newValue += ChanceMonster
						continue
					}else {
						newValue += string(value[s])
					}
				}else {
					newValue += string(value[s])
				}

			}

			newValue = strings.ReplaceAll(newValue, "%", "\033[48:2:0:150:150m%\033[0m")
			newValue = strings.ReplaceAll(newValue, "D", "\033[38:2:200:150:150mD\033[0m")
			newValue = strings.ReplaceAll(newValue, " ", "\033[48:2:0:200:150m \033[0m")
			populated[0].CoreBoard += newValue + "\n"
			play.CoreBoard += newValue + "\n"

	}
	return play, populated
}

const (
	mapPos = "\033[0;0H"
	descPos = "\033[0;50H"
	chatStart = "\033[38:2:200:50:50m{{=\033[38:150:50:150m"
	chatEnd = "\033[38:2:200:50:50m=}}"
	end = "\033[0m"

)

func showDesc(room Space) {
	fmt.Printf(descPos)
	splitOnNewline := strings.Split(room.Desc, "\n")
	for i := 0;i < len(splitOnNewline);i++ {
		splitPos := fmt.Sprint("\033["+strconv.Itoa(i)+";50H"+splitOnNewline[i]+end)
		fmt.Printf(splitPos)
	}
}

func showChat() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	collection := client.Database("chat").Collection("log")
	mess, err := collection.Find(context.Background(), bson.M{})

	count := 0
	for mess.Next(context.Background()) {
		var chatMess Chat
		err := mess.Decode(&chatMess)
		if err != nil {
			panic(err)
		}
		chatPos := fmt.Sprintf("\033["+strconv.Itoa(count)+";180H")
		count++
		fmt.Printf(chatPos)
		fmt.Printf(chatStart)
		fmt.Printf(chatMess.Message + " ")
		fmt.Printf(chatEnd)
		fmt.Printf(end)

	}
}

func main() {
	//TODO Get the Spaces that are already loaded in the database and skip
	//if vnum is taken
	//Get the flags passed in
	var populated []Space
	var play Player
	if len(os.Args) > 1 {
		if os.Args[1] == "--init" {
			//TODO testing suite - one test will be randomly generating 10,000 Spaces
			//and seeing if the system can take it
			InitZoneSpaces("0-100", "The Void", "The absence of light is blinding.\nThree large telephone poles illuminate a small square.")
			InitZoneSpaces("100-150", "Midgaard", "I wonder what day is recycling day.\nEven the gods create trash.")
			populated = PopulateAreas()
			play = InitPlayer("FSM")
			addPfile(play)
			createMobiles("Noodles")
		}
		if os.Args[1] == "--client" {
			//Continue on
			populated = PopulateAreas()
			play = InitPlayer("FSM")
			savePfile(play)
			fmt.Println("In client loop")
		}
	} else {
		fmt.Println("Use --init to build and launch the world, --client to just connect.")
		os.Exit(1)
	}


	//Game loop
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		fmt.Println("Saving")
		savePfile(play)
		input := scanner.Text()
		//Save pfile first

		if input == "quit" {
			fmt.Println("Bai!")
			os.Exit(1)
		}
		if strings.HasPrefix(input, "ooc") {
			createChat(input[3:], play)
			showChat()
		}

		if input == "look" {
			fmt.Sprintf("Current room is ", play.CurrentRoom)
			showDesc(play.CurrentRoom)
		}
		if strings.Contains(input, "gen map") {
			play, populated = genMap(play, populated)
		}

		if strings.Contains(input, "open map") {
			fmt.Printf(mapPos)
			fmt.Print(populated[0].CoreBoard)
		}
		if strings.HasPrefix(input, "view from") {
			splitCommand := strings.Split(input, "from")
			stripped := strings.TrimSpace(splitCommand[1])
			vnumLook, err := strconv.Atoi(stripped)
			if err != nil {
				fmt.Println("Error converting a stripped string")
			}
			DescribeSpace(vnumLook, populated)
		}
		if strings.HasPrefix(input, "go to") {
			splitCommand := strings.Split(input, "to")
			stripped := strings.TrimSpace(splitCommand[1])
			inp, err := strconv.Atoi(stripped)
			if err != nil {
				fmt.Println("Error converting a stripped string")
			}
			play.CurrentRoom = populated[inp]
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
