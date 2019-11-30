package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
	"log"
	"github.com/fsnotify/fsnotify"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func getConnectionString() string {
	f, err := os.Open("creds")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	scanned := scanner.Text()
	return scanned
}

func main() {
	fileChange := make(chan bool)
	if len(os.Args) == 2 {
	if os.Args[1] == "--main" {
			scanner := bufio.NewScanner(os.Stdin)
		//	fmt.Print("Enter your command")
			fmt.Print("Initializing a player")
			user, pword := LoginSC()
			play := InitPlayer(user, pword)
			whoList := who(play.Name)
			fmt.Println(whoList)
		//	go actOn() //for receiving in Go
//			go watch(play)
			for scanner.Scan() {
				input := scanner.Text()
				//Should probably do some error checking before
				//passing it along
				if len(strings.Split(input, ":")) <= 1 {
					continue
				}else {
					doPlayer(input, play)
					go doWatch(input, play, fileChange)
					go doInput(input, play, fileChange, whoList)
			//		fmt.Print("Enter your command")
				}
				fmt.Print("\033[26;53H\n")

			}
	}
	if os.Args[1] == "--secondary" {
			scanner := bufio.NewScanner(os.Stdin)
		//	fmt.Print("Enter your command")
			user, pword := LoginSC()
			fmt.Print("Initializing a player")
			play := InitPlayer(user, pword)
			whoList := who(play.Name)
			fmt.Println(whoList)
			go actOn(play, fileChange, whoList) //for receiving in Go
			go watch(play, fileChange)
			for scanner.Scan() {
				input := scanner.Text()
				//Should probably do some error checking before
				//passing it along
				if len(strings.Split(input, " ")) <= 1 {
					continue
				}else {
					play, input = parseInput(play, input)
					doPlayer(input, play)
					go watch(play, fileChange)
					go doInput(input, play, fileChange, whoList)
					go doWatch(input, play, fileChange)
			//		fmt.Print("Enter your command")
				}
				fmt.Print("\033[26;53H\n")

			}
		}
	}
}

func parseInput(play Player, input string) (Player, string) {
	var value []string

	fmt.Println("INPUT IS ",input)
	if strings.HasPrefix(input, "tell") {
		if strings.Split(input, " ")[1] == play.Name {
			input = strings.ReplaceAll(input, " ", ":")
			return play, input
		}
	}
	if strings.HasPrefix(input, "broadcast") {
		input = strings.ReplaceAll(input, " ", ":")
		return play, input
	}


	if strings.HasPrefix(input, "generate") {
		if len(strings.Split(input, " ")) > 1 {
			value = strings.Split(input, " ")
		}else {
			for len(value) < 2 {
				value = append(value, "broadcast: ")
			}
		}
		for i := 0;i < len(value);i++ {
			fmt.Println("VALUE IS :",value[1])
			switch value[i] {
			case "1":
				fmt.Println("generating a tiara")
				object := InitObject()
				play.Inventory[0].Item = object
				play.Inventory[0].Number++
			default:
				input += "broadcast:"
			}
		}
	}

	return play, input
}

func who(newPlayer string) []string {
	var oldPlayers []string
	f, err := os.OpenFile("../pot/who", os.O_RDWR|os.O_CREATE, 0644) 
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	longUser := strings.Split(string(content), "\n")
	for i := 0;i < len(longUser);i++ {
		if newPlayer == longUser[i] {
			continue
		}else {
			oldPlayers = append(oldPlayers, longUser[i])
		}
	}
	_, err = f.WriteString(newPlayer+"\n")
	if err != nil {
		panic(err)
	}
	err = f.Sync()
	if err != nil {
		panic(err)
	}
	f.Close()
	oldPlayers = append(oldPlayers, newPlayer)
	return oldPlayers
}

func doPlayer(input string, play Player) {
	play = decompEq(play)
	play = decompInv(play)
	DescribePlayer(play)
	describeEquipment(play)
	describeInventory(play)
}

