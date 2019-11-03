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
	"github.com/SolarLune/dngn"
	zmq "github.com/pebbe/zmq4"
)


type BroadcastPayload struct {
  Channel string
  Message string
  Game string
  Name string
}
type Broadcast struct {
    Event string
    Ref string
    Payload BroadcastPayload
}

type Descriptions struct {
	BATTLESPAM int
	ROOMDESC int
	PLAYERDESC int
	ROOMTITLE int
}
type Chat struct {
	User Player
	Message string
	Time time.Time
}
type Space struct{
	Room dngn.Room
	Vnums string
	Zone string
	ZonePos []int
	ZoneMap [][]int
	Vnum int
	Desc string
	Mobiles []int
	Items []int
	CoreBoard string
	Exits Exit
	Altered bool
}
type Exit struct {
	North int
	South int
	East int
	West int
	NorthWest int
	NorthEast int
	SouthWest int
	SouthEast int
	Up int
	Down int
}

type Player struct {
	Name string
	Title string
	Inventory []int
	Equipment []int
	CoreBoard string
	PlainCoreBoard string
	CurrentRoom Space
	PlayerHash string
	Target string
	TargetLong string
	TarX int
	TarY int
	OldX int
	OldY int
	CPU string

	MaxRezz int
	Rezz int
	Tech int

	Str int
	Int int
	Dex int
	Wis int
	Con int
	Cha int
}

type Mobile struct {
	Name string
	LongName string
	ItemSpawn []int
	Rep string
	MaxRezz int
	Rezz int
	Tech int
	Aggro int
	Align int
}



const (
	cmdPos = "\033[51;0H"
	mapPos = "\033[1;51H"
	descPos = "\033[0;50H"
	chatStart = "\033[38:2:200:50:50m{{=\033[38:2:150:50:150m"
	chatEnd = "\033[38:2:200:50:50m=}}"
	end = "\033[0m"

)



