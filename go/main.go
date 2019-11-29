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

	if len(os.Args) == 2 {
	if os.Args[1] == "--main" {
			scanner := bufio.NewScanner(os.Stdin)
		//	fmt.Print("Enter your command")
			fmt.Print("Initializing a player")
			play := InitPlayer("dorp", "norp")
		//	go actOn() //for receiving in Go
//			go watch(play)
			for scanner.Scan() {
				input := "broadcast:"+scanner.Text()
				//Should probably do some error checking before
				//passing it along
				doPlayer(input, play)
				doWatch(input, play)
				doInput(input)
		//		fmt.Print("Enter your command")

				fmt.Print("\033[26;53H\n")

			}
	}
	if os.Args[1] == "--secondary" {
			scanner := bufio.NewScanner(os.Stdin)
		//	fmt.Print("Enter your command")
			fmt.Print("Initializing a player")
			play := InitPlayer("dorp", "norp")
		//	go actOn() //for receiving in Go
			go watch(play)
			for scanner.Scan() {
				input := scanner.Text()
				//Should probably do some error checking before
				//passing it along
				play, input = parseInput(play, input)
				doPlayer(input, play)
				doWatch(input, play)
				doInput(input)
		//		fmt.Print("Enter your command")

				fmt.Print("\033[26;53H\n")

			}
		}
	}
}

func parseInput(play Player, input string) (Player, string) {

	fmt.Println("INPUT IS ",input)
	if strings.HasPrefix(input, "generate") {
		value := strings.Split(input, " ")[1]
		fmt.Println(value)
		switch value {
		case "1":
			fmt.Println("generating a tiara")
			object := InitObject()
			play.Inventory[0].Item = object
			play.Inventory[0].Number++
		default:
			input += "broadcast:"
		}
	}

	return play, input
}

func doPlayer(input string, play Player) {
	play = decompEq(play)
	play = decompInv(play)
	DescribePlayer(play)
	describeEquipment(play)
	describeInventory(play)
}

func doInput(input string) {
	connection := getConnectionString()
	conn, err := amqp.Dial(connection)

	left, right, both := false, false, false

	//Determine if we're sending to anyone in particular
	inputArray := strings.Split(input, ":")
	if len(inputArray) <= 2 {
		inputArray = append(inputArray, ":broadcasts")
	}
	if inputArray[1] == "left" {

		left = true
	}else if inputArray[1] == "right" {
		right = true
	}else {
		both = true
	}

	failOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"doot", //name
		true, // durable
		false, //delete when used
		false, //exclusive
		false, //no-wait
		nil, //arguments
	)
	qleft, err := ch.QueueDeclare(
		"left", //name
		true, // durable
		false, //delete when used
		false, //exclusive
		false, //no-wait
		nil, //arguments
	)
	qright, err := ch.QueueDeclare(
		"right", //name
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
	"broadcasts", //name
	"fanout", //type
	false, //durable
	false, //auto-delted
	false, //internal
	false, //no wait
	nil, //arguments
	)
	failOnError(err, "Failed to declare an exchange")
	err = ch.ExchangeDeclare(
	"broadcastsLeft", //name
	"direct", //type
	false, //durable
	false, //auto-delted
	false, //internal
	false, //no wait
	nil, //arguments
	)
	failOnError(err, "Failed to declare an exchange")
	err = ch.ExchangeDeclare(
	"broadcastsRight", //name
	"direct", //type
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
		"broadcasts",//exchange
		false,
		nil,
	)
	
	err = ch.QueueBind(
		qleft.Name, //queue name
		"left", //routing key
		"broadcastsLeft",//exchange
		false,
		nil,
	)
	
	err = ch.QueueBind(
		qright.Name, //queue name
		"right", //routing key
		"broadcastsRight",//exchange
		false,
		nil,
	)
	
	failOnError(err, "Failed to bind a queue")

	body := "broadcast:"+inputArray[2]
	if left {

		err = ch.Publish(
		"broadcastsLeft", //exchange
		"left", // routing key
		false, //mandatory
		false, //immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body: []byte(body),
		})
	}
	if right {
		err = ch.Publish(
		"broadcastsRight", //exchange
		"right", // routing key
		false, //mandatory
		false, //immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body: []byte(body),
		})
	}else if both {
		err = ch.Publish(
		"broadcasts", //exchange
		"", // routing key
		false, //mandatory
		false, //immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body: []byte(body),
		})

	}

//	fmt.Print("\033[26;53H\n")
//	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")

}


func actOn() {
        connection := getConnectionString()
        conn, err := amqp.Dial(connection)

        failOnError(err, "Failed to connect to RabbitMQ")

        defer conn.Close()

        ch, err := conn.Channel()

        failOnError(err, "Failed to open a channel")

        defer ch.Close()

	err = ch.ExchangeDeclare(
		"broadcasts",//name
		"fanout",//type
		true, //durable
		false, //auto-deleted
		false, //internal
		false, //no wait
		nil, //args
	)
	failOnError(err, "Failed to declare an exchange")

        q, err := ch.QueueDeclare(
                "doot", //name
                true, // durable
                false, //delete when used
                false, //exclusive
                false, //no-wait
                nil, //arguments
        )
        failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name, //queue name
		"", //routing key
		"broadcasts", //exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	forever := make(chan bool)
	for {

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
		go func() {
			for d := range msgs {
		//		fmt.Print("\033[26;53H\n")
		//		log.Printf("Received a message: %s", d.Body)
				message := string(d.Body)
				if strings.HasPrefix(message, "broadcast:") {
					var blank Player
					if !strings.Contains(message, "!:::tick:::!") {
						doWatch(string(d.Body), blank)
					}else {
						doWatch("!:::tick:::!", blank)
					}
				}
				if err != nil {
					panic(err)
				}
			}

		}()
		<-forever


	}
}

func watch(play Player) {
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
	        case event, ok := <-watcher.Events:
	            if !ok {
	                return
	            }
	           // fmt.Print("\033[26;53H\n")
		  //  log.Print("event:", event)
	            if event.Op&fsnotify.Write == fsnotify.Write {
	        //        log.Print("\033[48:2:150:0:150mmodified file:", event.Name,"\033[0m")
	            }
		if event.Name == "../pot/broadcast" {
			broadcastContainer = nil
			file, err := os.Open(event.Name)
			if err != nil {
				panic(err)
			}
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
func doWatch(input string, play Player) string {
	var broadcastContainer []string

	inputList := strings.Split(input, ":")

	if strings.Contains(input, "!:::tick:::!") {
		fmt.Println("\033[48:2:200:0:0mERROR\033[0m")
		return ""

		//do nothing but draw messages already there
	}
	if inputList[0] == "broadcast" {
		broadcastContainer = nil

		file, err := os.Open("../pot/broadcast")
		if err != nil {
			panic(err)
		}
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
