package main

import (
  "fmt"
  "strings"
  "math/rand"
  "strconv"
  term "github.com/nsf/termbox-go"
)
 func reset() {
         term.Sync() // cosmestic purpose
 }

 func battle(target string, play Player, sounds [31]chan bool) {
   tarX, tarY := 0, 0
   sounds[12] <- true
         err := term.Init()
         if err != nil {
                 panic(err)
         }
         targetXY := strings.Split(target, ":")[1]
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
         TL, out := "", ""
         var damMsg []string
         fmt.Print("ESC button to quit")
         play.Battling = true
//         showCoreBoard(play)

 keyPressListenerLoop:
         for {
                switch ev := term.PollEvent(); ev.Type {
            case term.EventKey:
                    switch ev.Key {
                    case term.KeyEsc:
                            break keyPressListenerLoop
                    default:

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
                                 }else {
                                   play.Target = string(splitCPU[i][r])
                          	//				fmt.Print("\033["+strconv.Itoa(i+20)+";"+strconv.Itoa(r+54)+"H\033[48:2:175:0:150m"+string(splitCPU[play.TarY][play.TarX])+"\033[0m")
                                   play.TargetLong = string(splitCPU[play.TarY][play.TarX])

                                 }

                                 targ = fmt.Sprint("\033["+strconv.Itoa(i+20)+";"+strconv.Itoa(r+54)+"H\033[48:2:175:0:150m"+string(splitCPU[play.TarY][play.TarX])+"\033[0m")

                               }else {
                              //   fmt.Print("\033["+strconv.Itoa(i+20)+";"+strconv.Itoa(r+54)+"H"+string(splitCPU[i][r]))
                               }
                             }
                           }

                           play.Target = targ

                           if ev.Ch == 'e' {
                                 for i := 0;i < len(play.Classes);i++ {

                                   if play.Classes[i].Skills[i].Name == "overcharge" {
                                     //fmt.Print("\033[1;53HOVERCHARGING")
                                     if play.Won <= len(play.Fights.Oppose) {

                                       if strings.Contains(play.Target, "M") {
                                         x, y := play.TarX, play.TarY
                                         hit := false
                                         for bat := 0;bat < len(play.Fights.Oppose);bat++ {
                                           if play.Fights.Oppose[bat].X == x && play.Fights.Oppose[bat].Y == y {
                                             if play.Fights.Oppose[bat].MaxRezz <= 0 {
                                               damMsg = append(damMsg, fmt.Sprint("No use beating a dead corpse"))
                                               hit = false
                                               continue
                                             }

                                             if rand.Intn(10) > 4 {
                                               hit = true
                                             }else {
                                               hit = false
                                             }
                                             if hit {
                                               damage := play.Classes[i].Skills[i].Dam + rand.Intn(5)

                                               damageString := strconv.Itoa(damage)
                                               damMsg = append(damMsg, fmt.Sprint("\033[38:2:200:0:0mDid "+damageString+" damage to "+play.Fights.Oppose[play.Won].Name+"\033[0m"))
                                               play.Fights.Oppose[bat].MaxRezz -= damage
                                               if play.Fights.Oppose[bat].MaxRezz >= 0 {
                                                 sounds[17] <- true
                                                 continue
                                               }

                                               if play.Fights.Oppose[bat].MaxRezz < 0 {
                                                 play.Fights.Oppose[bat].Char = "*"
                                                 play.Won++
                                                 TL, out = determine(play)
                                                fmt.Printf(out+play.Target+TL)
                                                damMsg = append(damMsg, fmt.Sprintf("You slay %v!",TL))
                                                sounds[14] <- true
                                                continue
                                               }

                                             }else {
                                               damMsg = append(damMsg, fmt.Sprint("\033[38:2:150:0:150mYou don't manage to do any damage.\033[0m"))
                                               sounds[3] <- true
                                               continue
                                             }

                                           }


                                         }

                                       }

                                     }

                                   }

                                 }
                               }
         //		}else {
      //        reset()
        // 			clearCoreBoard(play)
         //		}
            out = ""
            out += showBattle(damMsg)
            out += showDesc(play.CurrentRoom)
         		out += DescribePlayer(play)
         		//chats = showChat(play)
         		out += showCoreBoard(play)

         		_, outln := showCoreMobs(play)
            out += outln

         		//ShowOoc(response, play)
            //updateChat(play, response)
            TL, outChar := determine(play)
            fmt.Print(out)
            fmt.Printf(outChar+play.Target+TL)

         		fmt.Printf("\033[51;0H")


      //			fmt.Print("Input co-ordinates in the form of aA aB aC etc..")
           //play, err := target(play, populated)
         }
         case term.EventError:
                 panic(ev.Err)
         }

         }
         play.Battling = false
 }
 func determine(play Player) (string, string) {
    TL := ""
    out := ""
    switch play.TargetLong {
    case "T":
     TL = "A Bejewelled Tiara"
     TL = fmt.Sprint("\033[19;53H\033[48;2;175;0;150m<<<"+TL+">>>\033[0m                      ")
    case "M":
     TL = "A Rabid Ferret"
     for bat := 0;bat < len(play.Fights.Oppose);bat++ {
       if play.Fights.Oppose[bat].X == play.TarX && play.Fights.Oppose[bat].Y == play.TarY {
         if strings.Contains(play.Fights.Oppose[bat].Char, "C") {
            out = fmt.Sprint("\033[19;53H\033[48;2;175;0;0m<<<DEAD\033[48;2;5;0;150m"+TL+"\033[48;2;175;0;0mDEAD>>>\033[0m                      ")
             TL = "The twisted remains of a rabid ferret"
             break
           }
       }else {
           out = fmt.Sprint("\033[19;53H\033[48;2;175;0;150m<<<"+TL+">>>\033[0m                      ")

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