func doInput(input string, play Player, fileChange chan bool, whoList []string) {
	connection := getConnectionString()
	conn, err := amqp.Dial(connection)

	direct := false

	//Determine if we're sending to anyone in particular
	inputArray := strings.Split(input, ":")
	if len(inputArray) < 2 {
		inputArray = append(inputArray, ":")
	}
	tellTo := ""
	for i := 0;i < len(whoList);i++ {
		fmt.Println(inputArray[1])
		if inputArray[0] == "tell" && inputArray[1] == whoList[i] {
			direct = true
			tellTo = inputArray[1]
			fmt.Println("\033[48:2:0:0:120m",direct, tellTo,"\033[0m")
			break
		}
	}


	failOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")

	defer ch.Close()
	q, err := ch.QueueDeclare(
		"", //name
		true, // durable
		false, //delete when used
		false, //exclusive
		false, //no-wait
		nil, //arguments
	)
/*	q, err := ch.QueueDeclare(
		"input", //name
		true, // durable
		true, //delete when used
		false, //exclusive
		false, //no-wait
		nil, //arguments
	)
*/	failOnError(err, "Failed to declare a queue")

	err = ch.ExchangeDeclare(
	"ballast", //name
	"topic", //type
	false, //durable
	false, //auto-delted
	false, //internal
	false, //no wait
	nil, //arguments
	)
	failOnError(err, "Failed to declare an exchange")

	err = ch.QueueBind(
		q.Name, //queue name
		"", //routing key
		"ballast",//exchange
		false,
		nil,
	)
	body := ""
	for i := 0;i < len(inputArray);i++ {
		body += inputArray[i]+" "
	}
	if direct {
		//body = strings.ReplaceAll(body, "broadcast", "tell")
		err = ch.Publish(
		"ballast", //exchange
		tellTo+".tell", // routing key
		false, //mandatory
		false, //immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body: []byte(body),
		})
	}else {
		err = ch.Publish(
		"ballast", //exchange
		"", // routing key
		false, //mandatory
		false, //immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body: []byte(body),
		})

	}

//	fmt.Print("\033[26;53H\n")
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")

}


func actOn(play Player, fileChange chan bool, whoList []string) {
        connection := getConnectionString()
        conn, err := amqp.Dial(connection)

        failOnError(err, "Failed to connect to RabbitMQ")

        defer conn.Close()

        ch, err := conn.Channel()

        failOnError(err, "Failed to open a channel")

        defer ch.Close()
    /*    chDirect, err := conn.Channel()

        failOnError(err, "Failed to open a channel")

        defer chDirect.Close()
*/
	err = ch.ExchangeDeclare(
		"ballast",//name
		"topic",//type
		false, //durable
		false, //auto-deleted
		false, //internal
		false, //no wait
		nil, //args
	)
	failOnError(err, "Failed to declare an exchange")

        q, err := ch.QueueDeclare(
                "", //name
                false, // durable
                false, //delete when used
                false, //exclusive
                false, //no-wait
                nil, //arguments
        )
        failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name, //queue name
		"", //routing key
		"ballast", //exchange
		false,
		nil,
	)
	err = ch.QueueBind(
		q.Name, //queue name
		play.Name+".tell", //routing key
		"ballast", //exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")
	msgs, err := ch.Consume(
		q.Name, //queue
		"",
		true, //auto-ack
		false, //exclusive
		false, //no-local
		false, //no-wait
		nil, //args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	for {
		
		//select {
		//default:
		go func() {
			for msg := range msgs {
				message := string(msg.Body)

				if strings.Split(message, " ")[1] == play.Name {

					fmt.Print("\033[26;53H\n")
					log.Printf("\033[38:2:150:150:0mReceived a tell: %s\033[0m", msg.Body)
					if strings.HasPrefix(message, "tell") {
						if !strings.Contains(message, "!:::tick:::!") {
							f, err := os.OpenFile("../pot/tells", os.O_APPEND|os.O_WRONLY, 0644) 
							if err != nil {
								panic(err)
							}
							//strip the thingies out
	//						message = strings.ReplaceAll(message, "tell:", "\033[38:2:150:0:100mtell")

							_, err = f.WriteString(message+"\n")
							if err != nil {
								panic(err)
							}
							fileChange <- true
							var blank Player 
							f.Close()
							go doWatch(string(msg.Body), blank, fileChange)
							continue
						}
				
					}else {
						var blank Player 

						go doWatch("!:::tick:::!", blank, fileChange)
					}
				}
				if err != nil {
					panic(err)
				}
			
				fmt.Print("\033[26;53H\n")
				log.Printf("\033[38:2:0:150:150mReceived a message: %s\033[0m", msg.Body)
				if strings.HasPrefix(message, "broadcast") {
					var blank Player
					if !strings.Contains(message, "!:::tick:::!") {
						f, err := os.OpenFile("../pot/broadcast", os.O_APPEND|os.O_WRONLY, 0644) 
						if err != nil {
							panic(err)
						}
						//strip the thingies out
						message = strings.ReplaceAll(message, "broadcast", "")
						_, err = f.WriteString(message+"\n")
						if err != nil {
							panic(err)
						}
						f.Close()
						fileChange <- true

						go doWatch(string(msg.Body), blank, fileChange)
					}else {
						go doWatch("!:::tick:::!", blank, fileChange)
					}
				}
				if err != nil {
					panic(err)
				}
			}
		}()

//		}
		<-forever
	}
}



