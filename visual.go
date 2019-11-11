package main

import (
  "fmt"
  "math"
  "bufio"
  "os"
  "gocv.io/x/gocv"
  term "github.com/nsf/termbox-go"
)

func loadImages() []string {
  var result []string
  f, err := os.Open("dat/img/lorc")
  if err != nil {
    panic(err)
  }
  imgs, err := f.Readdirnames(789)
  if err != nil {
    panic(err)
  }
  for i := 0;i < len(imgs);i++ {

    result = append(result, loadPhotoLibrary("dat/img/lorc/"+imgs[i]))
  }

  return result

}


 func chooser(toChoose []string) string {
         err := term.Init()
         if err != nil {
                 panic(err)
         }

         defer term.Close()


          i := 0
 keyPressListenerLoop:
         for {
                 switch ev := term.PollEvent(); ev.Type {
                 case term.EventKey:
                         switch ev.Key {
                         case term.KeyEsc:
                                 break keyPressListenerLoop
                         case term.KeyArrowUp:
                                i++
                                fmt.Print(toChoose[i])
                                fmt.Print("user the arrow keys to choose a picture")
                         case term.KeyArrowDown:
                                i--
                                fmt.Print(toChoose[i])
                                fmt.Print("user the arrow keys to choose a picture")
                         case term.KeyArrowLeft:
                                i--
                                fmt.Print(toChoose[i])
                                fmt.Print("user the arrow keys to choose a picture")
                         case term.KeyArrowRight:
                                i++
                                fmt.Print(toChoose[i])
                                fmt.Print("user the arrow keys to choose a picture")
                         case term.KeyEnter:
                                 return toChoose[i]
                                 fmt.Print("user the arrow keys to choose a picture")
                         default:
                                 fmt.Print(toChoose[i])
                                 // we only want to read a single character or one key pressed event
                                 fmt.Print("user the arrow keys to choose a picture")
                         }
                 case term.EventError:
                         panic(ev.Err)
                 }
         }
         return toChoose[i]
 }

func loadPhotoLibrary(imgS string) string {
    		var frame string
    		img := gocv.IMRead(imgS, -1)

    		if img.Empty() {
    			fmt.Println("EMPTY FRAME")
    			img.Close()
    		}else {
          p := gocv.Split(img)
                var wordFinal []string
                for row := 38; row > 1; row-- {
                        for column := 37; column > 1; column-- {
                                rS := p[2].GetUCharAt((row*10)-1, (column*10)-1)
                                gS := p[1].GetUCharAt((row*10)-1, (column*10)-1)
                                bS := p[0].GetUCharAt((row*10)-1, (column*10)-1)
                                rowInt := math.Floor(float64(row)*0.65)+21

              position := fmt.Sprint("\033[",rowInt,";",column+51,"H")
                                word := fmt.Sprint(position,"\033[48;2;", rS, ";", gS, ";", bS, "m", "  ", "\033[0m")
                                wordFinal = append(wordFinal, word)
                        }
                }
          for i := len(wordFinal)-1;i > 0;i-- {
            frame += fmt.Sprintf(wordFinal[i])
          }

          img.Close()
      }
  return frame
}



func loadPhoto(play Player) Player {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("\033[21;90HEnter a full pathname in the type, \"/home/weasel/photo.png\"")
  fmt.Print("\033[22;90H400x400 resolution works best. \"done\" on a newline to choose\033[23;90H")
    for scanner.Scan() {
      if scanner.Text() == "done" {
        return play
      }

    		var frame string
    		img := gocv.IMRead(scanner.Text(), -1)

    		if img.Empty() {
    			fmt.Println("EMPTY FRAME")
    			img.Close()
    		}else {
          p := gocv.Split(img)
                var wordFinal []string
          var wordSecondary []string
                for row := 38; row > 1; row-- {
                        for column := 37; column > 1; column-- {
                                rS := p[2].GetUCharAt((row*10)-1, (column*10)-1)
                                gS := p[1].GetUCharAt((row*10)-1, (column*10)-1)
                                bS := p[0].GetUCharAt((row*10)-1, (column*10)-1)
                                rowInt := math.Floor(float64(row)*0.65)+21
              position := fmt.Sprint("\033[",rowInt,";",column+75,"H")
                                word := fmt.Sprint(position,"\033[48;2;", rS, ";", gS, ";", bS, "m", "  ", "\033[0m")
                                wordSecondary = append(wordSecondary, word)

              position = fmt.Sprint("\033[",rowInt,";",column+51,"H")
                                word = fmt.Sprint(position,"\033[48;2;", rS, ";", gS, ";", bS, "m", "  ", "\033[0m")
                                wordFinal = append(wordFinal, word)
                        }
                }
          var frameSecond string
          for i := len(wordFinal)-1;i > 0;i-- {
            frame += fmt.Sprintf(wordFinal[i])
            frameSecond += fmt.Sprintf(wordSecondary[i])
          }
          play.Profile = frame
          fmt.Print(play.Profile)
          img.Close()
      }
}
  return play
}


func clientLoops(in chan bool, out chan string) {
  camera, err := gocv.OpenVideoCapture(0)
  if err != nil {
    panic(err)
  }
	client := true
	CLIENT:
	for camera.IsOpened(){
    select {
    case val := <- in :
      if val == false {
        break CLIENT
      }else {
        img := gocv.NewMat()
    		var frame string
    		if ok := camera.Read(&img); !ok {
    			fmt.Println("Camera is closed")
    		}
    		if img.Empty() {
    			fmt.Println("EMPTY FRAME")
    			img.Close()
    			continue CLIENT
    		}else {
    			p := gocv.Split(img)
    		        var wordFinal []string
    			var wordSecondary []string
    		        for row := 24; row > 0; row-- {
    		                for column := 32; column > 0; column-- {
    		                        rS := p[2].GetUCharAt((row*10)-1, (column*10)-1)
    		                        gS := p[1].GetUCharAt((row*10)-1, (column*10)-1)
    		                        bS := p[0].GetUCharAt((row*10)-1, (column*10)-1)

    					position := fmt.Sprint("\033[",row+2,";",column+2+75,"H")
    		                        word := fmt.Sprint(position,"\033[48;2;", rS, ";", gS, ";", bS, "m", "==", "\033[0m")
    		                        wordSecondary = append(wordSecondary, word)

    					position = fmt.Sprint("\033[",row+2,";",column+2,"H")
    		                        word = fmt.Sprint(position,"\033[48;2;", rS, ";", gS, ";", bS, "m", "==", "\033[0m")
    		                        wordFinal = append(wordFinal, word)
    		                }
    		        }
    			var frameSecond string
    			for i := len(wordFinal)-1;i > 0;i-- {
    				frame += fmt.Sprintf(wordFinal[i])
    				frameSecond += fmt.Sprintf(wordSecondary[i])
    			}
    			if client {
    				//do client stuff
    				out <- frame
    			}
    			img.Close()
    		}
      }
      }
	}
  camera.Close()

}
