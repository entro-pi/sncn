package main

import (
  "os"
  "time"
  "fmt"
  "strconv"
  "github.com/faiface/beep"
  "github.com/faiface/beep/wav"
  "github.com/faiface/beep/mp3"
  "github.com/faiface/beep/speaker"
)

func playPew(crashOne chan bool, channelOne chan bool, channelTwo chan bool, channelThree chan bool, channelFour chan bool, channelFive chan bool, channelSix chan bool, channelSeven chan bool, channelEight chan bool, channelNine chan bool, channelTen chan bool) {

  var beeps []string
  for i := 0;i < 10;i++ {
    if i == 0 {
      beeps = append(beeps, "dat/sounds/beep.wav")
    }else {
      beeps = append(beeps, "dat/sounds/beep"+strconv.Itoa(i)+".wav")
    }
  }

  crash := "dat/sounds/crash.wav"


  fileName1 := "dat/sounds/49608__boilingsand__dialup-login-dec-2001-24-bit.wav"

  fileName2 := "dat/sounds/334914__robinhood76__06304-message-ding-1.wav"

  fileName3 := "dat/sounds/242341__ascap__metal-hit-medium-glass-bowl-8.mp3"

  crashFile, err := os.Open(crash)
  if err != nil {
    panic(err)
  }
  f4, err := os.Open(beeps[3])
  if err != nil {
    panic(err)
  }
  f5, err := os.Open(beeps[4])
  if err != nil {
    panic(err)
  }
  f6, err := os.Open(beeps[5])
  if err != nil {
    panic(err)
  }
  f7, err := os.Open(beeps[6])
  if err != nil {
    panic(err)
  }
  f8, err := os.Open(beeps[7])
  if err != nil {
    panic(err)
  }
  f9, err := os.Open(beeps[8])
  if err != nil {
    panic(err)
  }
  f0, err := os.Open(beeps[9])
  if err != nil {
    panic(err)
  }
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

  crashStream, crashForm, err := wav.Decode(crashFile)
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

  streamer4, format4, err := wav.Decode(f4)
  if err != nil {
  fmt.Println(err)
  }

  streamer5, format5, err := wav.Decode(f5)
  if err != nil {
  fmt.Println(err)
  }

  streamer6, format6, err := wav.Decode(f6)
  if err != nil {
  fmt.Println(err)
  }

  streamer7, format7, err := wav.Decode(f7)
  if err != nil {
  fmt.Println(err)
  }

  streamer8, format8, err := wav.Decode(f8)
  if err != nil {
  fmt.Println(err)
  }

  streamer9, format9, err := wav.Decode(f9)
  if err != nil {
  fmt.Println(err)
  }

  streamer10, format10, err := wav.Decode(f0)
  if err != nil {
    fmt.Println(err)
  }

  sr1 := format2.SampleRate
  speaker.Init(sr1, sr1.N(time.Second/10))

  buffer1 := beep.NewBuffer(format1)
  buffer1.Append(streamer1)

  crashBuf := beep.NewBuffer(crashForm)
  crashBuf.Append(crashStream)

  buffer2 := beep.NewBuffer(format2)
  buffer2.Append(streamer2)

  buffer3 := beep.NewBuffer(format3)
  buffer3.Append(streamer3)


  buffer4 := beep.NewBuffer(format4)
  buffer4.Append(streamer4)


  buffer5 := beep.NewBuffer(format5)
  buffer5.Append(streamer5)


  buffer6 := beep.NewBuffer(format6)
  buffer6.Append(streamer6)


  buffer7 := beep.NewBuffer(format7)
  buffer7.Append(streamer7)


  buffer8 := beep.NewBuffer(format8)
  buffer8.Append(streamer8)


  buffer9 := beep.NewBuffer(format9)
  buffer9.Append(streamer9)


  buffer10 := beep.NewBuffer(format10)
  buffer10.Append(streamer10)


  for {
  select {
  case <- crashOne:
    crashNoise := crashBuf.Streamer(0, crashBuf.Len())
    speaker.Play(crashNoise)
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
  case <- channelFour:
    //streamer2.Close()
    noise4 := buffer4.Streamer(0, buffer4.Len())
    speaker.Play(noise4)
    //f2.Close()
  case <- channelFive:
    //streamer2.Close()
    noise5 := buffer5.Streamer(0, buffer5.Len())
    speaker.Play(noise5)
    //f2.Close()
  case <- channelSix:
    //streamer2.Close()
    noise6 := buffer6.Streamer(0, buffer6.Len())
    speaker.Play(noise6)
    //f2.Close()
  case <- channelSeven:
    //streamer2.Close()
    noise7 := buffer7.Streamer(0, buffer7.Len())
    speaker.Play(noise7)
    //f2.Close()
  case <- channelEight:
    //streamer2.Close()
    noise8 := buffer8.Streamer(0, buffer8.Len())
    speaker.Play(noise8)
    //f2.Close()
  case <- channelNine:
    //streamer2.Close()
    noise9 := buffer9.Streamer(0, buffer9.Len())
    speaker.Play(noise9)
    //f2.Close()
  case <- channelTen:
    //streamer2.Close()
    noise10 := buffer10.Streamer(0, buffer10.Len())
    speaker.Play(noise10)
    //f2.Close()

  }
}
streamer1.Close()
f1.Close()
streamer2.Close()
f2.Close()
streamer3.Close()

}
