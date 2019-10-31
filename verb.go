package main

import (
	"github.com/SolarLune/dngn"
	"fmt"
	"strconv"
	"strings"
	"os"
	"bufio"
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
func target(play Player, populated []Space) error {

  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    out := ""
    topbar := " a b c d e f g h i j k l m n o p q r s t u v w "
    playPos := make([]int, 2)
    playPos[0], playPos[1] = 1, 1
    colPos := 25
    col := "z"
    sidebar := "A\nB\nC\nD\nE\nF\nG\nH\nI\nJ\nK\nL\nM\nN\nO\nP\nQ\nR\nS\nT\nU\nV\nW"
    row := "Z"
    //rowPos := 0
    input := scanner.Text()
		triggered := false
		for len(input) < 2 {
			input += "a"
			triggered = true
		}
		if triggered {
			input = "a"
			input += "A"
		}
    col = string(input[0])
    row = string(input[1])
    for i := 0;i < len(topbar);i++ {
      if string(topbar[i]) == col {
        out += fmt.Sprint("\033[48:2:200:0:0m"+string(topbar[i])+"\033[0m")
        colPos = i
      }else if i == 0 {
        out += fmt.Sprint("\033["+strconv.Itoa(i+20)+";51H"+string(topbar[i]))
      }else {
        out += string(topbar[i])
      }
    }
//    out += "\n"

	 	coreBoard := strings.Split(play.PlainCoreBoard, "\n")
    sidebarSplit := strings.Split(sidebar, "\n")
    for i := 0;i < len(sidebarSplit);i++ {
      out += fmt.Sprint("\033["+strconv.Itoa(i+21)+";51H\033[48:2:0:15:0m"+sidebarSplit[i])
      if sidebarSplit[i] == row {
      //	rowPos = i
        toOut := ""
        for c := 1;c < len(sidebar);c++ {
          if c == colPos - 1 || c == colPos + 1 {
            toOut += fmt.Sprint("\033[48:2:150:0:150m"+string(coreBoard[i][c])+"\033[0m")
          }else {
            toOut += fmt.Sprint("\033[48:2:0:200:0m"+string(coreBoard[i][c])+"\033[0m")
          }
        }
        out += toOut + "\n"
        continue
      }
      for c := 1;c < len(sidebar);c++ {
        if c == colPos {
          out += fmt.Sprint("\033[48:2:150:0:150m"+string(coreBoard[i][c])+"\033[0m")
        }else {
          out += fmt.Sprint(string(coreBoard[i][c]))
        }
      }
      out += "\n"
    }
    fmt.Print(out)
    fmt.Print("\033[51;1H")
    if scanner.Text() == "out" {
      fmt.Println("Seeyah!")
      return nil
    }

    }
    return nil
}
func genCoreBoard(play Player, populated []Space) (string, Player) {
	//Create a room map
	Room := dngn.NewRoom(128, 24)
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
			for s := 1;s < len(value);s++ {
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
		newValue = strings.ReplaceAll(newValue, " ", "\033[48;2;0;200;150m \033[0m")

    outVal += newValue + "\n"


	return outVal, play
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
