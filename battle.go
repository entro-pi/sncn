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

 func battle(play Player, sounds [31]chan bool) {
         err := term.Init()
         if err != nil {
                 panic(err)
         }

         defer term.Close()

         fmt.Print("ESC button to quit")

 keyPressListenerLoop:
         for {
                switch ev := term.PollEvent(); ev.Type {
            case term.EventKey:
                    switch ev.Key {
                    case term.KeyEsc:
                            break keyPressListenerLoop
                    case term.KeyF1:
                            reset()
                            fmt.Println("F1 pressed")
                    case term.KeyF2:
                            reset()
                            fmt.Println("F2 pressed")
                    case term.KeyF3:
                            reset()
                            fmt.Println("F3 pressed")
                    case term.KeyF4:
                            reset()
                            fmt.Println("F4 pressed")
                    case term.KeyF5:
                            reset()
                            fmt.Println("F5 pressed")
                    case term.KeyF6:
                            reset()
                            fmt.Println("F6 pressed")
                    case term.KeyF7:
                            reset()
                            fmt.Println("F7 pressed")
                    case term.KeyF8:
                            reset()
                            fmt.Println("F8 pressed")
                    case term.KeyF9:
                            reset()
                            fmt.Println("F9 pressed")
                    case term.KeyF10:
                            reset()
                            fmt.Println("F10 pressed")
                    case term.KeyF11:
                            reset()
                            fmt.Println("F11 pressed")
                    case term.KeyF12:
                            reset()
                            fmt.Println("F12 pressed")
                    case term.KeyInsert:
                            reset()
                            fmt.Println("Insert pressed")
                    case term.KeyDelete:
                            reset()
                            fmt.Println("Delete pressed")
                    case term.KeyHome:
                            reset()
                            fmt.Println("Home pressed")
                    case term.KeyEnd:
                            reset()
                            fmt.Println("End pressed")
                    case term.KeyPgup:
                            reset()
                            fmt.Println("Page Up pressed")
                    case term.KeyPgdn:
                            reset()
                            fmt.Println("Page Down pressed")
                    case term.KeyArrowUp:
                            reset()
                            fmt.Println("Arrow Up pressed")
                    case term.KeyArrowDown:
                            reset()
                            fmt.Println("Arrow Down pressed")
                    case term.KeyArrowLeft:
                            reset()
                            fmt.Println("Arrow Left pressed")
                    case term.KeyArrowRight:
                            reset()
                            fmt.Println("Arrow Right pressed")
                    case term.KeySpace:
                            reset()
                            fmt.Println("Space pressed")
                    case term.KeyBackspace:
                            reset()
                            fmt.Println("Backspace pressed")
                    case term.KeyEnter:
                            reset()
                            fmt.Println("Enter pressed")
                    case term.KeyTab:
                            reset()
                            fmt.Println("Tab pressed")

                    default:
                            // we only want to read a single character or one key pressed event
                            reset()
                            fmt.Println("ASCII : ", ev.Ch)

                           showDesc(play.CurrentRoom)
                           showChat(play)
                           showCoreBoard(play)
                           showCoreMobs(play)
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
                                               sounds[9] <- true

                                               if play.Fights.Oppose[bat].MaxRezz <= 0 {
                                                 play.Fights.Oppose[bat].Char = "*"
                                                 play.Won++
                                                 fmt.Println("Another one bites the dust!")

                                               }
                                             }else {
                                               sounds[9] <- true
                                             }

                                           }


                                         }

                                       }

                                     }

                                   }

                                 }
                               }
         //		}else {
         //			clearCoreBoard(play)
         //		}
             TL := ""
             out := ""
             fmt.Printf(play.Target)
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
             fmt.Print(TL)
             fmt.Print(out)
             fmt.Printf("\033[51;0H")


      //			fmt.Print("Input co-ordinates in the form of aA aB aC etc..")
           //play, err := target(play, populated)
         }
         case term.EventError:
                 panic(ev.Err)
         }

         }
 }
