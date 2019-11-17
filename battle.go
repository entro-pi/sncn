package main

import (
  "fmt"
  "strings"
//  "math/rand"
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
         var damMsg []string
         fmt.Print("ESC button to quit")
         play.Battling = true
//         showCoreBoard(play)
         play.Won = 0
         play.Found = 0

 keyPressListenerLoop:
         for {
                switch ev := term.PollEvent(); ev.Type {
            case term.EventKey:
                    switch ev.Key {
                    case term.KeyEsc:
                            play.Won = 0
                            play.Found = 0
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
                              case 'e':
                                
                              case 'c':
                                if play.Won >= len(play.Fights.Oppose) {
                                  fmt.Printf("\033[38:2:175:150:0mSlew %v monsters, clearing the core.\033[0m", play.Won)
                                  play.Won = 0
                                  fmt.Printf("\n\033[38:2:175:150:0mRe-joining social space.\033[0m")
                                  time.Sleep(5*time.Second)
                                  break keyPressListenerLoop
                                }else {
                                  fmt.Printf("Slew \033[38:2:200:0:0m%v\033[0m monsters.", play.Won)
                                  fmt.Printf("Gathered \033[38:2:175:150:0m%v\033[0m tiaras", play.Found)
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

                                  out = ""
                                  out += showBattle(damMsg)
                      //            out += showDesc(play.CurrentRoom)
                        //       		out += DescribePlayer(play)
                          //     		_, outln := showChat(play)
                            //      out += outln
                               		out += showCoreBoard(play)
                                  outln := ""
                               		_, outln = showCoreMobs(play)
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


         }

 }
 play.CoreShow = false
 play.Battling = false
 return play
}

func determine(play Player) (string, string) {
    TL := ""
    out := ""
    switch play.TargetLong {
    case "T":
     TL = "A Bejewelled Tiara"
     for i := 0;i < len(play.Fights.Treasure);i++ {
       if play.Fights.Treasure[i].X == play.TarX && play.Fights.Treasure[i].Y == play.TarY && play.Fights.Treasure[i].Owned == true {
         TL = ""
         }else {
           TL = fmt.Sprint("\033[19;53H\033[48;2;175;0;150m<<<"+TL+">>>\033[0m                      ")
           }
     }
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
