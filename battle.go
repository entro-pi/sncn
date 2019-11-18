package main

import (
  "fmt"
  "strings"
  "math/rand"
  "time"
  "strconv"
  term "github.com/nsf/termbox-go"
)
 func reset() {
         term.Sync() // cosmestic purpose
 }

 func battle(play Player, sounds [31]chan bool) Player {
   target := "tc:1|1"
   targetXY := "1|1"
   tarX, tarY := 0, 0
   //sounds[12] <- true
         err := term.Init()
         if err != nil {
                 panic(err)
         }
         targetXY = strings.Split(target, ":")[1]
         if strings.Contains(targetXY, "|") {
       		tarX, err = strconv.Atoi(strings.Split(targetXY, "|")[0])
       		if err != nil {
       			panic(err)
       		}
       		tarY, err = strconv.Atoi(strings.Split(targetXY, "|")[1])
       		if err != nil {
       			panic(err)
       		}

       		play.TarX = tarX
       		play.TarY = tarY
    //   	}else {
      // 		play.OldX, play.OldY = play.TarX, play.TarY
        }
         defer term.Close()
//         TL, out := "", ""
         _, out := "", ""
         //var damMsg []string
         fmt.Print("ESC button to quit")
    //     fmt.Print(play.Fights)
         play.Battling = true
//         showCoreBoard(play)
         play.Won = 0
         play.Found = 0

           var battleSpamList []string
 keyPressListenerLoop:
         for {
           usedQSpellSkill := false
           usedESpellSkill := false

                switch ev := term.PollEvent(); ev.Type {
            case term.EventKey:
                    switch ev.Key {
                    case term.KeyEsc:
                            play.Won = 0
                            play.Found = 0
                            break keyPressListenerLoop
                    default:
                              clearBigBroad()
                              play = showCoreMobs(play)
                              showBattleSpam(battleSpamList)
                              damage := 0
                              techUsed := 0
                              battleSpam := ""
                          		play.OldX, play.OldY = play.TarX, play.TarY

                          		switch ev.Ch {
                          		case 'w':
                          			play.TarY -= 1
                          		case 's':
                          			play.TarY += 1
                          		case 'a':
                          			play.TarX -= 1
                          		case 'd':
                          			play.TarX += 1
                              case 'q':
                                usedQSpellSkill = true
                              case 'e':
                                usedESpellSkill = true
                              case 'c':
                                play = calculateWon(play)
                                if play.Fights.Won >= len(play.Fights.Oppose) {
                                  fmt.Printf("\033[38:2:175:150:0mSlew %v monsters, clearing the core.\033[0m", play.Fights.Won)
                                  play.Won += play.Fights.Won
                                  play.Fights.Won = 0

                                  fmt.Printf("\n\033[38:2:175:150:0mRe-joining social space.\033[0m")
                                  time.Sleep(5*time.Second)
                                  break keyPressListenerLoop
                                }else {
                                  fmt.Printf("Slew \033[38:2:200:0:0m%v\033[0m monsters.", play.Fights.Won)
                                  fmt.Printf("Gathered \033[38:2:175:150:0m%v\033[0m tiaras", play.Fights.Found)
                                }
                              }
                              playY := strconv.Itoa(play.TarY+22)
                              playX := strconv.Itoa(play.TarX+53)

                              prepend := fmt.Sprint("\033["+playY+";"+playX+"H")
                                  play.Target = prepend+"\033[48:2:150:0:150m \033[0m"
                                  splitBoard := strings.Split(play.PlainCoreBoard, "\n")
                                  for r := 0;r < len(splitBoard);r++ {
                                    for c := 0;c < len(splitBoard[r]);c++ {
                                      if play.TarY == r && play.TarX == c {
                                        play.TargetLong = string(splitBoard[r][c])
                                      }
                                    }
                                  }

                                  if play.TargetLong == "%" {
                                    play.TarY = play.OldY
                                    play.TarX = play.OldX
                                    playY = strconv.Itoa(play.TarY+22)
                                    playX = strconv.Itoa(play.TarX+53)

                                    prepend = fmt.Sprint("\033["+playY+";"+playX+"H")
                                    play.Target = prepend+"\033[48:2:150:0:150m \033[0m"
                                  }

                                //  out = ""
                              //    out += showBattle(damMsg)
                      //            out += showDesc(play.CurrentRoom)
                        //       		out += DescribePlayer(play)
                          //     		_, outln := showChat(play)
                            //      out += outln
                               		out += showCoreBoard(play)
                                  showBattleSpam(battleSpamList)
//                                  out += fmt.Sprint("\033[100;41H"+battleSpam)
                                  //outln := ""
                            //   		_, outln = showCoreMobs(play)
                                  //out += outln

                               		//ShowOoc(response, play)
                                  //updateChat(play, response)

                                  TL, outChar := determine(play)

                                  if usedESpellSkill && len(TL) > 1 {
                                    if play.ESlotSkill.Name != "" {
                                      for i := 0;i < len(play.Fights.Oppose);i++ {
                                        if play.Fights.Oppose[i].X == play.TarX && play.Fights.Oppose[i].Y == play.TarY {
                                            var blank Spell
                                            blank.Name = ""
                                            battleSpam, damage, techUsed = resolve(blank, play.ESlotSkill, play.Fights.Oppose[i], play.Tech)
                                            play.Fights.Oppose[i].Rezz -=  damage
                                            play.Fights.Oppose[i].MaxRezz -=  damage
                                            battleSpamList = append(battleSpamList, battleSpam)

                                            if damage > 0 {
                                              sounds[play.ESlotSkill.Sound] <- true
                                            }
                                        }
                                  }

                                    }
                                    if play.ESlotSpell.Name != "" {
                                      for i := 0;i < len(play.Fights.Oppose);i++ {
                                        if play.Fights.Oppose[i].X == play.TarX && play.Fights.Oppose[i].Y == play.TarY {

                                        var blank Skill
                                        blank.Name = ""
                                        battleSpam, damage, techUsed = resolve(play.ESlotSpell, blank, play.Fights.Oppose[i], play.Tech)
                                        play.Fights.Oppose[i].Rezz -=  damage
                                        play.Fights.Oppose[i].MaxRezz -=  damage
                                        battleSpamList = append(battleSpamList, battleSpam)

                                        if damage > 0 {
                                          sounds[play.ESlotSpell.Sound] <- true
                                        }
                                    }
                                      }
                                    }
                                  }
                                  if usedQSpellSkill && len(TL) > 1 {
                                    if play.QSlotSkill.Name != "" {
                                      sounds[play.QSlotSkill.Sound] <- true
                                      for i := 0;i < len(play.Fights.Oppose);i++ {
                                        if play.Fights.Oppose[i].X == play.TarX && play.Fights.Oppose[i].Y == play.TarY {
                                        var blank Spell
                                        blank.Name = ""
                                        battleSpam, damage, techUsed = resolve(blank, play.ESlotSkill, play.Fights.Oppose[i], play.Tech)
                                        play.Fights.Oppose[i].Rezz -=  damage
                                        play.Fights.Oppose[i].MaxRezz -=  damage
                                        battleSpamList = append(battleSpamList, battleSpam)

                                        if damage > 0 {
                                          sounds[play.QSlotSkill.Sound] <- true
                                        }
                                    }
                                    }
                                    }
                                    if play.QSlotSpell.Name != "" {
                                      for i := 0;i < len(play.Fights.Oppose);i++ {
                                        if play.Fights.Oppose[i].X == play.TarX && play.Fights.Oppose[i].Y == play.TarY {

                                        var blank Skill
                                        blank.Name = ""
                                        battleSpam, damage, techUsed = resolve(play.ESlotSpell, blank, play.Fights.Oppose[i], play.Tech)
                                        play.Fights.Oppose[i].Rezz -=  damage
                                        play.Fights.Oppose[i].MaxRezz -=  damage
//                                        fmt.Println(play.Fights.Oppose[i].MaxRezz)
                                        battleSpamList = append(battleSpamList, battleSpam)

                                        if damage > 0 {
                                          sounds[play.QSlotSpell.Sound] <- true
                                        }

                                      }
                                      }

                                    }
                                    fmt.Println(battleSpamList)
                                    play.Tech -= techUsed
                                    }
                                  fmt.Print(out)

                                  fmt.Printf(outChar+play.Target+TL)

                               		fmt.Printf("\033[51;0H")


                            //			fmt.Print("Input co-ordinates in the form of aA aB aC etc..")
                                 //play, err := target(play, populated)
}


         }

 }
 play.CoreShow = false
 play.Battling = false
 return play
}
func calculateWon(play Player) Player {
  play.Fights.Won = 0
  for i := 0;i < len(play.Fights.Oppose);i++ {
    if play.Fights.Oppose[i].Rezz <= 0 {
      play.Fights.Won++
    }
  }
  return play
}
func resolve(spell Spell, skill Skill, mob Mobile, tech int) (string, int, int) {
  spam := ""
  damage := 0
  if mob.Rezz <= 0 {
    spam = fmt.Sprintln("You have reduced \033[38:2:200:0:0m"+mob.Name+" to a rotting corpse\033[0m.")
    //fmt.Println(spam)
    return spam, 0, 0
  }
  if len(spell.Name) > 1 {
    if spell.TechUsage > tech {
      fmt.Printf("You don't have enough Tech to perform that action!")

      return "You don't have enough Tech to perform that action!", 0, 0
    }else {
      roll := rand.Intn(200)
      if roll >= mob.AC {
        damage = rand.Intn(spell.Dam)+2
        spam = fmt.Sprintln("Your \033[38:2:200:0:0m"+spell.Name+"\033[0m does \033[38:2:20:75:75m"+strconv.Itoa(damage)+"\033[0m to "+mob.Name)
        //fmt.Println(spam)
        return spam, damage, spell.TechUsage
      }else {
        spam = fmt.Sprintln("You fail to damage "+mob.Name)
      }
      return spam, damage, spell.TechUsage
    }
  }
  if len(skill.Name) > 1 {
    roll := rand.Intn(200)
    if roll > mob.AC {
      damage := rand.Intn(skill.Dam)+1
      spam = fmt.Sprintln("Your \033[38:2:200:0:0m"+skill.Name+"\033[0m does \033[38:2:20:75:75m"+strconv.Itoa(damage)+"\033[0m to "+mob.Name)
      //fmt.Println(spam)
      return spam, damage, 0
    }else {
      //fmt.Println("You fail to damage " + mob.Name)
      spam = fmt.Sprintln("You fail to damage "+mob.Name)
    }
    return spam, damage, 0
  }


  return "", 0, 0
}
func determine(play Player) (string, string) {
    TL := ""
    out := ""
    switch play.TargetLong {
    case "T":
     TL = "A Bejewelled Tiara"
     for i := 0;i < len(play.Fights.Treasure);i++ {
       if play.Fights.Treasure[i].X == play.TarX && play.Fights.Treasure[i].Y == play.TarY {
         TL = ""
         }else {
           TL = fmt.Sprint("\033[19;53H\033[48;2;175;0;150m<<<"+TL+">>>\033[0m                      ")
           }
     }
   case "M":
     fallthrough
  case "F":
    fallthrough
   case "B":
     fallthrough
   case "R":
      for i := 0;i < len(play.Fights.Oppose);i++ {
        if play.TarX == play.Fights.Oppose[i].X && play.TarY == play.Fights.Oppose[i].Y {
          TL = play.Fights.Oppose[i].Name
          if play.Fights.Oppose[i].Rezz <= 0 {
              TL = play.Fights.Oppose[i].Corpse
          }
          out = fmt.Sprint("\033[19;53H\033[48;2;175;0;150m<<<"+TL+">>>\033[0m                       ")

        }



     }
    case "D":
     TL = "A Large Steel Door"
     TL = fmt.Sprint("\033[19;53H\033[48;2;175;0;150m<<<"+TL+">>>\033[0m                      ")
    default:
     TL = fmt.Sprint("\033[19;53H\033[48;2;5;0;150m<<<"+TL+">>>\033[0m                        ")
    }
    return TL, out
}
