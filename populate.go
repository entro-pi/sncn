package main

import (

  "context"
  "time"
  "strings"
  "math/rand"
  "fmt"
  "strconv"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func UIDMaker() string {
	hostname := "localhost"
	username := "username"
	//Inspired by 'Una (unascribed)'s bikeshed
	rand.Seed(int64(time.Now().Nanosecond()))
	adjectives := []string{"Accidental", "Allocated", "Asymptotic", "Background", "Binary",
		"Bit", "Blast", "Blocked", "Bronze", "Captured", "Classic",
		"Compact", "Compressed", "Concatenated", "Conventional",
		"Cryptographic", "Decimal", "Decompressed", "Deflated",
		"Defragmented", "Dirty", "Distinguished", "Dozenal", "Elegant",
		"Encrypted", "Ender", "Enhanced", "Escaped", "Euclidean",
		"Expanded", "Expansive", "Explosive", "Extended", "Extreme",
		"Floppy", "Foreground", "Fragmented", "Garbage", "Giga", "Gold",
		"Hard", "Helical", "Hexadecimal", "Higher", "Infinite", "Inflated",
		"Intentional", "Interlaced", "Kilo", "Legacy", "Lower", "Magical",
		"Mapped", "Mega", "Nonlinear", "Noodle", "Null", "Obvious", "Paged",
		"Parity", "Platinum", "Primary", "Progressive", "Prompt",
		"Protected", "Quick", "Real", "Recursive", "Replica", "Resident",
		"Retried", "Root", "Secure", "Silver", "SolidState", "Super",
		"Swap", "Switched", "Synergistic", "Tera", "Terminated", "Ternary",
		"Traditional", "Unlimited", "Unreal", "Upper", "Userspace",
		"Vector", "Virtual", "Web", "WoodGrain", "Written", "Zipped"}
	nouns := []string{"AGP", "Algorithm", "Apparatus", "Array", "Bot", "Bus", "Capacitor",
		"Card", "Chip", "Collection", "Command", "Connection", "Cookie",
		"DLC", "DMA", "Daemon", "Data", "Database", "Density", "Desktop",
		"Device", "Directory", "Disk", "Dongle", "Executable", "Expansion",
		"Folder", "Glue", "Gremlin", "IRQ", "ISA", "Instruction",
		"Interface", "Job", "Key", "List", "MBR", "Map", "Modem", "Monster",
		"Numeral", "PCI", "Paradigm", "Plant", "Port", "Process",
		"Protocol", "Registry", "Repository", "Rights", "Scanline", "Set",
		"Slot", "Smoke", "Sweeper", "TSR", "Table", "Task", "Thread",
		"Tracker", "USB", "Vector", "Window"}
	uniquefier := ""

	uniqe := ""
	for i := 0; i < 2; i++ {
		uniq := rand.Intn(15)
		if uniq >= 10 {
			switch uniq {
			case 10:
				uniqe = "A"
			case 11:
				uniqe = "B"
			case 12:
				uniqe = "C"
			case 13:
				uniqe = "D"
			case 14:
				uniqe = "E"
			case 15:
				uniqe = "F"
			}

			uniquefier += uniqe
		} else {
			uniquefier += fmt.Sprint(uniq)
		}
	}
	ind := rand.Intn(len(adjectives))
	indie := rand.Intn(len(adjectives))
	if indie == ind {
		indie = rand.Intn(len(adjectives))
	}
	thedog := rand.Intn(len(nouns))
	uniqueFied := fmt.Sprint(uniquefier, adjectives[ind], adjectives[indie], nouns[thedog])

	//fmt.Println(uniqueFied)

	UID := fmt.Sprint(uniqueFied, hostname, username)
	return UID
}

func PopulateAreaMobiles() []Mobile {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	var Mobiles []Mobile
	collection := client.Database("npcs").Collection("mobiles")
	results, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	for results.Next(context.Background()) {

			var Mobile Mobile
			err := results.Decode(&Mobile)
			if err != nil {
				panic(err)
			}
			Mobiles = append(Mobiles, Mobile)

//			fmt.Println(Spaces.Vnum)
	}
	return Mobiles
}

func PopulateAreaBuild(rangeVnums string) []Space {

  beginString := strings.Split(rangeVnums, "-")[0]

  endString := strings.Split(rangeVnums, "-")[1]

  begin, err := strconv.Atoi(beginString)
  if err != nil {
    panic(err)
  }
  end, err := strconv.Atoi(endString)
  if err != nil {
    panic(err)
  }
  length := end - begin
	areas := make([]Space, length)
	return areas
}

func PopulateAreas() []Space {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	var Spaces []Space
	collection := client.Database("zones").Collection("Spaces")
	results, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	for results.Next(context.Background()) {

			var Space Space
			err := results.Decode(&Space)
			if err != nil {
				panic(err)
			}
			Spaces = append(Spaces, Space)

//			fmt.Println(Spaces.Vnum)
	}
	return Spaces
}