func watch(play Player, fileChange chan bool) {
	var broadcastContainer []string

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
	    log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
	    for {
	        select {

	        case <-fileChange:
			drawTells(play, 0, 0)
			broadcastContainer = nil
			file, err := os.Open("../pot/broadcast")
			if err != nil {
				panic(err)
			}
			defer file.Close()
			contents, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			var lines []string
			lines = nil
			lines = strings.Split(string(contents), "\n")
			lineIn := strings.Split(string(contents), "\n")
			if len(lines) >= 20 {
				lines = nil
				for i := len(lineIn)-1;i > len(lineIn)-21;i-- {
					lineIn[i] = strings.ReplaceAll(lineIn[i], "broadcast:", "")
					lines = append(lines, lineIn[i])
				}
			}
//			var broadcastContainer []Broadcast
			col := 0
			row := 0
			colVal := 53
			rowVal := 0
			for i := 0;i < len(lines);i++ {
					var newBroad Broadcast
					newBroad.Payload.Message = lines[i]
					newBroad.Payload.Name = play.Name
					newBroad.Payload.Game = "snowcrash.network"
					if len(newBroad.Payload.Message) > 89 {
						newBroad.Payload.Message = lines[i][:89]
					}
					if strings.Contains(lines[i], "!:::tick:::!") {
						continue
					}

					newBroadPayload := AssembleBroadside(newBroad, rowVal, colVal)
					broadcastContainer = append(broadcastContainer, newBroadPayload)
					if row >= 5 {
						row = 0
						rowVal = 0
					}
					if col < 3 {
						col++
						colVal += 30
					}else {
						row++
						rowVal += 4
						col = 0
						colVal = 53
					}
				}
				for i := 0;i < len(broadcastContainer);i++ {
					fmt.Print(broadcastContainer[i])
				}
				fmt.Print("\033[26;53H\n")

				//log.Print(string(contents))
	        case event, ok := <-watcher.Events:
	            if !ok {
	                return
	            }
	           // fmt.Print("\033[26;53H\n")
		  //  log.Print("event:", event)
	            if event.Op&fsnotify.Write == fsnotify.Write {
	        //        log.Print("\033[48:2:150:0:150mmodified file:", event.Name,"\033[0m")
	            }
		if event.Name == "../pot/broadcast" || event.Name == "../pot/tells" {
			drawTells(play, 157, 24)
			broadcastContainer = nil
			file, err := os.Open(event.Name)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			contents, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			var lines []string
			lines = nil
			lines = strings.Split(string(contents), "\n")
			lineIn := strings.Split(string(contents), "\n")
			if len(lines) >= 20 {
				lines = nil
				for i := len(lineIn)-1;i > len(lineIn)-21;i-- {
					lineIn[i] = strings.ReplaceAll(lineIn[i], "broadcast:", "")
					lines = append(lines, lineIn[i])
				}
			}
//			var broadcastContainer []Broadcast
			col := 0
			row := 0
			colVal := 53
			rowVal := 0
			for i := 0;i < len(lines);i++ {
					var newBroad Broadcast
					newBroad.Payload.Message = lines[i]
					newBroad.Payload.Name = play.Name
					newBroad.Payload.Game = "snowcrash.network"
					if len(newBroad.Payload.Message) > 89 {
						newBroad.Payload.Message = lines[i][:89]
					}
					if strings.Contains(lines[i], "!:::tick:::!") {
						continue
					}

					newBroadPayload := AssembleBroadside(newBroad, rowVal, colVal)
					broadcastContainer = append(broadcastContainer, newBroadPayload)
					if row >= 5 {
						row = 0
						rowVal = 0
					}
					if col < 3 {
						col++
						colVal += 30
					}else {
						row++
						rowVal += 4
						col = 0
						colVal = 53
					}
				}
				for i := 0;i < len(broadcastContainer);i++ {
					fmt.Print(broadcastContainer[i])
				}
				fmt.Print("\033[26;53H\n")

				//log.Print(string(contents))
			}
	        case err, ok := <-watcher.Errors:
	            if !ok {
	                return
	            }
	            	fmt.Print("\033[26;53H\n")
			log.Print("error:", err)
		default:
//			for i := 0;i < len(broadcastContainer);i++ {
//				fmt.Print(broadcastContainer[i])
//			}
			//DO NOTHING
	        }
	    }
	}()

	err = watcher.Add("../pot")
	if err != nil {
	    log.Fatal(err)
	}
	<-done
}
func drawTells(play Player, colVal int, rowVal int) {
	var broadcastContainer []string
	file, err := os.Open("../pot/tells")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	row := 0

	tells, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(tells), "\n")
	for i := 0;i < len(lines);i++ {
			var newBroad Broadcast
			newBroad.Payload.Message = lines[i]
			newBroad.Payload.Name = play.Name
			newBroad.Payload.Game = "snowcrash.network"
			if len(newBroad.Payload.Message) > 89 {
				newBroad.Payload.Message = lines[i][:89]
			}
			if strings.Contains(lines[i], "!:::tick:::!") {
				continue
			}

			newBroadPayload := AssembleBM(newBroad, rowVal, colVal)
			broadcastContainer = append(broadcastContainer, newBroadPayload)
			if row >= 5 {
				row = 0
				rowVal = 0
			}else {
				row++
				rowVal += 4
				colVal = 153
			}
		}
		for i := 0;i < len(broadcastContainer);i++ {
			fmt.Print(broadcastContainer[i])
		}
		fmt.Print("\033[26;53H\n")

}

