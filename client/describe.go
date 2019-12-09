package main

import (
  "os"
  "fmt"
  term "github.com/nsf/termbox-go"
  "time"
  "strconv"
  "math/rand"
  "strings"
)
func clearBigBroad() {
  for i := 21;i <= 45;i++ {
    pos := strconv.Itoa(i)
    fmt.Print("\033["+pos+";53H                                                                                    ")

  }
}
func makePlate(plate string, play Player) []string {
  var out []string
  count := 1
  for i := 0;i < len(play.Inventory);i++ {
      countString := strconv.Itoa(count)
//      fmt.Println(play.Inventory[i].Number)

      if play.Inventory[i].Number != 0 {
        out = append(out, fmt.Sprint("\033[",countString,";174H\033[48;2;10;255;20m \033[0m\033[48;2;10;10;20m  x", strconv.Itoa(play.Inventory[i].Number)+plate[len(play.Inventory[i].Item.Name)+3:]+play.Inventory[i].Item.Name, "\033[48;2;10;255;20m \033[0m"))
        count++
      }else {
        out = append(out, fmt.Sprint("\033[",countString,";174H\033[48;2;10;255;20m \033[0m\033[48;2;10;10;20m", plate, "\033[48;2;10;255;20m \033[0m"))
        count++
      }

  }
  return out
}
func describeInventory(play Player) string {
  cel := ""
  plateString := "                                                            "
  plate := makePlate(plateString, play)
  for i := 0;i < len(plate);i++ {
    cel += plate[i]
    fmt.Println(plate[i])
  }
  cel += fmt.Sprint("\033[20;174H\033[48;2;10;255;20m", plateString, " \033[0m")
  return cel
}
func describeEquipment(play Player) string {
  cel := ""
  plateString := "                                                 "
  plate := makeEQPlate(plateString, play)
  for i := 0;i < len(plate);i++ {
    cel += plate[i]
    fmt.Println(plate[i])
  }
  cel += fmt.Sprint("\033[20;1H\033[48;2;10;255;20m", plateString, " \033[0m")
  return cel
}
func makeEQPlate(plate string, play Player) []string {
  var out []string
  count := 1
  for i := 0;i < len(play.Equipped);i++ {
      countString := strconv.Itoa(count)

    //    out = append(out, fmt.Sprint("\033[",countString,";1H\033[48;2;10;255;20m \033[0m\033[48;2;10;10;20m  x", plate, "\033[48;2;10;255;20m \033[0m"))
      if play.Equipped[i].Item.LongName != "" {
        out = append(out, fmt.Sprint("\033[",countString,";1H\033[48;2;10;255;20m \033[0m\033[48;2;10;10;20m",play.Equipped[i].Item.LongName, plate[len(play.Equipped[i].Item.LongName):], "\033[48;2;10;255;20m \033[0m"))
      }else {
        out = append(out, fmt.Sprint("\033[",countString,";1H\033[48;2;10;255;20m \033[0m\033[48;2;10;10;20m", plate, "\033[48;2;10;255;20m \033[0m"))
      }

      count++


  }
  return out
}

func JackIn(in chan bool) error {
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
  fmt.Printf("\033[17;32H")
out := ""
row := 0
for i := 0;i < 52;i++ {
  for count := 0;count < 250;count++ {
    select {
    case notConn := <- in:
      clearDirty()
      if notConn == false {

        return nil
      }

    default:
          if rand.Intn(45) > 35 {
            randPosX := strconv.Itoa(rand.Intn(200))
            randPosY := strconv.Itoa(rand.Intn(52))
            out += "\033["+randPosY+";"+randPosX+"H\033[48:2:250:250:250m \033[0m"
          }else {
            randPosX := strconv.Itoa(rand.Intn(200))
            randPosY := strconv.Itoa(rand.Intn(52))
            out += "\033["+randPosY+";"+randPosX+"H\033[48:2:25:35:25m \033[0m"
          }
          row++

          time.Sleep(10*time.Millisecond)
          fmt.Print(out)

    }
  }

  }
  return nil
}


