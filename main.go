package main

import (
	"bufio"
	"encoding/json"
	"os"
	"context"
	"time"
	"fmt"
	"strconv"
	"os/exec"
	"strings"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"

	zmq "github.com/pebbe/zmq4"
)



const (
	cmdPos = "\033[51;0H"
	mapPos = "\033[1;51H"
	descPos = "\033[0;50H"
	chatStart = "\033[38:2:200:50:50m{{=\033[38:2:150:50:150m"
	chatEnd = "\033[38:2:200:50:50m=}}"
	end = "\033[0m"
)



func main() {
	numSoundsnames, err := os.Open("dat/sounds")
	if err != nil {
		panic(err)
	}
	defer numSoundsnames.Close()
	soundFiles, err := numSoundsnames.Readdirnames(100)
	if err != nil {
		panic(err)
	}

	_ = len(soundFiles)
	//TODO Get the Spaces that are already loaded in the database and skip
	//if vnum is taken
	//Get the flags passed in
	var sounds [31]chan bool
	for i := 0;i < 30;i++ {
		sound := make(chan bool)
		sounds[i] = sound
	}
	go playSounds(sounds)
	//sounds[0] <- true
	var populated []Space
	var mobiles []Mobile
	//var chats int
	var chatsCurrent int
	var grapevines int
	var grapevinesCurrent int
	chatsCurrent = 0
	grapevinesCurrent = 0

//	chats = 0
	grapevines = 0
	var play Player
	var hostname string
	var response *zmq.Socket
	chatBoxes := true
	grape := true
	//Make this relate to character level
	var dug []Space
	play.CoreShow = false
	out := ""
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
			play = InitPlayer("FSM", "noodles")
			addPfile(play)
			createMobiles("Noodles")
			fmt.Println("\033[38:2:0:250:0mAll tests passed and world has been initialzed\n\033[0mYou may now start with --login.")
			os.Exit(1)
		}else if os.Args[1] == "--guest" {
			//Continue on
			populated = PopulateAreas()
			play = InitPlayer("Wallace", "gromit")
			savePfile(play)
			fmt.Println("In client loop")
			fmt.Printf("\033[51;0H")
		}else if os.Args[1] == "--login" {
			//Continue on
			user, pword := LoginSC()

			populated = PopulateAreas()
			play = InitPlayer(user, pword)
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
					out += showDesc(play.CurrentRoom)
					out += DescribePlayer(play)
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
			play = InitPlayer("FlyingSpaghettiMonster", "monster")
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
				hostname = "tcp://snowcrashnetwork.vineyard.haus:7777"
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
	connected := make(chan bool)

	if len(os.Args) >= 2 {
		if len(os.Args) > 2 && os.Args[2] == "--safe-mode"{
				play.Channels = play.Channels[0:]
					//noot noot
		}else {
//			play.Channels = append(play.Channels, "")
			play.Channels = append(play.Channels, "gossip")
			go JackIn(connected)
			sounds[29] <- true
		}

	}
	for i := 0;i < len(play.Channels);i++ {
		response.Recv(0)
		fmt.Println("Subscribing to "+play.Channels[i])
		_, err := response.Send(play.Name+"+|+"+play.Channels[i], 0)
		if err != nil {
			panic(err)
		}
		_, err = response.Recv(0)
		if err != nil {
			panic(err)
		}

		fmt.Println("ok")
		connected <- false
		sounds[9] <- true
		clearDirty()
		updateWho(play, true)
	}


	//Show the screen first off
	play.CurrentRoom = populated[1]
	out += showDesc(play.CurrentRoom)
	out += DescribePlayer(play)
	chats, outln := showChat(play)
	out += outln
	updateChat(play, response)
	out += ShowOoc(response, play)
	var ShowSoc bool

	var socBroadcasts []Broadcast
	//Game loop
	fmt.Println("#of mobiles:"+strconv.Itoa(len(mobiles)))
	firstDig := false
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		out = ""
		if chatsCurrent != chats {
		//	sounds[9] <- true
			chatsCurrent = chats
		}
		if grapevinesCurrent != grapevines {
		//	sounds[9] <- true
			grapevinesCurrent = grapevines
		}
		//clearCmd()
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
							out += drawDig(digFrame, play.CurrentRoom.ZonePos)
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
		if input == "save" {
				message := fmt.Sprint("++SAVE++"+play.PlayerHash)
				response.Recv(0)
				_, err = response.Send(message, 0)
				if err != nil {
					panic(err)
				}
				_, err = response.Recv(0)
				if err != nil {
					panic(err)
				}
				playerBytes, err := bson.Marshal(play)
				if err != nil {
					panic(err)
				}
				_, err = response.SendBytes(playerBytes, 0)
				if err != nil {
					panic(err)
				}

		}
		if input == "==INVALIDATE::>" {
			fmt.Println("INVALIDATING ALL SESSIONS")
		//	response.Recv(0)
			_, err = response.Send("::INVALIDATE::", 0)
			if err != nil {
				panic(err)
			}
		}
		if input == "CHECK SESSION" {
			fmt.Println(play.Session)
			response.Recv(0)
			_, err = response.Send(play.PlayerHash+"::CHECK::"+play.Session, 0)
			if err != nil {
				panic(err)
			}
			result, err := response.Recv(0)
			if err != nil {
				panic(err)
			}

			if result == "+__+SHUTDOWN+__+" {
				fmt.Println("\033[48:2:200:0:0mMIS-SESSION-TOKEN \nABORT\nABORT\nABORT")
				os.Exit(1)
			}

		}
		if strings.HasPrefix(input, "g:") {
			message := strings.Split(input, ":")[1]
			longMessage := ""
			channel := "gossip"
			fmt.Print("\033[35;53HComposing |+"+message+"+|\033[36;53H@ on a newline to end.\033[37;53H")
			count := 1
			for scanner.Scan() {
				if scanner.Text() == "@" {
					fmt.Print("Done composing!")
					break
				}
				longMessage += scanner.Text()
				fmt.Print("\033["+strconv.Itoa(35+count)+";53H")
			}

			response.Recv(0)
			fmt.Println("\033[38:2:0:150:150m[["+message+"]]\033[0m")
			_, err := response.Send(play.Name+"||UWU||"+channel+"||}}{{||"+message+"+++"+longMessage, 0)
			if err != nil {
				panic(err)
			}
			sounds[9] <- true
		}
		if input == "who" {
			who := fmt.Sprint(showWho(play))
			fmt.Printf("\033[38:2:175:0:150m"+who+"\033[0m")
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
			updateWho(play, false)
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
			playerBytes, err := response.RecvBytes(0)
			if err != nil {
				panic(err)
			}
			err = bson.Unmarshal(playerBytes, &play)
			if err != nil {
				panic(err)
			}
//			play = InitPlayer(name, password)
//			play.PlayerHash, err = response.Recv(0)
	//		if err != nil {
		//		panic(err)
		//	}
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
			play = battle(input, play, sounds)
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
			updateWho(play, false)
			fmt.Println("Have a great day!")
			time.Sleep(1*time.Second)
			os.Exit(1)
		}
		if input == "PAINT" {
			url := "dat/ASCIIpaint/index.html"
			cmd := exec.Command("xdg-open", url)
			cmd.Run()
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
			sounds[9] <- true
		}
		if input == "blit" {
			clearDirty()
		}
		if input == "show channels" {
			fmt.Print(play.Channels)
		}
		if input == "show chat" {
			chatBoxes = true
		}
		if input == "hide chat" {
				chatBoxes = false
		}
		if input == "count keys" {
			countKeys()
			out += showDesc(play.CurrentRoom)
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
			out += showDesc(play.CurrentRoom)
			out += DescribePlayer(play)
			//chats, out += showChat(play)
			if play.CoreShow {
				showCoreBoard(play)
			}
			if chatBoxes {
				_, outln := showChat(play)
				out += outln
			}
			sounds[2] <- true
			fmt.Printf("\033[51;0H")

		}
		if input == "show grape" {
			grape = true
			sounds[9] <- true
		}
		if input == "hide chat" {
			chatBoxes = false
			clearDirty()
			out += showDesc(play.CurrentRoom)
			out += DescribePlayer(play)
			//chats, out += showChat(play)
			if play.CoreShow {
				showCoreBoard(play)
			}
			if chatBoxes {
				_, outln := showChat(play)
				out += outln
			}
			sounds[9] <- true
			fmt.Printf("\033[51;0H")
		}
		if input == "show chat" {
			chatBoxes = true
			clearDirty()
			out += showDesc(play.CurrentRoom)
			out += DescribePlayer(play)
			//chats, out += showChat(play)
			if play.CoreShow {
				out += showCoreBoard(play)
			}
			if chatBoxes {
				_, outln := showChat(play)
				out += outln
			}
			sounds[9] <- true
			fmt.Printf("\033[51;0H")
		}
		if input == "report" {
			fmt.Print(play.Classes)
		}
		if input == "look" {
			fmt.Sprintf("Current room is ", play.CurrentRoom)
			out += showDesc(play.CurrentRoom)
			out += DescribePlayer(play)
		}
		if strings.Contains(input, "gen coreboard") {
			//TODO make this so one doesn't loose the
			//old coreboard, or convert it to xp, i dunno
			if len(strings.Split(input, "=")) > 1 {
				size := strings.Split(input, "=")[1]
				play.CoreBoard, play = genCoreBoard(size, play, populated)
				out += showCoreBoard(play)
				play.CoreShow = true
			}
		}
		if strings.Contains(input, "open map") {
			//// TODO:
			//This
		}
		if strings.Contains(input, "craft mobile"){
			craftMob(*scanner)

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
		if input == "SAVE ZONES" {
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
		if input == "addclass" {
			play = addClass(play)
		}
		if strings.HasPrefix(input, "go to") {
			splitCommand := strings.Split(input, "to")
			stripped := strings.TrimSpace(splitCommand[1])
			inp, err := strconv.Atoi(stripped)
			if err != nil {
				fmt.Println("Error converting a stripped string")
			}
			play, populated = goTo(inp, play, populated)
		}
		if input == "score" {
			out += DescribePlayer(play)
		}
		if input == "load profile picture" {
			play = importPhoto(play)
		}
		if input == "capture profile picture" {
			in := make(chan bool)
			out := make(chan string)
			frame := ""
			go clientLoops(in, out)
			newScan := bufio.NewScanner(os.Stdin)
			PHOTO:
			for newScan.Scan() {
				in <- true
				fmt.Printf("Press Enter to freeze frame\n@ on a newline to use.\nout on a newline to exit without saving.")
				select {
				case frame = <- out:
					fmt.Printf(frame)
					if newScan.Text() == "out" {
						fmt.Print("Quitting without saving!")
						in <- false
						break PHOTO
					}
					if newScan.Text() == "@" {
						play.Profile = frame
						in <- false
						break PHOTO
					}
					//nothing
				}


			}
		}
		if strings.HasPrefix(input, "soc") {
			if len(strings.Split(input, " ")) > 1 {
				response.Recv(0)
				//clear the selection
				for i := 0;i < len(socBroadcasts);i++ {
					socBroadcasts[i].Payload.Selected = false
				}

				fmt.Println("Sending --+--")
				_, err := response.Send(play.Name+"--+--", 0)
				_, err = response.Recv(0)
				if err != nil {
					panic(err)
				}
		//		out += string(result)
				grapevines = updateChat(play, response)
				fmt.Println("Sending ok")
				_, err = response.Send("--SELECT:"+strings.Split(input, " ")[1], 0)
				if err != nil {
					panic(err)
				}
				socBytes, err := response.RecvBytes(0)
				if err != nil {
					panic(err)
				}
				err = json.Unmarshal(socBytes, &socBroadcasts)
				if err != nil {
					panic(err)
				}

			}
//			fmt.Println(string(socBytes))
		}
		if input == "show soc" {
			ShowSoc = true
		}else if input == "hide soc" {
			ShowSoc = false
		}
		if strings.HasPrefix(input, "bs=") {
			numBS, err := strconv.Atoi(strings.Split(input, "=")[1])
			if err != nil {
				fmt.Println("Error, was that a number?")
			}
			for i := 0;i < len(socBroadcasts);i++ {
				socBroadcasts[i].Payload.Selected = false
				if i == numBS {
					socBroadcasts[i].Payload.Selected = true
				}
			}
		}
		if strings.Contains(input, "broadside=") {
			rowCol := strings.Split(input, "=")[1]
			row, err := strconv.Atoi(strings.Split(rowCol, ":")[0])
			col, err := strconv.Atoi(strings.Split(rowCol, ":")[1])
			if err != nil {
				panic(err)
			}
			var bs Broadcast
			bs.Payload.Message = "Kaboom!"
			bs.Payload.Channel = "BS"
			bs.Payload.Name = play.Name
			bs.Payload.Game = "snowcrash"

			broad := AssembleBroadside(bs, row, col)
			fmt.Printf(broad)
		}
		if strings.Contains(input, "pewpew") {
			if len(strings.Split(input, "pewpew ")) > 1 {
				numString := strings.Split(input, "pewpew ")[1]
				num, err := strconv.Atoi(numString)
				if err != nil {
					fmt.Println("Valid pews are 0-30")
					fmt.Println("Interesting sounds, 9, 17, 29")
				}
				sounds[num] <- true
			}
		}

		//Reset the input to a standardized place
		out += showDesc(play.CurrentRoom)
		out += DescribePlayer(play)
		//chats, out += showChat(play)
		if play.CoreShow {
			outln := ""
			out += showCoreBoard(play)
			play, outln = showCoreMobs(play)
			out += outln
		}

		if chatBoxes {
			ShowOoc(response, play)
//			chats, out += showChat(play)
		}
		if grape {
			grapevines = updateChat(play, response)
		}
//		}else {
//			clearCoreBoard(play)
//		}
		if ShowSoc {
			for i := 0;i < len(socBroadcasts);i++ {
	//			if socBroadcasts[i].Payload.Selected {
//					fmt.Println("\033[38:2:200:0:0mDOOOT\033[0m")
	//			}
				out += AssembleBroadside(socBroadcasts[i], socBroadcasts[i].Payload.Row, socBroadcasts[i].Payload.Col)
			}
			for i := 0;i < len(socBroadcasts);i++ {
				if socBroadcasts[i].Payload.Selected {
					out += AssembleBroadside(socBroadcasts[i], socBroadcasts[i].Payload.Row, socBroadcasts[i].Payload.Col)
				}
			}
			out += play.Profile
		}
		fmt.Print(out)
		//fmt.Printf(play.Target)

		fmt.Printf("\033[51;0H")
	}
		fmt.Sprint(chats)
//	res, err := collection.InsertOne(context.Background(), bson.M{"Noun":"x"})
//	res, err = collection.InsertOne(context.Background(), bson.M{"Verb":"+"})
//	res, err = collection.InsertOne(context.Background(), bson.M{"ProperNoun":"y"})
}