func doWatch(input string, play Player, fileChange chan bool) string {
	var broadcastContainer []string
	var do bool
	do = false
	select {
	case do = <-fileChange:
	}
	inputList := strings.Split(input, ":")

	if strings.Contains(input, "!:::tick:::!") {
		fmt.Println("\033[48:2:200:0:0mERROR\033[0m")
		return ""

		//do nothing but draw messages already there
	}
	if inputList[0] == "broadcast" || do {
		broadcastContainer = nil
		if do {
			fmt.Println("DOING")
		}
		file, err := os.Open("../pot/broadcast")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		contents, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		var lines []string
		lines = nil
		
		lines = strings.Split(string(contents), "\n")
		lineIn := strings.Split(string(contents), "\n")	
		if len(lines) >= 21 {
			lines = nil
			for i := len(lineIn)-1;i > len(lineIn)-21;i-- {
				lineIn[i] = strings.ReplaceAll(lineIn[i], "broadcast:", "")
				lines = append(lines, lineIn[i])
			}
		}
//			var broadcastContainer []Broadcast
		col := 0
		row := 0
		colVal := 53
		rowVal := 0
		for i := 0;i < len(lines);i++ {
			if strings.Contains(lines[i], "!:::tick:::!") {
				continue
			}
			var newBroad Broadcast
			newBroad.Payload.Name = play.Name
			newBroad.Payload.Game = "snowcrash.network"
			newBroad.Payload.Message = lines[i]
			if len(newBroad.Payload.Message) > 24 {
				newBroad.Payload.Message = lines[i][:24]
			}
			newBroadPayload := AssembleBroadside(newBroad, rowVal, colVal)
			broadcastContainer = append(broadcastContainer, newBroadPayload)
			if row >= 5 {
				row = 0
				rowVal = 0
			}
			if col < 3 {
				col++
				colVal += 30
			}else {
				row++
				rowVal += 4
				col = 0
				colVal = 53
			}
		}
		//log.Print(string(contents))
	}
	for i := 0;i < len(broadcastContainer);i++ {
		fmt.Print(broadcastContainer[i])
	}
	fmt.Print("\033[26;53H\n")

	return ""
}