func main() {
	//TODO Get the Spaces that are already loaded in the database and skip
	//if vnum is taken
	//Get the flags passed in
	var populated []Space
	var mobiles []Mobile
	var play Player
	var hostname string
	var response *zmq.Socket
	chatBoxes := false
	grape := true
	//Make this relate to character level
	var dug []Space
	coreShow := false
	if len(os.Args) > 1 {
		if os.Args[1] == "--init" {
			//TODO testing suite - one test will be randomly generating 10,000 Spaces
			//and seeing if the system can take it
			descString := "The absence of light is blinding.\nThree large telephone poles illuminate a small square."
			for len(strings.Split(descString, "\n")) < 8 {
				descString += "\n"
			}
			InitZoneSpaces("0-5", "The Void", descString)
			descString = "I wonder what day is recycling day.\nEven the gods create trash."
			for len(strings.Split(descString, "\n")) < 8 {
				descString += "\n"
			}
			InitZoneSpaces("5-15", "Midgaard", descString)
			populated = PopulateAreas()
			play = InitPlayer("FSM")
			addPfile(play)
			createMobiles("Noodles")
			fmt.Println("\033[38:2:0:250:0mAll tests passed and world has been initialzed\n\033[0mYou may now start with --login.")
			os.Exit(1)
		}else if os.Args[1] == "--guest" {
			//Continue on
			populated = PopulateAreas()
			play = InitPlayer("Wallace")
			savePfile(play)
			fmt.Println("In client loop")
			fmt.Printf("\033[51;0H")
		}else if os.Args[1] == "--login" {
			//Continue on
			user, pword := LoginSC()

			populated = PopulateAreas()
			play = InitPlayer(user)
			//just hang on to the password for now
			fmt.Sprint(pword)
			savePfile(play)
			fmt.Println("In client loop")
			input := "go to 1"
			//this is pretty incomprehensible
			//TODO
			splitCommand := strings.Split(input, "to")
			stripped := strings.TrimSpace(splitCommand[1])
			inp, err := strconv.Atoi(stripped)
			if err != nil {
				fmt.Println("Error converting a stripped string")
			}
			for i := 0;i < len(populated);i++ {
				if inp == populated[i].Vnum {
					play.CurrentRoom = populated[i]
					fmt.Print(populated[i].Vnum, populated[i].Vnums, populated[i].Zone)
					showDesc(play.CurrentRoom)
					DescribePlayer(play)
					fmt.Printf("\033[0;0H\033[38:2:0:255:0mPASS\033[0m")
					break
				}else {
					fmt.Printf("\033[0;0H\033[38:2:255:0:0mERROR\033[0m")
				}
			}
			//log the character in

			response.Recv(0)
			_, err = response.Send(user + ":=:" + pword, 0)
			if err != nil {
				panic(err)
			}
			playBytes, err := response.RecvBytes(0)
			if err != nil {
				panic(err)
			}
			err = bson.Unmarshal(playBytes, &play)
			if err != nil || play.PlayerHash == "2" {
				panic(err)
			}
			fmt.Println(play.PlayerHash)
			fmt.Printf("\033[51;0H")
		}else if os.Args[1] == "--builder" {
			//Continue on
			populated = PopulateAreas()
			play = InitPlayer("FlyingSpaghettiMonster")
			savePfile(play)

			fmt.Println("Builder log-in")

			fmt.Printf("\033[51;0H")
		}	else if strings.Contains(os.Args[1], "--connect-core") {
				//TODO move these to after authentication
				user, pword := LoginSC()

				populated = PopulateAreas()
				mobiles = PopulateAreaMobiles()
				savePfile(play)

				fmt.Println("Core login procedure started")
				response, _ = zmq.NewSocket(zmq.REQ)

				defer response.Close()
				//Preferred way to connec
				hostname = "tcp://91.121.154.192:7777"
				err := response.Connect(hostname)
				fmt.Printf("\033[51;0H")
				user = strings.TrimSpace(user)
				pword = strings.TrimSpace(pword)
				_, err = response.Send(user+":=:"+pword, 0)
				if err != nil {
					panic(err)
				}
				playBytes, err := response.RecvBytes(0)
				if err != nil {
					panic(err)
				}
				err = bson.Unmarshal(playBytes, &play)
				if err != nil || play.PlayerHash == "2"{
					fmt.Print("\033[38:2:150:0:150mAuthorization failed\033[0m")
					os.Exit(1)
				}
			}else {
			fmt.Println("Unrecognized flag")
			os.Exit(1)
		}
 }else {
		fmt.Println("Use --init to build and launch the world, --user to just connect.")
		fmt.Println("--builder for a building session")
		os.Exit(1)
	}


	//Show the screen first off
	play.CurrentRoom = populated[1]
	showDesc(play.CurrentRoom)
	DescribePlayer(play)
	//showChat(play)

	//Game loop
	fmt.Println("#of mobiles:"+strconv.Itoa(len(mobiles)))
	firstDig := false
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		clearCmd()
		savePfile(play)
		input := scanner.Text()
		//Save pfile first
		save := false
		if strings.HasPrefix(input, "dig") {
			if strings.Split(input, " ")[1] == "new" {
				firstDig = true
			}else {
				firstDig = false
			}
			if firstDig {
				fmt.Println("Now specify the zone name and vnums required")
				fmt.Println("as in, \"dig zem 0 15\"")
				scanner.Scan()
				input = scanner.Text()
			}
			var digFrame [][]int
			for i := 0;i < 30;i++ {
				Frame := make([]int, 50)
				digFrame = append(digFrame, Frame)
			}

			fmt.Println("\033[38:2:255:0:0m", len(digFrame), "\033[0m")

			//Make a bar that fills with how many rooms you dig

			pos := make([]int, 2)

			if firstDig {
				pos[0] = 25
				pos[1] = 25
			}else {
				pos[0] = play.CurrentRoom.ZonePos[0]
				pos[1] = play.CurrentRoom.ZonePos[1]

			}

			if len(strings.Split(input, " ")) == 4 {
				digZone := strings.Split(input, " ")[1]
				digVnumStart := strings.Split(input, " ")[2]
				digVnumEnd := strings.Split(input, " ")[3]

				//Error was nil so start the digging protocol
				save = false
				dug = dug[:0]

				digNums := digVnumStart + "-" + digVnumEnd
				toDig := PopulateAreaBuild(digNums)
				for i := 0;i < len(toDig);i++ {
					populated = append(populated, toDig[i])

				}

				digNum, err := strconv.Atoi(digVnumStart)
				if err != nil {
					panic(err)
				}
				DIG:
				for scanner.Scan() {
					input = scanner.Text()
					inp, err := strconv.Atoi(input)
					if err != nil {
						fmt.Sprint("\033[0;0HAlphabetic code entry found")
						switch input {
						case "update zonemap":
							updateZoneMap(play, populated)
						case "edit desc":
							//desc
							//room has to exist before we edit it
							digDug(pos, play, digFrame, digNums, digZone, digNum, populated)
							//dignum shouldn't change because we're editing the same room

							play.CurrentRoom.Desc = ""
							fmt.Println("Enter the room's new description, enter for a new line, @ on a new line to end.")
							descScanner := bufio.NewScanner(os.Stdin)
							DESC:
							for descScanner.Scan() {
								if descScanner.Text() == "@" || len(strings.Split(populated[play.CurrentRoom.Vnum].Desc, "\n")) < 8 {
									if descScanner.Text() == "@" {
										for len(strings.Split(populated[play.CurrentRoom.Vnum].Desc, "\n")) < 8 {
											populated[play.CurrentRoom.Vnum].Desc += "\n"
										}
									}
									populated[play.CurrentRoom.Vnum].Desc = play.CurrentRoom.Desc
									break DESC
								}else {
									play.CurrentRoom.Desc += descScanner.Text() + "\n"
								}
							}
							client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
							if err != nil {
								panic(err)
							}
							ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
							err = client.Connect(ctx)
							if err != nil {
								panic(err)
							}
							filter := bson.M{"vnum": play.CurrentRoom.Vnum}
							collection := client.Database("zones").Collection("Spaces")
							update := bson.M{"$set": bson.M{"vnums":populated[play.CurrentRoom.Vnum].Vnums,
								 "desc":populated[play.CurrentRoom.Vnum].Desc,"exits": populated[play.CurrentRoom.Vnum].Exits,
									 "altered": true }}

							result, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
							if err != nil {
								panic(err)
							}
							fmt.Println("\033[38:2:255:0:0m", result, "\033[0m")
						case "edit title":
							//room title
						case "edit mobiles":
							//mobiles
						case "edit items":
							//items
						default:
							fmt.Println("I don't understand")
						}

						err = nil
					}
					//Set up the whole keypad for "digging"
					switch inp {
					case 1101:
						save = false
						break DIG
					case 1111:
						save = true
						break DIG
					case 1:
						//Sw

						if digFrame[pos[0]+1][pos[1]-1] != 1 {
							digFrame[pos[0]][pos[1]] = 1
							pos[0] += 1
							pos[1] -= 1
							digNum, play.CurrentRoom = digDug(pos, play, digFrame, digNums, digZone, digNum, populated)
							play.CurrentRoom.Vnum = digNum
						}
					case 2:
						//S
						if digFrame[pos[0]+1][pos[1]] != 1 {
							digFrame[pos[0]][pos[1]] = 1
							pos[0] += 1
							digNum, play.CurrentRoom = digDug(pos, play, digFrame, digNums, digZone, digNum, populated)
							play.CurrentRoom.Vnum = digNum
						}
					case 3:
						//Se
						if digFrame[pos[0]+1][pos[1]+1] != 1 {
							digFrame[pos[0]][pos[1]] = 1
							pos[0] += 1
							pos[1] += 1
							digNum, play.CurrentRoom = digDug(pos, play, digFrame, digNums, digZone, digNum, populated)
							play.CurrentRoom.Vnum = digNum
						}
					case 4:
						//W
						if digFrame[pos[0]][pos[1]-1] != 1 {
							digFrame[pos[0]][pos[1]] = 1
							pos[1] -= 1
							digNum, play.CurrentRoom = digDug(pos, play, digFrame, digNums, digZone, digNum, populated)
							play.CurrentRoom.Vnum = digNum
							}
					case 5:
						//TODO, make a selector for which level is shown
						//Down

						save = true
					case 6:
						//E
						if digFrame[pos[0]][pos[1]+1] != 1 {
							digFrame[pos[0]][pos[1]] = 1
							pos[1] += 1
							digNum, play.CurrentRoom = digDug(pos, play, digFrame, digNums, digZone, digNum, populated)
							play.CurrentRoom.Vnum = digNum
						}
					case 7:
						//Nw
						if digFrame[pos[0]-1][pos[1]-1] != 1 {
							digFrame[pos[0]][pos[1]] = 1
							pos[0] -= 1
							pos[1] -= 1
							digNum, play.CurrentRoom = digDug(pos, play, digFrame, digNums, digZone, digNum, populated)
							play.CurrentRoom.Vnum = digNum
						}
					case 8:
						//N
						if digFrame[pos[0]-1][pos[1]] != 1 {
							digFrame[pos[0]][pos[1]] = 1
							pos[0] -= 1
							digNum, play.CurrentRoom = digDug(pos, play, digFrame, digNums, digZone, digNum, populated)
							play.CurrentRoom.Vnum = digNum
						}
					case 9:
						//Ne
						if digFrame[pos[0]-1][pos[1]+1] != 1 {
							digFrame[pos[0]][pos[1]] = 1
							pos[0] -= 1
							pos[1] += 1
							digNum, play.CurrentRoom = digDug(pos, play, digFrame, digNums, digZone, digNum, populated)
							play.CurrentRoom.Vnum = digNum
						}
					default:
						if len(play.CurrentRoom.ZonePos) >= 2 {
							drawDig(digFrame, play.CurrentRoom.ZonePos)
						}
						fmt.Println("Dug ", digNum, " rooms of ", digVnumEnd)
					}
				}


			}
			if save {
				client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
				if err != nil {
					panic(err)
				}
				ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
				err = client.Connect(ctx)
				if err != nil {
					panic(err)
				}

				file, err := os.Create("dat/zone.bson")
				if err != nil {
					panic(err)
				}
				defer file.Close()
				writer := bufio.NewWriter(file)
				fmt.Println("\033[38:2:200:50:50mUpdating the zone with final map.\033[0m")
				updateZoneMap(play, populated)
				fmt.Println("Dumping the area list to dat/zone.bson")
				for i := 0;i < len(populated);i++ {
					marshalledBson, err := bson.Marshal(populated[i])
					if err != nil {
						panic(err)
					}
					writer.Write(marshalledBson)
					writer.Flush()
				}
			}

			}
		//COMMAND SECTION
		if input == "pew" {
			go playPew(1)
		}
		if strings.HasPrefix(input, "g ") {
			message := strings.Split(input, " ")[2]
			channel := strings.Split(input, " ")[1]
			response.Recv(0)
			fmt.Println("\033[38:2:0:150:150m[["+message+"]]\033[0m")
			_, err := response.Send(play.Name+"||UWU||"+channel+"||}}{{||"+message, 0)
			if err != nil {
				panic(err)
			}
		}
		if strings.Contains(input, "gvsub ") {

			channel := strings.Split(input, "gvsub ")[1]
			response.Recv(0)
			fmt.Println("Subscribing to "+channel)
			_, err := response.Send(play.Name+"+|+"+channel, 0)
			if err != nil {
				panic(err)
			}

		}
		if strings.Contains(input, "gvunsub ") {

			channel := strings.Split(input, "gvunsub ")[1]
			response.Recv(0)
			fmt.Println("Unsubscribing from "+channel)
			_, err := response.Send(play.Name+"-|-"+channel, 0)
			if err != nil {
				panic(err)
			}

		}
		if input == "logout" {
			response.Recv(0)
			fmt.Println(play.Name+"+==LOGOUT")
			_, err := response.Send(play.Name+"+==LOGOUT", 0)
			if err != nil {
				panic(err)
			}
			bye, err := response.Recv(0)
			if err != nil {
				panic(err)
			}
			fmt.Println(bye)
			fmt.Println("Have a great day!")
			time.Sleep(1*time.Second)
			os.Exit(1)
		}
		if strings.HasPrefix(input, "create") {
			name, password := strings.Split(input, " ")[1], strings.Split(input, " ")[2]
			response.Recv(0)
			_, err := response.Send(name + ":-:" + password, 0)
			if err != nil {
				panic(err)
			}
			play.PlayerHash, err = response.Recv(0)
			if err != nil {
				panic(err)
			}
			fmt.Println(play.PlayerHash)
		}
		if strings.HasPrefix(input, "login") {
			userPass := strings.Split(input, " ")
			user, pass := userPass[1], userPass[2]
			response.Recv(0)
			_, err := response.Send(user + ":=:" + pass, 0)
			if err != nil {
				panic(err)
			}
			playBytes, err := response.RecvBytes(0)
			if err != nil {
				panic(err)
			}
			err = bson.Unmarshal(playBytes, &play)
			if err != nil || play.Name == "" {
				fmt.Print("\033[38:2:150:0:150mAuthorization failed\033[0m")
				os.Exit(1)
			}
			fmt.Println(play.PlayerHash)
		}
		if strings.HasPrefix(input, "wizinit:") {
			fmt.Println("Sending init world command")
			pass := strings.Split(input, "--")[1]
			response.Recv(0)
			_, err := response.Send("init world:"+play.Name+"--"+pass, 0)
			if err != nil {
				panic(err)
			}
		}
		if input == "shutdown server" {
			fmt.Println("Sending shutdown signal")
			response.Recv(0)
			_, err := response.Send("+===shutdown===+", 0)
			if err != nil {
				panic(err)
			}
		}
		//secondary commands
		if strings.HasPrefix(input, "tc:") {
			TARG:
			for scanner.Scan() {
				inputTarg := scanner.Text()
				if strings.HasPrefix(input, "tc:") {
						targString := strings.Split(input, "tc:")[1]
						play = improvedTargeting(play, targString)
						input = ""
				}else if scanner.Text() == "out" {
					fmt.Println("Seeyah!")
					break TARG
				}else {
					play = improvedTargeting(play, inputTarg)
				}
				showDesc(play.CurrentRoom)
				//showChat(play)
				showCoreBoard(play)

		//		}else {
		//			clearCoreBoard(play)
		//		}
				TL := ""
				fmt.Printf(play.Target)
				switch play.TargetLong {
				case "T":
					TL = "A Bejewelled Tiara"
					TL = fmt.Sprint("\033[19;53H\033[48;2;175;0;150m<<<"+TL+">>>\033[0m                      ")
				case "M":
					TL = "A Rabid Ferret"
					TL = fmt.Sprint("\033[19;53H\033[48;2;175;0;150m<<<"+TL+">>>\033[0m                      ")
				case "D":
					TL = "A Large Steel Door"
					TL = fmt.Sprint("\033[19;53H\033[48;2;175;0;150m<<<"+TL+">>>\033[0m                      ")

				default:
					TL = fmt.Sprint("\033[19;53H\033[48;2;5;0;150m<<<"+TL+">>>\033[0m                        ")

				}
				fmt.Print(TL)
				fmt.Printf("\033[51;0H")

			}
//			fmt.Print("Input co-ordinates in the form of aA aB aC etc..")
			//play, err := target(play, populated)

		}
		if input == "show room vnum" {
			fmt.Print("\033[38;2;150;0;150mROOM VNUM :"+strconv.Itoa(play.CurrentRoom.Vnum)+"\033[0m")
		}
		if input == "dam rezz" {
			play.Rezz -= 5
		}
		if input == "dam tech" {
			play.Tech -= 6
		}
		if input == "heal" {
			play.Rezz = 17
			play.Tech = 17
		}
		if input == "show zone info" {
			fmt.Println("\033[38;2;150;0;150mZONE NAME :"+play.CurrentRoom.Zone+"\033[0m")
			fmt.Print("\033[38;2;150;0;150mZONE VNUMS :"+play.CurrentRoom.Vnums+"\033[0m")
		}
		if input == "edit desc"{

			play.CurrentRoom.Desc = ""
			fmt.Println("Enter the room's new description, enter for a new line, @ on a new line to end.")
			descScanner := bufio.NewScanner(os.Stdin)
			DESCREG:
			for descScanner.Scan() {
				if descScanner.Text() == "@" {
					if descScanner.Text() == "@" {
						for len(strings.Split(populated[play.CurrentRoom.Vnum].Desc, "\n")) < 8 {
							populated[play.CurrentRoom.Vnum].Desc += "\n"
						}
					}
					populated[play.CurrentRoom.Vnum].Desc = play.CurrentRoom.Desc
					break DESCREG
				}else {
					play.CurrentRoom.Desc += descScanner.Text() + "\n"
				}
			}

			client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
			if err != nil {
				panic(err)
			}
			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			err = client.Connect(ctx)
			if err != nil {
				panic(err)
			}
			filter := bson.M{"vnum": play.CurrentRoom.Vnum}
			collection := client.Database("zones").Collection("Spaces")
			update := bson.M{"$set": bson.M{"vnums":populated[play.CurrentRoom.Vnum].Vnums,
				 "desc":populated[play.CurrentRoom.Vnum].Desc,"exits": populated[play.CurrentRoom.Vnum].Exits,
					 "altered": true }}

			result, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
			if err != nil {
				panic(err)
			}
			fmt.Println("\033[38:2:255:0:0m", result, "\033[0m")
			populated = PopulateAreas()
		}



		if input == "quit" {
			fmt.Println("Bai!")
			zmq.AuthStop()
			os.Exit(1)
		}
		if strings.HasPrefix(input, "ooc") {
			input = strings.Replace(input, "ooc ", "+=+", 1)
			input = play.Name+input
			//createChat(input[3:], play)
			//todo
			response.Recv(0)
			_, err := response.Send(input, 0)
			if err != nil {
				panic(err)
			}
			chat, err := response.Recv(0)
			if err != nil {
				panic(err)
			}
			fmt.Printf(chat)
		}
		if input == "blit" {
			clearDirty()
		}
		if input == "show chat" {
			chatBoxes = true
		}
		if input == "hide chat" {
				chatBoxes = false
		}
		if input == "count keys" {
			countKeys()
			showDesc(play.CurrentRoom)
		}
		if strings.HasPrefix(input, "merge") {
			fmt.Println("Merging area zone map data")
			split := strings.Split(input, " ")
			sourceName, destName := split[1], split[2]
			var sourceDat [][]int
			var destDat [][]int
			for i := 0;i < len(populated);i++ {
				if populated[i].Zone == sourceName {
					sourceDat = populated[i].ZoneMap
				}
			}
			for i := 0;i < len(populated);i++ {
				if populated[i].Zone == destName {
					destDat = populated[i].ZoneMap
				}
			}
			zoneDat := mergeMaps(sourceDat, destDat)
			populated[play.CurrentRoom.Vnum].ZoneMap = zoneDat
			play.CurrentRoom.ZoneMap = zoneDat
			play.CurrentRoom.Zone = sourceName
			updateZoneMap(play, populated)
			play.CurrentRoom.Zone = destName
			updateZoneMap(play, populated)
		}
		if input == "update zonemap" {
			updateZoneMap(play, populated)
		}
		if input == "hide grape" {
			grape = false
			clearDirty()
			showDesc(play.CurrentRoom)
			DescribePlayer(play)
			//showChat(play)
			if coreShow {
				showCoreBoard(play)
			}
			if chatBoxes {
				showChat(play)
			}

			fmt.Printf("\033[51;0H")

		}
		if input == "show grape" {
			grape = true
		}
		if input == "hide chat" {
			chatBoxes = false
			clearDirty()
			showDesc(play.CurrentRoom)
			DescribePlayer(play)
			//showChat(play)
			if coreShow {
				showCoreBoard(play)
			}
			if chatBoxes {
				showChat(play)
			}

			fmt.Printf("\033[51;0H")
		}
		if input == "show chat" {
			chatBoxes = true
			clearDirty()
			showDesc(play.CurrentRoom)
			DescribePlayer(play)
			//showChat(play)
			if coreShow {
				showCoreBoard(play)
			}
			if chatBoxes {
				showChat(play)
			}

			fmt.Printf("\033[51;0H")
		}
		if input == "look" {
			fmt.Sprintf("Current room is ", play.CurrentRoom)
			showDesc(play.CurrentRoom)
			DescribePlayer(play)
		}
		if strings.Contains(input, "gen coreboard") {
			//TODO make this so one doesn't loose the
			//old coreboard, or convert it to xp, i dunno
			go playPew(2)
			play.CoreBoard, play = genCoreBoard(play, populated)
		}
		if strings.Contains(input, "open map") {
			//// TODO:
			//This
		}
		if strings.Contains(input, "craft mobile"){
			craftMob()
			for scanner.Scan() {
				if scanner.Text() != "exit" {
					continue
				}else {
					break
				}


			}
		}
		if strings.Contains(input, "unlock coreboard") {
			coreShow = false
		}
		if strings.Contains(input, "lock coreboard") {
//			fmt.Printf(mapPos)
				showCoreBoard(play)
				coreShow = true
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
			go playPew(1)
			play, populated = goTo(inp, play, populated)
		}
		if input == "score" {
			DescribePlayer(play)
		}
		if input == "updateChat" {
			updateChat(play, response)
		}

		//Reset the input to a standardized place
		showDesc(play.CurrentRoom)
		DescribePlayer(play)
		//showChat(play)
		if coreShow {
			showCoreBoard(play)
		}
		if chatBoxes {
			showChat(play)
		}
		if grape {
			updateChat(play, response)
		}
//		}else {
//			clearCoreBoard(play)
//		}
		fmt.Printf(play.Target)

		fmt.Printf("\033[51;0H")
	}
//	res, err := collection.InsertOne(context.Background(), bson.M{"Noun":"x"})
//	res, err = collection.InsertOne(context.Background(), bson.M{"Verb":"+"})
//	res, err = collection.InsertOne(context.Background(), bson.M{"ProperNoun":"y"})
}
