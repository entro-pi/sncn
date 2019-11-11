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