func LoginSC() (string, string){
       err := term.Init()
         if err != nil {
                 panic(err)
         }

         defer term.Close()

         fmt.Println("Enter any key to begin login sequence, ESC to cancel")

	userEntered := false
	pwdHide := "********************************************************************************************************"
	user := ""
	pword := ""
	escaped := false
	KEYPRESS:
         for {
                 switch ev := term.PollEvent(); ev.Type {
                 case term.EventKey:
                         switch ev.Key {
			case term.KeyBackspace2:
				fallthrough
			case term.KeyBackspace:
				if !userEntered {
					if len(user) - 1 > 0 {
						user = user[:len(user)-1]
					}else {
						user = ""
					}
					userLine := "________________"
	                                 fmt.Printf("\033[14;32H" + userLine + "\033[0m")
	                                 fmt.Printf("\033[14;32H" + user + "\033[0m")
				}else if userEntered {
					if len(pword)-1 > 0 {
						pword = pword[:len(pword)-1]
					}else {
						pword = ""
					}
					pwdLine := "________________"
					pwordMask := 0
					fmt.Printf("\033[17;32H" + pwdLine+ "\033[0m")
					pwordMask = len(pword)
					fmt.Printf("\033[17;32H" + pwdHide[:pwordMask] + "\033[0m")
				}
                         case term.KeyEsc:
				escaped = true
				break KEYPRESS
                        case term.KeyEnter:
				if !userEntered && len(user) > 3 {
					userEntered = true
					continue 
				}else if userEntered {
					return user, pword
				}
			default:
				
                                 // we only want to read a single character or one key pressed event
				  clearDirty()
				if !userEntered {
				  user += string(ev.Ch)
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
                                 }else if userEntered {
			    		pword += string(ev.Ch)
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
					fmt.Printf("\033[17;32H" + pwdHide[:len(pword)] + "\033[0m")


				}

                         }
                 case term.EventError:
                         panic(ev.Err)
                 }
         }
 if len(user) < 3 || len(pword) < 3 || escaped {
	term.Close()
	os.Exit(1)
 } 

  return "", ""
}
func clearDirty() {
  for i := 0;i < 255;i++ {
    fmt.Println("")
  }
}

func showBattle(damMsg []string) string {
	out := ""
  clear := false
  outClear := ""
  for i := 0;i < len(damMsg);i++ {
    if len(damMsg) > 17 {
        damMsg = damMsg[17:]
        //clearDirty()
        //reset()
        i = 0
        clear = true
    }
    if clear {
      for c := 0;c < 17;c++ {
        outClear += fmt.Sprint("\033["+strconv.Itoa(i)+";53H                                                                                                                                    ")
        if c == 16 {
            fmt.Print(outClear)
            clear = false
        }
      }

    }else {
      out += fmt.Sprint("\033["+strconv.Itoa(i+1)+";53H"+damMsg[i]+"\033["+strconv.Itoa(i+2)+";53H                                                                    ")

    }
  }
  clearCore()
  fmt.Print(outClear)
	return out
}

func clearCore() {
  for i := 0;i < 42;i++ {
    Y := strconv.Itoa(i)

    fmt.Print("\033["+Y+";0H                                                                                                                                                                                   ")
  }
}

func clearCmd() {
		fmt.Print("\033[52;0H                                                                                                                                                                                   ")
		fmt.Print("\033[53;0H                                                                                                                                                                                   ")
		fmt.Print("\033[54;0H                                                                                                                                                                                   ")
		fmt.Print("\033[55;0H                                                                                                                                                                                   ")
		fmt.Print("\033[56;0H                                                                                                                                                                                   ")
}



func showBattleSpam(spam []string) {
  count := 0
  if len(spam) >= 1 {
    for i := len(spam)-1;i > 0;i-- {
      if i > 22 {
        i = 0
        count++
        i += count
      }
      fmt.Print("\033["+strconv.Itoa(i+22)+";160H"+spam[i])
    }
  }
}

