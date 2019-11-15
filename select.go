package main

import (
  "fmt"
)



func selected(item interface{}) {
  switch item.(type) {
  case Player:
  case Space:
  case Mobile:
  case EquipmentItem:
  case InventoryItem:
  case Broadcast:
  default:
    fmt.Printf("Nonetype")
  }
}
