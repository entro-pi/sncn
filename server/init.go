package main

import (
	"strings"
)

func initDigRoom(digFrame [][]int, zoneVnums string, zoneName string, play Player, vnum int) (Space, int) {
	var dg Space
	dg.Vnums = zoneVnums
	dg.Zone = zoneName
	dg.ZonePos = make([]int, 2)
	dg.ZoneMap = digFrame
	//todo directions
	vnum += 1
	dg.Vnum = vnum
	dg.Altered = true
	dg.Desc = "Nothing but some cosmic rays"
	for len(strings.Split(dg.Desc, "\n")) < 8 {
		dg.Desc += "\n"
	}
	return dg, vnum
}

func addClass(play Player) Player {
  var class Class
  play.Classes = append(play.Classes, class)
  play.Classes[0].Level = 1
  play.Classes[0].Name = "wildling"
  var rip Skill
  rip.DamType = "slash"
  rip.Level = 0
  rip.Usage = 'e'
  play.Classes[0].Skills = append(play.Classes[0].Skills, rip)
  var blast Spell
  blast.Usage = 'w'
  blast.TechUsage = 2
  blast.Level = 1
  blast.Consumed = false
  play.Classes[0].Spells = append(play.Classes[0].Spells, blast)
  return play
}

func InitPlayer(name string, pass string) Player {
	var play Player
	var inv []int
	inv = append(inv, 1)
	play.Name = name
	play.Title = "The Unknown"
  play.Classes = make([]Class, 1, 1)
  var rip Skill
  rip.Name = "overcharge"
  rip.DamType = "slash"
  rip.Level = 0
  rip.Usage = 'e'
  play.Classes[0].Skills = append(play.Classes[0].Skills, rip)
  var blast Spell
  blast.TechUsage = 2
  blast.Level = 1
  blast.Consumed = false
  blast.Name = "blast"
  play.Classes[0].Spells = append(play.Classes[0].Spells, blast)

  play.Inventory = make([]InventoryItem, 1, 1)
  play.Equipped = make([]EquipmentItem, 1, 1)
  play.Rezz = 100
  play.MaxRezz = play.Rezz
  play.Tech = 100
  play.MaxTech = play.Tech
  play.Mana = 100
  play.MaxMana = play.Mana
  play.PlayerHash = hash(name+pass)

  var bank Account
  bank.Owner = play.Name
  bank.Amount = 0.0
  play.BankAccount = bank

	play.Str = 10
	play.Int = 10
	play.Dex = 10
	play.Wis = 10
	play.Con = 10
	play.Cha = 10
  play.Channels = append(play.Channels, "")
	return play

}
func InitObject() Object {
  var obj Object
  obj.Name = "a golden tiara"
  obj.LongName = "A golden tiara lies here."
  obj.Vnum = 1
  obj.Owned = false
  return obj
}
func InitMob() Mobile {
  var mob Mobile
  mob.Name = "rabid ferret"
  mob.LongName = "A rabid ferret charges towards you!"
  mob.Rezz = 30
  mob.Tech = 2
  mob.Aggro = false
  mob.Align = -1
  return mob
}

func InitFight() Fight {
  var newFight Fight
  return newFight
}