func showCoreMobs(play Player) (Player) {
  for i := 0;i < len(play.Fights.Oppose);i++ {
    if play.Fights.Oppose[i].Rezz <= 0 {
      play.Fights.Oppose[i].Name = play.Fights.Oppose[i].Corpse
    }
  }
  return play
}

func showCoreBoard(play Player) string {
  core := ""
	out := ""
  coreSplit := strings.Split(play.CoreBoard, "\n")
  for i := 0;i < len(coreSplit);i++ {
      core += fmt.Sprint("\033[",strconv.Itoa(i+22),";53H",coreSplit[i])
    }
    out = fmt.Sprint(core)
		return out
}
func clearCoreBoard(play Player) {
  coreSplit := strings.Split(play.CoreBoard, "\n")
  //This needs to be made dynamic for when we adjust the view. for now it's fine
  coreSpace := "                          "
  for i := 0;i < len(coreSplit);i++ {
    core := fmt.Sprint("\033[",strconv.Itoa(i+22),";53H ")

    fmt.Print(core+coreSpace)
  }
}
func DescribeSpace(vnum int, Spaces []Space) string {
	out := ""
	for i := 0; i < len(Spaces);i++ {
		if Spaces[i].Vnum == vnum {
			out += fmt.Sprint(Spaces[i].Zone)
			out += fmt.Sprint(Spaces[i].Desc)
      for m := 0;m < len(Spaces[i].MobilesInRoom);m++ {
          countString := strconv.Itoa(47+m)
          out += fmt.Sprint("\033["+countString+";53H"+Spaces[i].MobilesInRoom[m].LongName)
      }
    }
	}
	return out
}

func showPages(socBroadcasts []Broadcast, page int) string {
  numberBroadcasts := len(socBroadcasts)
  var out string
  col := 80
  curPage := page
  numPages := numberBroadcasts / 20
  for i := 1;i <= numPages;i++ {
    position := fmt.Sprint(strconv.Itoa(col+2*i))
    if i == curPage {
      out += fmt.Sprint("\033[48;2;200;200;10m\033[38;2;10;200;200m\033[21;"+position+"H"+strconv.Itoa(i)+"\033[0m")
      }else {
        out += fmt.Sprint("\033[48;2;10;200;10m\033[38;2;150;10;100m\033[21;"+position+"H"+strconv.Itoa(i)+"\033[0m")
      }

  }
  return out
}

