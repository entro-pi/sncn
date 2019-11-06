package main

import (
  "os"
  "time"
  "github.com/faiface/beep"
  "github.com/faiface/beep/wav"
  "github.com/faiface/beep/speaker"
)


func playSounds(sounds [31]chan bool){
  f, err := os.Open("dat/sounds/")
  if err != nil {
    panic(err)
  }
  beeps, err := f.Readdirnames(31)
//  fmt.Println(beeps)
  if err != nil {
    panic(err)
  }
  streams, forms := make([]beep.Streamer, 31), make([]beep.Format, 31)
  var sr beep.SampleRate
  var buffers []beep.Buffer
  for i := 0;i < len(beeps);i++ {
    f, err := os.Open("dat/sounds/"+beeps[i])
  //  fmt.Println(f.Name())
    if err != nil {
      panic(err)
    }
    soundStream, soundForm, err := wav.Decode(f)
    if err != nil {
      panic(err)
    }
    streams = append(streams, soundStream)
    forms = append(forms, soundForm)
    if i == 29 {
      sr = soundForm.SampleRate

    }

    soundBuffer := beep.NewBuffer(soundForm)
    soundBuffer.Append(soundStream)
    buffers = append(buffers, *soundBuffer)
  }
  speaker.Init(sr, sr.N(time.Second/10))

  numSoundsnames, err := os.Open("dat/sounds")
	if err != nil {
		panic(err)
	}
	defer numSoundsnames.Close()
	soundFiles, err := numSoundsnames.Readdirnames(100)
	if err != nil {
		panic(err)
	}

	_ = len(soundFiles)

  for {
    select {
    case <- sounds[0]:
      noise := buffers[0].Streamer(0, buffers[0].Len())
      speaker.Play(noise)

    case <- sounds[1]:
      noise := buffers[1].Streamer(1, buffers[1].Len())
      speaker.Play(noise)

    case <- sounds[2]:
      noise := buffers[2].Streamer(2, buffers[2].Len())
      speaker.Play(noise)
    case <- sounds[3]:
      noise := buffers[3].Streamer(3, buffers[3].Len())
      speaker.Play(noise)
    case <- sounds[4]:
      noise := buffers[4].Streamer(4, buffers[4].Len())
      speaker.Play(noise)
    case <- sounds[5]:
      noise := buffers[5].Streamer(5, buffers[5].Len())
      speaker.Play(noise)
    case <- sounds[6]:
      noise := buffers[6].Streamer(6, buffers[6].Len())
      speaker.Play(noise)
    case <- sounds[7]:
      noise := buffers[7].Streamer(7, buffers[7].Len())
      speaker.Play(noise)
    case <- sounds[8]:
      noise := buffers[8].Streamer(8, buffers[8].Len())
      speaker.Play(noise)
    case <- sounds[9]:
      noise := buffers[9].Streamer(9, buffers[9].Len())
      speaker.Play(noise)
    case <- sounds[10]:
      noise := buffers[10].Streamer(10, buffers[10].Len())
      speaker.Play(noise)
    case <- sounds[11]:
      noise := buffers[11].Streamer(11, buffers[11].Len())
      speaker.Play(noise)
    case <- sounds[12]:
      noise := buffers[12].Streamer(12, buffers[12].Len())
      speaker.Play(noise)
    case <- sounds[13]:
      noise := buffers[13].Streamer(13, buffers[13].Len())
      speaker.Play(noise)
    case <- sounds[14]:
      noise := buffers[14].Streamer(14, buffers[14].Len())
      speaker.Play(noise)
    case <- sounds[15]:
      noise := buffers[15].Streamer(15, buffers[15].Len())
      speaker.Play(noise)
    case <- sounds[16]:
      noise := buffers[16].Streamer(16, buffers[16].Len())
      speaker.Play(noise)
    case <- sounds[17]:
      noise := buffers[17].Streamer(17, buffers[17].Len())
      speaker.Play(noise)
    case <- sounds[18]:
      noise := buffers[18].Streamer(18, buffers[18].Len())
      speaker.Play(noise)
    case <- sounds[19]:
      noise := buffers[19].Streamer(19, buffers[19].Len())
      speaker.Play(noise)
    case <- sounds[20]:
      noise := buffers[20].Streamer(20, buffers[20].Len())
      speaker.Play(noise)
    case <- sounds[21]:
      noise := buffers[21].Streamer(21, buffers[21].Len())
      speaker.Play(noise)
    case <- sounds[22]:
      noise := buffers[22].Streamer(22, buffers[22].Len())
      speaker.Play(noise)
    case <- sounds[23]:
      noise := buffers[23].Streamer(23, buffers[23].Len())
      speaker.Play(noise)
    case <- sounds[24]:
      noise := buffers[24].Streamer(24, buffers[24].Len())
      speaker.Play(noise)
    case <- sounds[25]:
      noise := buffers[25].Streamer(25, buffers[25].Len())
      speaker.Play(noise)
    case <- sounds[26]:
      noise := buffers[26].Streamer(26, buffers[26].Len())
      speaker.Play(noise)
    case <- sounds[27]:
      noise := buffers[27].Streamer(27, buffers[27].Len())
      speaker.Play(noise)
    case <- sounds[28]:
      noise := buffers[28].Streamer(28, buffers[28].Len())
      speaker.Play(noise)
    case <- sounds[29]:
      noise := buffers[29].Streamer(29, buffers[29].Len())
      speaker.Play(noise)
    case <- sounds[30]:
      noise := buffers[30].Streamer(30, buffers[30].Len())
      speaker.Play(noise)
    default:
//        fmt.Println(strconv.Itoa(numSounds)+"MAKE SOME NOISE")
  //      noise := buffers[1].Streamer(0, buffers[1].Len())
    //    speaker.Play(noise)

    }
  }
}
