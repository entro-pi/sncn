package main

import (
  "os"
  "time"
  "github.com/faiface/beep"
  "github.com/faiface/beep/wav"
  "github.com/faiface/beep/speaker"
)

func playPew(sound int) {
  fileName := ""
  if sound == 1 {
    fileName = "dat/sounds/482273__seanporio__lazer-pluck-2-mod.wav"
  }else if sound == 2 {
    fileName = "dat/sounds/482273__seanporio__lazer-pluck-2-slow-mod.wav"
  }
  f, err := os.Open(fileName)
  if err != nil {
    panic(err)
  }
  defer f.Close()
  streamer, format, err := wav.Decode(f)
  if err != nil {
    panic(err)
  }
  defer streamer.Close()

  sr := format.SampleRate
  speaker.Init(sr, sr.N(time.Second/10))

  	done := make(chan bool)
  	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
  		done <- true
  	})))

  	<-done
}
