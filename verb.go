package main

import (
	"github.com/SolarLune/dngn"
	"fmt"
	"strconv"
	"strings"

  "math/rand"
	"context"
	"time"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func digDug(pos []int, play Player, digFrame [][]int, digNums string, digZone string, digNum int, populated []Space) (int, Space) {
	digVnumEnd := strings.Split(digNums, "-")[1]
	dg, digNum := initDigRoom(digFrame, digNums, digZone, play, digNum)
	play.CurrentRoom = dg
	for len(populated) <= digNum {
		populated = append(populated, dg)
	}
	populated[digNum] = dg
	dg.Vnum = digNum
	digFrame[pos[0]][pos[1]] = 8
	dg.ZonePos = dg.ZonePos[:0]
	dg.ZonePos = append(dg.ZonePos, pos[0])
	dg.ZonePos = append(dg.ZonePos, pos[1])
	fmt.Println("dug ", dg)
	drawDig(digFrame, dg.ZonePos)
	updateRoom(play, populated)
	fmt.Println("Dug ", digNum, " rooms of ", digVnumEnd)
	return digNum, dg
}


func AssembleComposeCel(chatMess Chat, row int) (string, int) {
	var cel string
	inWord := chatMess.Message
	wor := ""
	word := ""
	words := ""
	if len(inWord) > 68 {
		return "DONE COMPOSTING", 0
	}
	if len(inWord) > 28 && len(inWord) > 54 {
		wor += inWord[:28]
		word += inWord[28:54]
		words += inWord[54:]
		for i := len(words); i <= 28; i++ {
			words += " "
		}
	}
	if len(inWord) > 28 && len(inWord) < 54 {
		wor += inWord[:28]
		word += inWord[28:]
		for i := len(word); i <= 28; i++ {
			word += " "
		}
		words = "                            "

	}
	if len(inWord) <= 28 {
		wor = "                            "
		word += ""
		word += inWord
		for i := len(word); i <= 28; i++ {
			word += " "
		}
		words = "                            "
	}
	timeString := strings.Split(chatMess.Time.String(), " ")
	row++
	cel += fmt.Sprint("\033["+strconv.Itoa(row)+";180H\033[48;2;10;255;20m \033[48;2;10;10;20m", wor, "\033[48;2;10;255;20m \033[0m")
	row++
	cel += fmt.Sprint("\033["+strconv.Itoa(row)+";180H\033[48;2;10;255;20m \033[48;2;10;10;20m", word, "\033[48;2;10;255;20m \033[0m"+timeString[1])
	row++
	cel += fmt.Sprint("\033["+strconv.Itoa(row)+";180H\033[48;2;10;255;20m \033[48;2;10;10;20m", words, "\033[48;2;10;255;20m \033[0m"+timeString[0])
	row++
	namePlate := "                            "[len(chatMess.User.Name):]
	cel += fmt.Sprint("\033["+strconv.Itoa(row)+";180H\033[48;2;10;255;20m\033[38:2:50:0:50m@"+chatMess.User.Name+namePlate+"\033[48;2;10;255;20m \033[0m")

	return cel, row
	//	fmt.Println(cel)
}


func AssembleDescCel(room Space, row int) (string) {
	var cel string
	inWord := strings.Split(room.Desc, "\n")
	for len(strings.Split(room.Desc, "\n")) < 9 {
		room.Desc += "\n"
		inWord = strings.Split(room.Desc, "\n")
	}
	for len(inWord[0]) < 100 {
		inWord[0] += " "
	}
	row++
	cel += fmt.Sprint("\033["+strconv.Itoa(row+20)+";51H\033[48;2;100;5;100m \033[48;2;10;10;20m", inWord[0], "\033[48;2;100;5;100m \033[0m")
	for len(inWord[1]) < 100 {
		inWord[1] += " "
	}
	row++
	cel += fmt.Sprint("\033["+strconv.Itoa(row+20)+";51H\033[48;2;100;5;100m \033[48;2;10;10;20m", inWord[1], "\033[48;2;100;5;100m \033[0m")
	for len(inWord[2]) < 100 {
		inWord[2] += " "
	}
	row++
	cel += fmt.Sprint("\033["+strconv.Itoa(row+20)+";51H\033[48;2;100;5;100m \033[48;2;10;10;20m", inWord[2], "\033[48;2;100;5;100m \033[0m")
	for len(inWord[3]) < 100 {
		inWord[3] += " "
	}
	row++
	cel += fmt.Sprint("\033["+strconv.Itoa(row+20)+";51H\033[48;2;100;5;100m \033[48;2;10;10;20m", inWord[3], "\033[48;2;100;5;100m \033[0m")
	for len(inWord[4]) < 100 {
		inWord[4] += " "
	}
	row++
	cel += fmt.Sprint("\033["+strconv.Itoa(row+20)+";51H\033[48;2;100;5;100m \033[48;2;10;10;20m", inWord[4], "\033[48;2;100;5;100m \033[0m")
	for len(inWord[5]) < 100 {
		inWord[5] += " "
	}
	row++
	cel += fmt.Sprint("\033["+strconv.Itoa(row+20)+";51H\033[48;2;100;5;100m \033[48;2;10;10;20m", inWord[5], "\033[48;2;100;5;100m \033[0m")
	for len(inWord[6]) < 100 {
		inWord[6] += " "
	}
	row++
	cel += fmt.Sprint("\033["+strconv.Itoa(row+20)+";51H\033[48;2;100;5;100m \033[48;2;10;10;20m", inWord[6], "\033[48;2;100;5;100m \033[0m")
	for len(inWord[7]) < 100 {
		inWord[7] += " "
	}
	row++
	cel += fmt.Sprint("\033["+strconv.Itoa(row+20)+";51H\033[48;2;100;5;100m\033[38:2:50:0:50m@", inWord[7], "\033[48;2;100;5;100m \033[0m")

	return cel
}
func countKeys() {
  keys := "abcdefghijklmnopqrstuvwxyz0123456789"
  fmt.Println("\033[38:2:150:0:150m",len(keys),"in :",keys)

  keys = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
  fmt.Println("\033[38:2:175:0:150m",len(keys),"in :",keys)

  keys = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
  fmt.Println("\033[38:2:185:0:150m",len(keys),"in :",keys)
}

func goTo(dest int, play Player, populated []Space) (Player, []Space) {

	for i := 0;i < len(populated);i++ {
		if dest == populated[i].Vnum {
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
	return play, populated
}


func mergeMaps(source [][]int, dest [][]int) ([][]int) {
  for i := 0;i < len(source);i++ {
    for c := 0;c < len(source[i]);c++ {
      if source[i][c] == 1 {
        dest[i][c] = 1
      }
    }
  }
  return dest
}
func improvedTargeting(play Player, target string) (Player) {

	if strings.Contains(target, "|") {
		tarX, err := strconv.Atoi(strings.Split(target, "|")[0])
		if err != nil {
			panic(err)
		}
		tarY, err := strconv.Atoi(strings.Split(target, "|")[1])
		if err != nil {
			panic(err)
		}

		play.TarX = tarX
		play.TarY = tarY
	}else {
		play.OldX, play.OldY = play.TarX, play.TarY

		switch target {
		case "8":
			play.TarY -= 1
		case "2":
			play.TarY += 1
		case "4":
			play.TarX -= 1
		case "6":
			play.TarX += 1
		}

	}
	targ := ""
//	fmt.Print(play.CPU)
	splitCPU := strings.Split(play.CPU, "\n")
	CPU:
	for i := 0;i < len(splitCPU);i++ {
		for r := 0;r < len(splitCPU[i]);r++ {
			if play.TarX == r && play.TarY == i {
				if string(splitCPU[i][r]) == "%" {
					play.TarX, play.TarY = play.OldX, play.OldY
					targ = fmt.Sprint("\033["+strconv.Itoa(play.TarY+20)+";"+strconv.Itoa(play.TarX+54)+"H\033[48:2:175:0:150m"+string(splitCPU[play.TarY][play.TarX])+"\033[0m")
					break CPU
				}

				fmt.Print("\033["+strconv.Itoa(i+20)+";"+strconv.Itoa(r+54)+"H\033[48:2:175:0:150m"+string(splitCPU[play.TarY][play.TarX])+"\033[0m")
				play.TargetLong = string(splitCPU[play.TarY][play.TarX])

				targ = fmt.Sprint("\033["+strconv.Itoa(i+20)+";"+strconv.Itoa(r+54)+"H\033[48:2:175:0:150m"+string(splitCPU[play.TarY][play.TarX])+"\033[0m")

			}else {
				fmt.Print("\033["+strconv.Itoa(i+20)+";"+strconv.Itoa(r+54)+"H"+string(splitCPU[i][r]))
			}
		}
	}
	play.Target = targ
	return play
}


func genCoreBoard(play Player, populated []Space) (string, Player) {
	//Create a room map
	Room := dngn.NewRoom(126, 24)
	splits := rand.Intn(75)
	Room.GenerateBSP('%', 'D', splits)
//	_, err = collection.InsertOne(context.Background(), bson.M{"room":Room})
//	if err != nil {
//		panic(err)
//	}
  newValue := ""
  outVal := ""
//	fmt.Println("Generating and populating map")
	for i := 0;i < len(Room.Data);i++ {

	//				fmt.Println(populated[0].Room.Data[populated[0].Room.Width-1][i])
			value := string(Room.Data[i])
//      newValue = ""
			for s := 0;s < len(value);s++ {
				if string(value[s]) == " " {
					ChanceTreasure := "T"
					if rand.Intn(100) > 98 {
							newValue += ChanceTreasure
							continue
					}
					if rand.Intn(100) > 95 {
						ChanceMonster := "M"
						newValue += ChanceMonster
						continue
					}else {
						newValue += string(value[s])
					}
				}else {
					newValue += string(value[s])
				}

			}
      newValue += "\n"
    }
		play.CPU = newValue + "\n"

    play.PlainCoreBoard = newValue
    play.CoreBoard = newValue
    showCoreBoard(play)
    showChat(play)
    showDesc(play.CurrentRoom)
    time.Sleep(250*time.Millisecond)
    newValue = strings.ReplaceAll(newValue, "T", "\033[48;2;200;150;0mT\033[0m")

    play.CoreBoard = newValue
    showCoreBoard(play)
    showChat(play)
    showDesc(play.CurrentRoom)
    time.Sleep(250*time.Millisecond)
    newValue = strings.ReplaceAll(newValue, "M", "\033[48;2;200;50;50mM\033[0m")

    play.CoreBoard = newValue
    showCoreBoard(play)
    showChat(play)
    showDesc(play.CurrentRoom)
    time.Sleep(250*time.Millisecond)
		newValue = strings.ReplaceAll(newValue, "%", "\033[38;2;0;150;150m%\033[0m")

    play.CoreBoard = newValue
    showCoreBoard(play)
    showChat(play)
    showDesc(play.CurrentRoom)
    time.Sleep(250*time.Millisecond)
		newValue = strings.ReplaceAll(newValue, "D", "\033[48;2;200;150;150mD\033[0m")

    play.CoreBoard = newValue
    showCoreBoard(play)
    showChat(play)
    showDesc(play.CurrentRoom)
    time.Sleep(250*time.Millisecond)
		newValue = strings.ReplaceAll(newValue, " ", "\033[48:2:0:200:150m \033[0m")
		play.CoreBoard = newValue
    outVal += newValue + "\n"
		//fmt.Println(play.CPU)
	return outVal, play
}


func craftMob() Mobile {
	mob := InitMob()
	val := ""
  val += "\033[0;53H\033[48:2:120:120:0m                                                                              \033[0m"
  val += "\033[32;53H\033[48:2:120:120:0m                                                                              \033[0m"

  for i := 2;i < 32;i++ {
		if i == 2 || i == 31 {
			val += "\033["+strconv.Itoa(i)+";53H\033[48:2:120:120:0m \033[0m                                                                            \033[48:2:120:120:0m \033[0m"
		}else if i == 4 || i > 11 && i < 15 {
			if i == 4 {
				val += fmt.Sprint("\033[38:2:225:0:225m\033["+strconv.Itoa(i)+";53H\033[48:2:120:120:0m$                                                                             \033[0m")

			}else if i > 11 && i < 15 {
				val += fmt.Sprint("\033[38:2:225:0:225m\033["+strconv.Itoa(i)+";53H\033[48:2:120:120:0m&                                                                             \033[0m")

			}else {
				val += fmt.Sprint("\033[38:2:225:0:225m\033["+strconv.Itoa(i)+";53H\033[48:2:120:120:0m                                                                              \033[0m")

			}
		}else {
			val += "\033["+strconv.Itoa(i)+";53H\033[48:2:120:120:0m \033[0m                                                                            \033[48:2:120:120:0m \033[0m"

		}


  }
  fmt.Println(val)
	return mob
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
						"message":message, "time":time.Now(), "user":play })
	if err != nil {
		panic(err)
	}
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
