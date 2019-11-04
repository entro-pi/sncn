package main

import (
  "os"
  "time"
  "fmt"
  "github.com/faiface/beep"
  "github.com/faiface/beep/wav"
  "github.com/faiface/beep/mp3"
  "github.com/faiface/beep/speaker"
)

func playPew(channelOne chan bool, channelTwo chan bool, channelThree chan bool) {

  fileName1 := "dat/sounds/49608__boilingsand__dialup-login-dec-2001-24-bit.wav"

  fileName2 := "dat/sounds/334914__robinhood76__06304-message-ding-1.wav"

  fileName3 := "dat/sounds/242341__ascap__metal-hit-medium-glass-bowl-8.mp3"

  f1, err := os.Open(fileName1)
  if err != nil {
    panic(err)
  }
  f2, err := os.Open(fileName2)
  if err != nil {
    panic(err)
  }
  f3, err := os.Open(fileName3)
  if err != nil {
    panic(err)
  }
  streamer1, format1, err := wav.Decode(f1)
  if err != nil {
    panic(err)
  }
  streamer2, format2, err := wav.Decode(f2)
  if err != nil {
    panic(err)
  }

  streamer3, format3, err := mp3.Decode(f3)
  if err != nil {
    fmt.Println(err)
  }

  sr1 := format2.SampleRate
  speaker.Init(sr1, sr1.N(time.Second/10))

  buffer1 := beep.NewBuffer(format1)
  buffer1.Append(streamer1)


  buffer2 := beep.NewBuffer(format2)
  buffer2.Append(streamer2)

  buffer3 := beep.NewBuffer(format3)
  buffer3.Append(streamer3)


  for {
  select {
  case <- channelOne:
    //streamer1.Close()
    noise1 := buffer1.Streamer(0, buffer1.Len())
    speaker.Play(noise1)
    //f1.Close()

  case <- channelTwo:
    //streamer2.Close()
    noise2 := buffer2.Streamer(0, buffer2.Len())
    speaker.Play(noise2)
    //f2.Close()
  case <- channelThree:
    //streamer2.Close()
    noise3 := buffer3.Streamer(0, buffer3.Len())
    speaker.Play(noise3)
    //f2.Close()

  }
}
streamer1.Close()
f1.Close()
streamer2.Close()
f2.Close()


}
