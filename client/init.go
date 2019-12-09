package main

import (
  "os"
  "bufio"
  "context"
  "time"
	"strconv"
	"strings"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
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
  play.Classes = make([]Class, 10, 10)
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

	play.Inventory = make([]InventoryItem, 20, 20)
  play.Equipped = make([]EquipmentItem, 20, 20)
  play.Rezz = 17
  play.MaxRezz = play.Rezz
  play.Tech = 17
  play.PlayerHash = hash(name+pass)

  var bank Account
  bank.Owner = play.Name
  bank.Amount = 0.0
  play.BankAccount = bank

	play.Str = 1
	play.Int = 1
	play.Dex = 1
	play.Wis = 1
	play.Con = 1
	play.Cha = 1
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


func InitZoneSpaces(SpaceRange string, zoneName string, desc string) {
  userFile, err := os.Open("creds")
  if err != nil {
    panic(err)
  }
  defer userFile.Close()
  scanner := bufio.NewScanner(userFile)
  scanner.Scan()
  user := scanner.Text()
  scanner.Scan()
  pass := scanner.Text()
  client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://"+user+":"+pass+"@sncn-hifs4.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	collection := client.Database("zones").Collection("Spaces")
	vnums := strings.Split(SpaceRange, "-")
	vnumStart, err := strconv.Atoi(vnums[0])
	if err != nil {
		panic(err)
	}

	vnumEnd, err := strconv.Atoi(vnums[1])
	if err != nil {
		panic(err)
	}
	for i := vnumStart;i < vnumEnd;i++ {
		var mobiles []int
		var items []int
		mobiles = append(mobiles, 0)
		items = append(items, 0)
		_, err = collection.InsertOne(context.Background(), bson.M{"vnums":SpaceRange,"zone":zoneName,"vnum":i, "desc":desc,
							"mobiles": mobiles, "items": items })
	}
	if err != nil {
		panic(err)
	}
}
