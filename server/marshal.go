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
func decompEq(play Player) Player {

  num := play.EqBank.SlotOneAmount
  for i := 0;i < num;i++ {
    //fmt.Println("\033[38:2:200:0:0mslot one\033[0m")

    play.Equipped[0] = play.EqBank.SlotOne
    play.Equipped[0].Number++
  }
  play.EqBank.SlotOneAmount = 0
  num = play.EqBank.SlotTwoAmount
  for i := 0;i < num;i++ {
    //fmt.Println("\033[38:2:200:0:0mslot two\033[0m")
    play.Equipped[1] = play.EqBank.SlotTwo
    play.Equipped[1].Number++
  }
  play.EqBank.SlotTwoAmount = 0
  num = play.EqBank.SlotThreeAmount
  for i := 0;i < num;i++ {
    play.Equipped[2] = play.EqBank.SlotThree
    play.Equipped[2].Number++
  }
  play.EqBank.SlotThreeAmount = 0
  num = play.EqBank.SlotFourAmount
  for i := 0;i < num;i++ {
    play.Equipped[3] = play.EqBank.SlotFour
    play.Equipped[3].Number++
  }
  play.EqBank.SlotFourAmount = 0
  num = play.EqBank.SlotFiveAmount
  for i := 0;i < num;i++ {
    play.Equipped[4] = play.EqBank.SlotFive
    play.Equipped[4].Number++
  }
  play.EqBank.SlotFiveAmount = 0
  num = play.EqBank.SlotSixAmount
  for i := 0;i < num;i++ {
    play.Equipped[5] = play.EqBank.SlotSix
    play.Equipped[5].Number++
  }
  play.EqBank.SlotSixAmount = 0
  num = play.EqBank.SlotSevenAmount
  for i := 0;i < num;i++ {
    play.Equipped[6] = play.EqBank.SlotSeven
    play.Equipped[6].Number++
  }
  play.EqBank.SlotSevenAmount = 0
  num = play.EqBank.SlotEightAmount
  for i := 0;i < num;i++ {
    play.Equipped[7] = play.EqBank.SlotEight
    play.Equipped[7].Number++
  }
  play.EqBank.SlotEightAmount = 0
  num = play.EqBank.SlotNineAmount
  for i := 0;i < num;i++ {
    play.Equipped[8] = play.EqBank.SlotNine
    play.Equipped[8].Number++
  }
  play.EqBank.SlotNineAmount = 0
  num = play.EqBank.SlotTenAmount
  for i := 0;i < num;i++ {
    play.Equipped[9] = play.EqBank.SlotTen
    play.Equipped[9].Number++
  }
  play.EqBank.SlotTenAmount = 0
  return play
}

func composeEq(play Player) Player {
  if play.Equipped[0].Number >= 1 {
    play.EqBank.SlotOne.Item = play.Equipped[0].Item
    play.EqBank.SlotOneAmount = play.Equipped[0].Number
  }
  if play.Equipped[1].Number >= 1 {
    play.EqBank.SlotTwo.Item = play.Equipped[1].Item
    play.EqBank.SlotTwoAmount = play.Equipped[1].Number
  }
  if play.Equipped[2].Number >= 1 {
    play.EqBank.SlotThree.Item = play.Equipped[2].Item
    play.EqBank.SlotThreeAmount = play.Equipped[2].Number
  }
  if play.Equipped[3].Number >= 1 {
    play.EqBank.SlotFour.Item = play.Equipped[3].Item
    play.EqBank.SlotFourAmount = play.Equipped[3].Number
  }
  if play.Equipped[4].Number >= 1 {
    play.EqBank.SlotFive.Item = play.Equipped[4].Item
    play.EqBank.SlotFiveAmount = play.Equipped[4].Number
  }
  if play.Equipped[5].Number >= 1 {
    play.EqBank.SlotSix.Item = play.Equipped[5].Item
    play.EqBank.SlotSixAmount = play.Equipped[5].Number
  }
  if play.Equipped[6].Number >= 1 {
    play.EqBank.SlotSeven.Item = play.Equipped[6].Item
    play.EqBank.SlotSevenAmount = play.Equipped[6].Number
  }
  if play.Equipped[7].Number >= 1 {
    play.EqBank.SlotEight.Item = play.Equipped[7].Item
    play.EqBank.SlotEightAmount = play.Equipped[7].Number
  }
  if play.Equipped[8].Number >= 1 {
    play.EqBank.SlotNine.Item = play.Equipped[8].Item
    play.EqBank.SlotNineAmount = play.Equipped[8].Number
  }
  if play.Equipped[9].Number >= 1 {
    play.EqBank.SlotTen.Item = play.Equipped[9].Item
    play.EqBank.SlotTenAmount = play.Equipped[9].Number
  }
  return play
}
