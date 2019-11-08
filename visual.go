package main

import (
  "fmt"
  "bufio"
  "os"
  "gocv.io/x/gocv"
)

func importPhoto(play Player) Player {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Enter a full pathname in the type, \"/home/weasel/photo.png\"")
  fmt.Println("320x240 resolution works best")
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
                for row := 24; row > 0; row-- {
                        for column := 32; column > 0; column-- {
                                rS := p[2].GetUCharAt((row*10)-1, (column*10)-1)
                                gS := p[1].GetUCharAt((row*10)-1, (column*10)-1)
                                bS := p[0].GetUCharAt((row*10)-1, (column*10)-1)

              position := fmt.Sprint("\033[",row+2,";",column+2+75,"H")
                                word := fmt.Sprint(position,"\033[48;2;", rS, ";", gS, ";", bS, "m", "  ", "\033[0m")
                                wordSecondary = append(wordSecondary, word)

              position = fmt.Sprint("\033[",row+2,";",column+2,"H")
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