func showDesc(room Space) string {
	out := ""
	splitPos := AssembleDescCel(room, 25)
	out += fmt.Sprint(splitPos)
	out += fmt.Sprint("\033[38:2:140:40:140m[[")
	if room.Exits.North != 0 {
		out += fmt.Sprint("\033[38:2:180:20:180mNorth")
	}
	if room.Exits.South != 0 {
		out += fmt.Sprint("\033[38:2:180:20:180mSouth")
	}
	if room.Exits.East != 0 {
		out += fmt.Sprint("\033[38:2:180:20:180mEast")
	}
	if room.Exits.West != 0 {
		out += fmt.Sprint("\033[38:2:180:20:180mWest")
	}
	if room.Exits.NorthWest != 0 {
		out += fmt.Sprint("\033[38:2:180:20:180mNorthWest")
	}
	if room.Exits.NorthEast != 0 {
		out += fmt.Sprint("\033[38:2:180:20:180mNorthEast")
	}
	if room.Exits.SouthWest != 0 {
		out += fmt.Sprint("\033[38:2:180:20:180mSouthWest")
	}
	if room.Exits.SouthEast != 0 {
		out += fmt.Sprint("\033[38:2:180:20:180mSouthEast")
	}
	if room.Exits.Up != 0 {
		out += fmt.Sprint("\033[38:2:180:20:180mUp")
	}
	if room.Exits.Down != 0 {
		out += fmt.Sprint("\033[38:2:180:20:180mDown")
	}

	out += fmt.Sprint("\033[38:2:140:40:140m]]\033[0m\033[0;0H")
//	if len(room.ZonePos) >= 2 {
//		out += drawDig(room.ZoneMap, room.ZonePos)
//	}
  for m := 0;m < len(room.MobilesInRoom);m++ {
      countString := strconv.Itoa(45-m)
      out += fmt.Sprint("\033["+countString+";53H\033[38:2:200:10:175m"+room.MobilesInRoom[m].LongName+"\033[0m")
  }
	return out
}
func drawDig(digFrame [][]int, zonePos []int) string {
	out := ""
	for i := 0;i < len(digFrame);i++ {
		out += fmt.Sprint("\033[48;2;10;255;20m \033[0m")
		for c := 0;c < len(digFrame[i]);c++ {
				prn := ""
				val := fmt.Sprint(digFrame[i][c])
				if i == zonePos[0] && c == zonePos[1] {
					prn = "8"
				}
				if prn == "8" {
					out += fmt.Sprint("\033[38:2:150:10:50m"+val+"\033[0m")
				}else if val == "1" || val == "8" {
					val = "1"
					out += fmt.Sprint("\033[38:2:50:10:50m"+val+"\033[0m")
				}else if c == 0 || c == len(digFrame[i])-1 || i == 0 || i == len(digFrame)-1{
          out += fmt.Sprint("\033[48;2;10;255;20m \033[0m")
        }else {
						out += fmt.Sprint(val)
				}
		}
		out += fmt.Sprintln("\033[48;2;10;255;20m \033[0m")
	}
	return out
}

func DescribePlayer(play Player) string {
	out := ""
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
  out += fmt.Sprint(techShow)
  out += fmt.Sprint(hp)
	out += fmt.Sprint("\033[28;0H")
//  out += listMyClasses(play)
  if len(play.ESlotSpell.Name) > 1 {
    out += fmt.Sprintln("\033[38:2:150:150:5m'e'\033[0m = "+play.ESlotSpell.Name)
  }else if len(play.ESlotSkill.Name) > 1 {
    out += fmt.Sprintln("\033[38:2:150:10:105m'e'\033[0m = "+play.ESlotSkill.Name)
  }else {
    out += fmt.Sprintln("")
  }
  if len(play.QSlotSpell.Name) > 1 {
    out += fmt.Sprintln("\033[38:2:150:150:5m'q'\033[0m = "+play.QSlotSpell.Name)
  }else if len(play.QSlotSkill.Name) > 1 {
    out += fmt.Sprintln("\033[38:2:150:10:105m'q'\033[0m = "+play.QSlotSkill.Name)
  }else {
    out += fmt.Sprintln("")
  }

  out += fmt.Sprintln("<<====", play.BankAccount.Amount, "====>>")
  out += fmt.Sprintln("\033[38:2:200:0:0mHas slain "+strconv.Itoa(play.Slain)+" monsters\033[0m")
  out += fmt.Sprintln("\033[38:2:150:150:0mHas found "+strconv.Itoa(play.Hoarded)+" treasures\033[0m")
	out += fmt.Sprintln("======================")
	out += fmt.Sprintln("\033[38:2:0:200:0mStrength     :\033[0m", play.Str)
	out += fmt.Sprintln("\033[38:2:0:200:0mIntelligence :\033[0m", play.Int)
	out += fmt.Sprintln("\033[38:2:0:200:0mDexterity    :\033[0m", play.Dex)
	out += fmt.Sprintln("\033[38:2:0:200:0mWisdom       :\033[0m", play.Wis)
	out += fmt.Sprintln("\033[38:2:0:200:0mConstitution :\033[0m", play.Con)
	out += fmt.Sprintln("\033[38:2:0:200:0mCharisma     :\033[0m", play.Cha)
	out += fmt.Sprintln("======================")
	return out
}
