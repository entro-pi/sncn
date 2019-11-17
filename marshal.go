package main


func decompInv(play Player) Player {

  num := play.ItemBank.SlotOneAmount
  for i := 0;i < num;i++ {
    //fmt.Println("\033[38:2:200:0:0mslot one\033[0m")

    play.Inventory[0] = play.ItemBank.SlotOne
    play.Inventory[0].Number++
  }
  play.ItemBank.SlotOneAmount = 0
  num = play.ItemBank.SlotTwoAmount
  for i := 0;i < num;i++ {
    //fmt.Println("\033[38:2:200:0:0mslot two\033[0m")
    play.Inventory[1] = play.ItemBank.SlotTwo
    play.Inventory[1].Number++
  }
  play.ItemBank.SlotTwoAmount = 0
  num = play.ItemBank.SlotThreeAmount
  for i := 0;i < num;i++ {
    play.Inventory[2] = play.ItemBank.SlotThree
    play.Inventory[2].Number++
  }
  play.ItemBank.SlotThreeAmount = 0
  num = play.ItemBank.SlotFourAmount
  for i := 0;i < num;i++ {
    play.Inventory[3] = play.ItemBank.SlotFour
    play.Inventory[3].Number++
  }
  play.ItemBank.SlotFourAmount = 0
  num = play.ItemBank.SlotFiveAmount
  for i := 0;i < num;i++ {
    play.Inventory[4] = play.ItemBank.SlotFive
    play.Inventory[4].Number++
  }
  play.ItemBank.SlotFiveAmount = 0
  num = play.ItemBank.SlotSixAmount
  for i := 0;i < num;i++ {
    play.Inventory[5] = play.ItemBank.SlotSix
    play.Inventory[5].Number++
  }
  play.ItemBank.SlotSixAmount = 0
  num = play.ItemBank.SlotSevenAmount
  for i := 0;i < num;i++ {
    play.Inventory[6] = play.ItemBank.SlotSeven
    play.Inventory[6].Number++
  }
  play.ItemBank.SlotSevenAmount = 0
  num = play.ItemBank.SlotEightAmount
  for i := 0;i < num;i++ {
    play.Inventory[7] = play.ItemBank.SlotEight
    play.Inventory[7].Number++
  }
  play.ItemBank.SlotEightAmount = 0
  num = play.ItemBank.SlotNineAmount
  for i := 0;i < num;i++ {
    play.Inventory[8] = play.ItemBank.SlotNine
    play.Inventory[8].Number++
  }
  play.ItemBank.SlotNineAmount = 0
  num = play.ItemBank.SlotTenAmount
  for i := 0;i < num;i++ {
    play.Inventory[9] = play.ItemBank.SlotTen
    play.Inventory[9].Number++
  }
  play.ItemBank.SlotTenAmount = 0
  return play
}

func composeInv(play Player) Player {
  if play.Inventory[0].Number >= 1 {
    play.ItemBank.SlotOne.Item = play.Inventory[0].Item
    play.ItemBank.SlotOneAmount = play.Inventory[0].Number
  }
  if play.Inventory[1].Number >= 1 {
    play.ItemBank.SlotTwo.Item = play.Inventory[1].Item
    play.ItemBank.SlotTwoAmount = play.Inventory[1].Number
  }
  if play.Inventory[2].Number >= 1 {
    play.ItemBank.SlotThree.Item = play.Inventory[2].Item
    play.ItemBank.SlotThreeAmount = play.Inventory[2].Number
  }
  if play.Inventory[3].Number >= 1 {
    play.ItemBank.SlotFour.Item = play.Inventory[3].Item
    play.ItemBank.SlotFourAmount = play.Inventory[3].Number
  }
  if play.Inventory[4].Number >= 1 {
    play.ItemBank.SlotFive.Item = play.Inventory[4].Item
    play.ItemBank.SlotFiveAmount = play.Inventory[4].Number
  }
  if play.Inventory[5].Number >= 1 {
    play.ItemBank.SlotSix.Item = play.Inventory[5].Item
    play.ItemBank.SlotSixAmount = play.Inventory[5].Number
  }
  if play.Inventory[6].Number >= 1 {
    play.ItemBank.SlotSeven.Item = play.Inventory[6].Item
    play.ItemBank.SlotSevenAmount = play.Inventory[6].Number
  }
  if play.Inventory[7].Number >= 1 {
    play.ItemBank.SlotEight.Item = play.Inventory[7].Item
    play.ItemBank.SlotEightAmount = play.Inventory[7].Number
  }
  if play.Inventory[8].Number >= 1 {
    play.ItemBank.SlotNine.Item = play.Inventory[8].Item
    play.ItemBank.SlotNineAmount = play.Inventory[8].Number
  }
  if play.Inventory[9].Number >= 1 {
    play.ItemBank.SlotTen.Item = play.Inventory[9].Item
    play.ItemBank.SlotTenAmount = play.Inventory[9].Number
  }
  return play
}


