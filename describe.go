package main

import (
  "fmt"
  "os"
	"time"
	"context"
  "strconv"
  "bufio"
  "strings"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func LoginSC() (string, string){
  clearDirty()
  loginScanner := bufio.NewScanner(os.Stdin)
  user := ""
  fmt.Printf("\033[10;28H\033[0m")
  fmt.Printf("\033[11;28H \033[48;2;10;255;20m\033[38;2;10;10;255m         LOGIN         \033[0m")
  fmt.Printf("\033[12;28H\033[48;2;10;255;20m \033[48;2;10;10;20m                       \033[48;2;10;255;20m \033[0m")
  fmt.Printf("\033[13;28H\033[48;2;10;255;20m \033[48;2;10;10;20m   \033[38;2;10;200;150mUSER                \033[48;2;10;255;20m \033[0m")
  fmt.Printf("\033[14;28H\033[48;2;10;255;20m \033[48;2;10;10;20m   ________________    \033[48;2;10;255;20m \033[0m")
  fmt.Printf("\033[15;28H\033[48;2;10;255;20m \033[48;2;10;10;20m                       \033[48;2;10;255;20m \033[0m")
  fmt.Printf("\033[16;28H\033[48;2;10;255;20m \033[48;2;10;10;20m   \033[38;2;10;200;150mPASSWORD            \033[48;2;10;255;20m \033[0m")
  fmt.Printf("\033[17;28H\033[48;2;10;255;20m \033[48;2;10;10;20m   ________________    \033[48;2;10;255;20m \033[0m")
  fmt.Printf("\033[18;28H\033[48;2;10;255;20m \033[48;2;10;10;20m                       \033[48;2;10;255;20m \033[0m")
  fmt.Printf("\033[19;28H \033[48;2;10;255;20m                       \033[0m")
  fmt.Printf("\033[14;32H" + user + "\033[0m")

  loginScanner.Scan()

  user = loginScanner.Text()
  fmt.Printf("\033[17;32H")
  for {

		fmt.Printf("\033[10;28H\033[0m")
		fmt.Printf("\033[11;28H \033[48;2;10;255;20m\033[38;2;10;10;255m         LOGIN         \033[0m")
		fmt.Printf("\033[12;28H\033[48;2;10;255;20m \033[48;2;10;10;20m                       \033[48;2;10;255;20m \033[0m")
		fmt.Printf("\033[13;28H\033[48;2;10;255;20m \033[48;2;10;10;20m   \033[38;2;10;200;150mUSER                \033[48;2;10;255;20m \033[0m")
		fmt.Printf("\033[14;28H\033[48;2;10;255;20m \033[48;2;10;10;20m   ________________    \033[48;2;10;255;20m \033[0m")
		fmt.Printf("\033[15;28H\033[48;2;10;255;20m \033[48;2;10;10;20m                       \033[48;2;10;255;20m \033[0m")
		fmt.Printf("\033[16;28H\033[48;2;10;255;20m \033[48;2;10;10;20m   \033[38;2;10;200;150mPASSWORD            \033[48;2;10;255;20m \033[0m")
		fmt.Printf("\033[17;28H\033[48;2;10;255;20m \033[48;2;10;10;20m   ________________    \033[48;2;10;255;20m \033[0m")
		fmt.Printf("\033[18;28H\033[48;2;10;255;20m \033[48;2;10;10;20m                       \033[48;2;10;255;20m \033[0m")
		fmt.Printf("\033[19;28H \033[48;2;10;255;20m                       \033[0m")
		fmt.Printf("\033[14;32H" + user + "\033[0m")
    fmt.Printf("\033[17;32H")
		loginScanner.Scan()
    pword := loginScanner.Text()
    clearDirty()
    //Only use clearDirty at major intersections, it will cause flicker
		return user, pword

  }
  return "", ""
}
func clearDirty() {
  for i := 0;i < 255;i++ {
    fmt.Println("")
  }
}

func clearCmd() {
		fmt.Print(cmdPos+"                                                                                                                                                                                   ")
		fmt.Print("\033[52;0H                                                                                                                                                                                   ")
		fmt.Print("\033[53;0H                                                                                                                                                                                   ")
		fmt.Print("\033[54;0H                                                                                                                                                                                   ")
		fmt.Print("\033[55;0H                                                                                                                                                                                   ")
		fmt.Print("\033[56;0H                                                                                                                                                                                   ")
		fmt.Print(cmdPos)
}

func showCoreBoard(play Player) {
  coreSplit := strings.Split(play.CoreBoard, "\n")
  for i := 0;i < len(coreSplit);i++ {
    core := ""
//    if i == 0  || i == 1{
//      for len(core) < len(coreSplit[0]) {
        //core += fmt.Sprint(" ")
//      }
//      core += "\n"
//    }else {
      core = fmt.Sprint("\033[",strconv.Itoa(i+20),";52H",coreSplit[i])
//    }
    fmt.Print(core)
  }
}
func clearCoreBoard(play Player) {
  coreSplit := strings.Split(play.CoreBoard, "\n")
  //This needs to be made dynamic for when we adjust the view. for now it's fine
  coreSpace := "                          "
  for i := 0;i < len(coreSplit);i++ {
    core := fmt.Sprint("\033[",strconv.Itoa(i+20),";52H ")

    fmt.Print(core+coreSpace)
  }
}
func DescribeSpace(vnum int, Spaces []Space) {
	for i := 0; i < len(Spaces);i++ {
		if Spaces[i].Vnum == vnum {
			fmt.Println(Spaces[i].Zone)
			fmt.Println(Spaces[i].Desc)
		}
	}
}

func showDesc(room Space) {
	splitPos := AssembleDescCel(room, 25)
	fmt.Printf(splitPos)
	fmt.Printf("\033[38:2:140:40:140m[[")
	if room.Exits.North != 0 {
		fmt.Printf("\033[38:2:180:20:180mNorth")
	}
	if room.Exits.South != 0 {
		fmt.Printf("\033[38:2:180:20:180mSouth")
	}
	if room.Exits.East != 0 {
		fmt.Printf("\033[38:2:180:20:180mEast")
	}
	if room.Exits.West != 0 {
		fmt.Printf("\033[38:2:180:20:180mWest")
	}
	if room.Exits.NorthWest != 0 {
		fmt.Printf("\033[38:2:180:20:180mNorthWest")
	}
	if room.Exits.NorthEast != 0 {
		fmt.Printf("\033[38:2:180:20:180mNorthEast")
	}
	if room.Exits.SouthWest != 0 {
		fmt.Printf("\033[38:2:180:20:180mSouthWest")
	}
	if room.Exits.SouthEast != 0 {
		fmt.Printf("\033[38:2:180:20:180mSouthEast")
	}
	if room.Exits.Up != 0 {
		fmt.Printf("\033[38:2:180:20:180mUp")
	}
	if room.Exits.Down != 0 {
		fmt.Printf("\033[38:2:180:20:180mDown")
	}

	fmt.Printf("\033[38:2:140:40:140m]]\033[0m\033[0;0H")
	if len(room.ZonePos) >= 2 {
		drawDig(room.ZoneMap, room.ZonePos)
	}
}

func showChat(play Player) {
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
	var row int
	for mess.Next(context.Background()) {
		var chatMess Chat
		err := mess.Decode(&chatMess)
		if err != nil {
			panic(err)
		}
		chatPos := fmt.Sprintf("\033["+strconv.Itoa(count+3)+";180H")
		count++
		fmt.Printf(chatPos)
		if row >= 51 {
			row = 0
		}
		message, position := AssembleComposeCel(chatMess, row)
		row = position
		fmt.Printf(message)
//		fmt.Printf(chatStart)
//		fmt.Printf(chatMess.Message + " ")
//		fmt.Printf(chatEnd)
//  	fmt.Printf(end)

	}
}
func drawDig(digFrame [][]int, zonePos []int) {
	for i := 0;i < len(digFrame);i++ {
		for c := 0;c < len(digFrame[i]);c++ {
				prn := ""
				val := fmt.Sprint(digFrame[i][c])
				if i == zonePos[0] && c == zonePos[1] {
					prn = "8"
				}
				if prn == "8" {
					fmt.Printf("\033[38:2:150:10:50m"+val+"\033[0m")
				}else if val == "1" || val == "8" {
					val = "1"
					fmt.Printf("\033[38:2:50:10:50m"+val+"\033[0m")
				}else {
						fmt.Printf(val)
				}
		}
		fmt.Println("")
	}
}

func DescribePlayer(play Player) {

  ratio := ""
  count := 18
  for   rezz := 0;rezz < play.Rezz;rezz++ {

    ratio += "\033["+strconv.Itoa(count+30)+";25H\033[48:2:175:50:50m \033[0m\n"
    count--
  }
  for count > 0 {
      ratio += "\033["+strconv.Itoa(count+30)+";25H\033[48:2:15:50:50m \033[0m\n"

    count--
  }

  ratio += "\033[31;24H+++\n"
  ratio += "\033[49;24H+++"
  hp := ratio
  count = 18
  ratio = ""
  for tech := 0;tech < play.Tech;tech++ {
    ratio += "\033["+strconv.Itoa(count+30)+";31H\033[48:2:75:150:50m \033[0m\n"
    count--
  }
  for count > 0 {
      ratio += "\033["+strconv.Itoa(count+30)+";31H\033[48:2:15:50:50m \033[0m\n"
      count--
  }
  ratio += "\033[31;30H===\n"
  ratio += "\033[49;30H==="
  techShow := ratio
  fmt.Print(techShow)
  fmt.Print(hp)
	fmt.Printf("\033[40;0H")
	fmt.Println("======================")
	fmt.Println("\033[38:2:0:200:0mStrength     :\033[0m", play.Str)
	fmt.Println("\033[38:2:0:200:0mIntelligence :\033[0m", play.Int)
	fmt.Println("\033[38:2:0:200:0mDexterity    :\033[0m", play.Dex)
	fmt.Println("\033[38:2:0:200:0mWisdom       :\033[0m", play.Wis)
	fmt.Println("\033[38:2:0:200:0mConstitution :\033[0m", play.Con)
	fmt.Println("\033[38:2:0:200:0mCharisma     :\033[0m", play.Cha)
	fmt.Println("======================")
}
