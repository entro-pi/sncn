package main

import (
  "bufio"
  "os"
  "fmt"
  "strconv"
)

func readItemsFromFile(filePath string) []Object {
  file, err := os.Open(filePath)
  if err != nil {
    panic(err)
  }
  fmt.Println("LOADING ITEMS")
  var objHolder []Object
  scanner := bufio.NewScanner(file)
  var obj Object

  for scanner.Scan() {

    if scanner.Text() == "VNUM" {
      scanner.Scan()
      //fmt.Println("VNUM")
      //fmt.Println(scanner.Text())
      obj.Vnum, err = strconv.Atoi(scanner.Text())
      if err != nil {
        panic(err)
      }
    }
    if scanner.Text() == "NAME" {
      scanner.Scan()
      //fmt.Println("NAME")
      //fmt.Println(scanner.Text())
      obj.Name = scanner.Text()
    }
    if scanner.Text() == "LONGNAME" {
      scanner.Scan()
      //fmt.Println("LONGNAME")
      //fmt.Println(scanner.Text())
      obj.LongName = scanner.Text()
    }
    if scanner.Text() == "ZONE" {
      scanner.Scan()
      //fmt.Println("ZONE")
      //fmt.Println(scanner.Text())
      obj.Zone = scanner.Text()
    }
    if scanner.Text() == "VALUE" {
      scanner.Scan()
      //fmt.Println("VALUE")
      //fmt.Println(scanner.Text())
      obj.Value, err = strconv.Atoi(scanner.Text())
      if err != nil {
        panic(err)
      }
    }
    if scanner.Text() == "OWNED" {
      scanner.Scan()
      //fmt.Println("OWNED")
      //fmt.Println(scanner.Text())
      if scanner.Text() == "true"{
        obj.Owned = true
      }else {
        obj.Owned = false
      }
    }
    if scanner.Text() == "SLOT" {
      scanner.Scan()
      //fmt.Println("SLOT")
      //fmt.Println(scanner.Text())
      obj.Slot, err = strconv.Atoi(scanner.Text())
      if err != nil {
        panic(err)
      }
      objHolder = append(objHolder, obj)

    }
  }
  fmt.Println("DONE LOADING ITEMS")
  return objHolder
}

func onlineTransaction(advert *Broadcast, customer Player, allItems []Object) (Player, string) {
  output := ""
  if advert.Payload.Transaction.Sold {
    return customer, "SOLD OUT"
  }
  if len(advert.Payload.Store.Inventory) > 0 {
      for i := 0;i < len(advert.Payload.Store.Inventory);i++ {
        hasSpace := false
        slot := 0
        vnum := advert.Payload.Store.Inventory[i].Item.Vnum
        hash := advert.Payload.Store.Inventory[i].ItemHash
        price := advert.Payload.Store.Inventory[i].Price
        customerCash := customer.BankAccount.Amount
        isSold := advert.Payload.Store.Inventory[i].Sold
        fmt.Println("VNUM",vnum,"HASH",hash,"PRICE",price,"CUSTOMERCASH",customerCash,"ISSOLD",isSold)
        for c := len(customer.Inventory) - 1;c > 0;c-- {
          if customer.Inventory[c].Item.Name == "nothing" {
            hasSpace = true
            slot = c
          }
        }
        if customerCash >= price && hasSpace {
            customer.BankAccount.Amount -= price
        }
        if hash == onlineHash(allItems[vnum].LongName) {
          customer.Inventory[slot].Item = allItems[vnum]
          customer.Inventory[slot].Number++
          advert.Payload.Store.Inventory[i].Sold = true
          fmt.Println("\033[38:2:0:200:0mTransaction approved.\033[0m")
          output += fmt.Sprintln("\033[38:2:0:200:0mTransaction approved.\033[0m")

          return customer, output
        }
      }
  }else if !advert.Payload.Transaction.Sold {
      hasSpace := false
      hasCash := false
      slot := 0
      vnum := advert.Payload.Transaction.Item.Vnum
      hash := advert.Payload.Transaction.ItemHash
      price := advert.Payload.Transaction.Price
      customerCash := customer.BankAccount.Amount

      fmt.Println("VNUM",vnum,"HASH",hash,"PRICE",price,"CUSTOMERCASH",customerCash)
      for c := len(customer.Inventory) - 1;c > 0;c-- {
        if customer.Inventory[c].Item.Name == "nothing" || customer.Inventory[c].Item.Name == "" || customer.Inventory[c].Item.Name == advert.Payload.Transaction.Item.Name{
          hasSpace = true
          slot = c
        }
      }
      if customerCash >= price && hasSpace {
          customer.BankAccount.Amount -= price
          hasCash = true
      }
      if hasSpace && hasCash {
        customer.Inventory[slot].Item = allItems[vnum]
        customer.Inventory[slot].Number++
        fmt.Println("\033[38:2:0:200:0mTransaction approved.\033[0m")
        output += fmt.Sprintln("\033[38:2:0:200:0mTransaction approved.\033[0m")
        return customer, output

      }
  }else {
    output += fmt.Sprint("Looks like you missed out on the sale!")
    output += fmt.Sprint("That is sold out!")
  }
  output += fmt.Sprintln("\033[38:2:200:0:0mTransaction declined.\033[0m")
  fmt.Println("\033[38:2:200:0:0mTransaction declined.\033[0m")
  return customer, output
}

func stack(play Player) Player {
  //need to do a deepequals

  return play
}
func lookupObject(vnum int) Object {
  var obj Object
  switch vnum {
  case 1:
    obj.Vnum = 1
    obj.Name = "a red rose"
    obj.LongName = "A rose floats here, slowly rotating."
    obj.Zone = "zem"
    obj.Value = 1
    obj.Owned = false
    obj.Slot = 2
    return obj
  case 2:
    obj.Vnum = 2
    obj.Name = "a nyancat"
    obj.LongName = "A poptart kitten happily miaos in a circle."
    obj.Zone = "zem"
    obj.Value = 100
    obj.Owned = false
    obj.Slot = 1
  default:
    obj.Vnum = 0
    obj.Name = "nothing"
    obj.LongName = "The lack of things is...lacking"
    obj.Zone = "zem"
    obj.Value = 0
    obj.Owned = false
    obj.Slot = 0
  }
  return obj
}