//Todo, make these receivers on a type
func initInv(play Player) Player {
  for i := 1;i < play.Inventory[0].Number;i++ {
    if i >= 1 {
//          fmt.Println("\033[38:2:0:200:0mINV\033[0m")
      play.ItemBank.SlotOne.Item = play.Inventory[0].Item
      play.ItemBank.SlotOneAmount++
    }else {
      var blank Object
      play.ItemBank.SlotOne.Item = blank
      play.ItemBank.SlotOneAmount = 0
    }
  }


  for i := 1;i < play.Inventory[1].Number;i++ {
    if i >= 1 {
//          fmt.Println("\033[38:2:0:200:0mINV\033[0m")
      play.ItemBank.SlotTwo.Item = play.Inventory[1].Item

      play.ItemBank.SlotTwoAmount++
    }else {
      var blank Object
      play.ItemBank.SlotTwo.Item = blank
      play.ItemBank.SlotTwoAmount = 0
    }
  }
  for i := 1;i < play.Inventory[2].Number;i++ {
    if i >= 1 {
//          fmt.Println("\033[38:2:0:200:0mINV\033[0m")
      play.ItemBank.SlotThree.Item = play.Inventory[2].Item
      play.ItemBank.SlotThreeAmount++
    }else {
      var blank Object
      play.ItemBank.SlotThree.Item = blank
      play.ItemBank.SlotThreeAmount = 0
    }
  }
  for i := 1;i < play.Inventory[3].Number;i++ {
    if i >= 1 {
//          fmt.Println("\033[38:2:0:200:0mINV\033[0m")
      play.ItemBank.SlotFour.Item = play.Inventory[3].Item
      play.ItemBank.SlotFourAmount++
    }else {
      var blank Object
      play.ItemBank.SlotFour.Item = blank
      play.ItemBank.SlotFourAmount = 0
    }
  }
  for i := 1;i < play.Inventory[4].Number;i++ {
    if i >= 1 {
//          fmt.Println("\033[38:2:0:200:0mINV\033[0m")
      play.ItemBank.SlotFive.Item = play.Inventory[4].Item
      play.ItemBank.SlotFiveAmount++
    }else {
      var blank Object
      play.ItemBank.SlotFive.Item = blank
      play.ItemBank.SlotFiveAmount = 0
    }
  }
  for i := 1;i < play.Inventory[5].Number;i++ {
    if i >= 1 {
//          fmt.Println("\033[38:2:0:200:0mINV\033[0m")
      play.ItemBank.SlotSix.Item = play.Inventory[5].Item
      play.ItemBank.SlotSixAmount++
    }else {
      var blank Object
      play.ItemBank.SlotSix.Item = blank
      play.ItemBank.SlotSixAmount = 0
    }
  }
  for i := 1;i < play.Inventory[6].Number;i++ {
    if i >= 1 {
//          fmt.Println("\033[38:2:0:200:0mINV\033[0m")
      play.ItemBank.SlotSeven.Item = play.Inventory[6].Item
      play.ItemBank.SlotSevenAmount++
    }else {
      var blank Object
      play.ItemBank.SlotSeven.Item = blank
      play.ItemBank.SlotSevenAmount = 0
    }
  }
  for i := 1;i < play.Inventory[7].Number;i++ {
    if i >= 1 {
//          fmt.Println("\033[38:2:0:200:0mINV\033[0m")
      play.ItemBank.SlotEight.Item = play.Inventory[7].Item
      play.ItemBank.SlotEightAmount++
    }else {
      var blank Object
      play.ItemBank.SlotEight.Item = blank
      play.ItemBank.SlotEightAmount = 0
    }
  }
  for i := 1;i < play.Inventory[8].Number;i++ {
    if i >= 1 {
//          fmt.Println("\033[38:2:0:200:0mINV\033[0m")
      play.ItemBank.SlotNine.Item = play.Inventory[8].Item
      play.ItemBank.SlotNineAmount++
    }else {
      var blank Object
      play.ItemBank.SlotNine.Item = blank
      play.ItemBank.SlotNineAmount = 0
    }
  }
  for i := 1;i < play.Inventory[9].Number;i++ {
    if i >= 1 {
//          fmt.Println("\033[38:2:0:200:0mINV\033[0m")
      play.ItemBank.SlotTen.Item = play.Inventory[9].Item
      play.ItemBank.SlotTenAmount++
    }else {
      var blank Object
      play.ItemBank.SlotTen.Item = blank
      play.ItemBank.SlotTenAmount = 0
    }
  }
  return play
}
