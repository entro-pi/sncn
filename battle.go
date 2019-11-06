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

         fmt.Print("ESC button to quit")
         showCoreBoard(play)

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
                          //					fmt.Print("\033["+strconv.Itoa(i+20)+";"+strconv.Itoa(r+54)+"H\033[48:2:175:0:150m"+string(splitCPU[play.TarY][play.TarX])+"\033[0m")
                                   play.TargetLong = string(splitCPU[play.TarY][play.TarX])

                                 }

                                 targ = fmt.Sprint("\033["+strconv.Itoa(i+20)+";"+strconv.Itoa(r+54)+"H\033[48:2:175:0:150m"+string(splitCPU[play.TarY][play.TarX])+"\033[0m")

                               }else {
                                 fmt.Print("\033["+strconv.Itoa(i+20)+";"+strconv.Itoa(r+54)+"H"+string(splitCPU[i][r]))
                               }
                             }
                           }

                           play.Target = targ

                           if ev.Ch == 'e' {
                                 for i := 0;i < len(play.Classes);i++ {

                                   if play.Classes[i].Skills[i].Name == "overcharge" {
                                     fmt.Println("OVERCHARGING")
                                     if play.Won <= len(play.Fights.Oppose) {

                                       if strings.Contains(play.Target, "M") {
                                         x, y := play.TarX, play.TarY
                                         for bat := 0;bat < len(play.Fights.Oppose);bat++ {
                                           if play.Fights.Oppose[bat].X == x && play.Fights.Oppose[bat].Y == y {
                                             if rand.Intn(10) > 4 {
                                               damage := play.Classes[i].Skills[i].Dam + rand.Intn(5)

                                               damageString := strconv.Itoa(damage)
                                               fmt.Print("\033[52;5H\033[38:2:200:0:0mDid "+damageString+" damage to "+play.Fights.Oppose[play.Won].Name+"\033[0m")
                                               play.Fights.Oppose[bat].MaxRezz -= damage
                                               if play.Fights.Oppose[bat].MaxRezz >= 0 {
                                                 sounds[3] <- true
                                               }

                                               if play.Fights.Oppose[bat].MaxRezz < 0 {
                                                 play.Fights.Oppose[bat].Char = "*"
                                                 play.Won++
                                                 fmt.Println("Another one bites the dust!")
                                                sounds[14] <- true
                                               }
                                             }else {
                                               sounds[17] <- true
                                             }

                                           }


                                         }

                                       }

                                     }

                                   }

                                 }
                               }
         //		}else {
              reset()
         			clearCoreBoard(play)
         //		}
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

             showDesc(play.CurrentRoom)
         		DescribePlayer(play)
         		//chats = showChat(play)
         		showCoreBoard(play)
         		showCoreMobs(play)

         		//ShowOoc(response, play)
            //updateChat(play, response)

         		fmt.Printf(out+play.Target+TL)

         		fmt.Printf("\033[51;0H")


      //			fmt.Print("Input co-ordinates in the form of aA aB aC etc..")
           //play, err := target(play, populated)
         }
         case term.EventError:
                 panic(ev.Err)
         }

         }
 }
