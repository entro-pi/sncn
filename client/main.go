package main

import (
	"strconv"
	"time"
	"context"
	"os/signal"
	"strings"
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
	"log"
	"github.com/gotk3/gotk3/gtk"
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

func handleBreak() {
	//handle signal interrupt
	ctx := context.Background()

	//trap ctrl-c and call cancel
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

}


func main() {
	go handleBreak()

	fileChange := make(chan bool)
	if len(os.Args) == 2 {
	if os.Args[1] == "--gui" {
		LaunchGUI(fileChange)
	}
	if os.Args[1] == "--1920x1080main" {
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
					doPlayer(input, play, os.Args[1])
					go watch(play, fileChange)
					go doInput(input, play, fileChange, whoList)
					go doWatch(input, play, fileChange)
			//		fmt.Print("Enter your command")
				}
				fmt.Print("\033[26;53H\n")

			}
		}
	}
	if os.Args[1] == "--4x3" {
                        scanner := bufio.NewScanner(os.Stdin)
                //      fmt.Print("Enter your command")
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
                                        doPlayer(input, play, os.Args[1])
                                        go watch(play, fileChange)
                                        go doInput(input, play, fileChange, whoList)
                                        go doWatch(input, play, fileChange)
                        //              fmt.Print("Enter your command")
                                }
                                fmt.Print("\033[26;53H\n")
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

func logout(playName string) {
	f, err := os.OpenFile("../pot/who", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	newF, err := os.OpenFile("../pot/newWho", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	longUser := strings.Split(string(content), "\n")
	for i := 1;i < len(longUser);i++ {
		if playName == longUser[i] {
			newF.WriteString("<<logout>>\n")
			continue
		}else if len(longUser[i]) > 1 && longUser[i] != playName {
			newF.WriteString(longUser[i]+"\n")
		}
	}
	f.Sync()
	newF.Sync()
	f.Close()
	newF.Close()
//	os.Remove("../pot/who")
	os.Rename("../pot/newWho", "../pot/who")
}

func who(playName string) []string {
	var oldPlayers []string
	
	time.Sleep(100*time.Millisecond)
	f, err := os.Open("../pot/who")
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	longUser := strings.Split(string(content), "\n")
	for i := 1;i < len(longUser);i++ {
		if playName == longUser[i] {
			fmt.Print("YOU\n")
			continue
		}else if len(longUser[i]) > 1 && longUser[i] != playName {
			fmt.Print(longUser[i]+"\n")
		}
		oldPlayers = append(oldPlayers, longUser[i])
	}
	f.Close()
	return oldPlayers
}

func doPlayer(input string, play Player, format string) {
	
	if format == "--1920x1080main" {
		play = decompEq(play)
		play = decompInv(play)
		

		DescribePlayer(play)
		describeEquipment(play)
		describeInventory(play)
	}
}

func doInput(input string, play Player, fileChange chan bool, whoList []string) {
	connection := getConnectionString()
	conn, err := amqp.Dial(connection)

	direct := false

	//Determine if we're sending to anyone in particular

	inputArray := strings.Split(input, ":")
	if len(inputArray) < 2 {
		inputArray = strings.Split(input, " ")
	}
	if inputArray[0] == "quit" {
		os.Exit(1)
	}
	tellTo := ""
	if inputArray[0] == "tell" {
		direct = true
		tellTo = inputArray[1]
	}
	/*for i := 0;i < len(whoList);i++ {
		fmt.Println(inputArray[1])
		if inputArray[0] == "tell" && inputArray[1] == whoList[i] {
			direct = true
			tellTo = inputArray[1]
			fmt.Println("\033[48:2:0:0:120m",direct, tellTo,"\033[0m")
			break
		}
	}*/


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
	body := "::SENDER::"+play.Name+"::SENDER::"
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
	}else if inputArray[0] == "broadcast" {
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
func doGUIInput(play Player, input string) {
	connection := getConnectionString()
	conn, err := amqp.Dial(connection)

	direct := false

	//Determine if we're sending to anyone in particular

	inputArray := strings.Split(input, "::SENDER::")
	tellToArray := strings.Split(input, "@")
	tellTo := ""
	if len(tellToArray) > 1 {
		direct = true
		tellTo = tellToArray[1]
	}else {
		tellTo = ""
	}
	/*for i := 0;i < len(whoList);i++ {
		fmt.Println(inputArray[1])
		if inputArray[0] == "tell" && inputArray[1] == whoList[i] {
			direct = true
			tellTo = inputArray[1]
			fmt.Println("\033[48:2:0:0:120m",direct, tellTo,"\033[0m")
			break
		}
	}*/


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
	body := "::SENDER::"+play.Name+"::SENDER::="
	for i := 0;i < len(inputArray);i++ {
		body += inputArray[i]+" "
	}
	if direct {
		body += "::=::SENDTO::"+tellTo+"::SENDTO::"
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
		body += "::=::SENDTO::ALL::SENDTO::"
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
//	return tell
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

//		select {
//		default:
		go func() {
			fmt.Println("Awaiting messages")
			for msg := range msgs {
				fmt.Println("Message!")
				message := string(msg.Body)

				if strings.Split(message, "::SENDTO::")[1] == play.Name {

					log.Printf("\033[38:2:150:150:0mReceived a tell: %s\033[0m", msg.Body)
						if !strings.Contains(message, "!:::tick:::!") {
							f, err := os.OpenFile("../pot/tells", os.O_APPEND|os.O_WRONLY, 0644)
							if err != nil {
								panic(err)
							}
							//strip the thingies out
	//						message = strings.ReplaceAll(message, "tell:", "\033[38:2:150:0:100mtell")
							_, err = f.WriteString(message+"::TIMESTAMP::"+time.Now().Weekday().String()+"-"+strconv.Itoa(time.Now().Hour())+":"+strconv.Itoa(time.Now().Minute())+"::TIMESTAMP::\n")
							if err != nil {
								panic(err)
							}
//							f.Sync()
							forever <- true
							f.Close()
						}

				}else if strings.Split(message, "::SENDTO::")[1] == "ALL" {
					log.Printf("\033[38:2:0:150:150mReceived a message: %s\033[0m", msg.Body)
					if !strings.Contains(message, "!:::tick:::!") {
						f, err := os.OpenFile("../pot/broadcast", os.O_APPEND|os.O_WRONLY, 0644)
						if err != nil {
							panic(err)
						}
						//strip the thingies out
						_, err = f.WriteString(message+"::TIMESTAMP::"+time.Now().Weekday().String()+"-"+strconv.Itoa(time.Now().Hour())+":"+strconv.Itoa(time.Now().Minute())+"::TIMESTAMP::\n")
						if err != nil {
							panic(err)
						}
						f.Close()
						forever <- true

						//go doWatch(string(msg.Body), blank, fileChange)
					}
				}
			}
		}()
		<-forever
//		}
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
			if os.Args[1] == "--4x3" {
				drawTells(os.Args[1], play, 150, 24)
			}
			broadcastContainer = nil
	        	broadcastContainer = drawBroadcasts(os.Args[1], play, broadcastContainer)
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
			broadcastContainer = nil
			if event.Name == "../pot/broadcast" {
				broadcastContainer = drawBroadcasts(os.Args[1], play, broadcastContainer)
			}else if event.Name == "../pot/tells" && os.Args[1] == "--4x3" {
				drawTells(os.Args[1], play, 150, 24)
			}
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
func drawBroadcasts(format string, play Player, broadcastContainer []string) []string {
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
	colVal := 0
	rowVal := 0
	colValHolder := 0
	colNumber := 0
	rowNumber := 0
	if format == "--1920x1080main" {
		colVal = 53
		rowNumber = 5
		colValHolder = 53
		rowVal = 0
		colNumber = 3
	}else if format == "--4x3" {
		colVal = 0
		rowNumber = 5
		colValHolder = 0
		rowVal = 0
		colNumber = 4
	}
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
			if row >= rowNumber {
				row = 0
				rowVal = 0
			}
			if col < colNumber {
				col++
				colVal += 30
			}else {
				row++
				rowVal += 4
				col = 0
				colVal = colValHolder
			}
		}
		for i := 0;i < len(broadcastContainer);i++ {
			//fmt.Print(broadcastContainer[i])
		}
		//fmt.Print("\033[26;53H\n")

		//log.Print(string(contents))
	return broadcastContainer
}
func drawPlainBroadcasts(play Player) []string {
	var broadcastContainer []string
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
	for i := 0;i < len(lines);i++ {
			var newBroad Broadcast
			newBroad.Payload.Message = lines[i]
			newBroad.Payload.Name = play.Name
			newBroad.Payload.Game = "snowcrash.network"
			if len(newBroad.Payload.Message) > 500 {
				newBroad.Payload.Message = lines[i][:500]
			}
			if strings.Contains(lines[i], "!:::tick:::!") {
				continue
			}
			if len(newBroad.Payload.Message) > 1 {
				broadcastContainer = append(broadcastContainer, newBroad.Payload.Message)
			}
	}
	return broadcastContainer
}

func drawTells(format string, play Player, colVal int, rowVal int) []string {
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
	colValHolder := 0
	rowNumber := 0
	col := 0
	colNumber := 0
	if format == "--1920x1080main" {
		colVal = 150
		colValHolder = 150
	}else if format == "--4x3" {
		colVal = 0
		colValHolder = 0
		rowNumber = 6
		colNumber = 4
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
			if row >= rowNumber {
				row = 0
				rowVal = 24
			}else if col < colNumber {
				colVal += 30
				col++
			}else {
				rowVal += 4

				row++
				col = 0
				colVal = colValHolder
			}
		}
		for i := 0;i < len(broadcastContainer);i++ {
		//	fmt.Print(broadcastContainer[i])
		}
//		fmt.Print("\033[26;53H\n")
	return broadcastContainer
}


func drawPlainTells(play Player) []string {
	var broadcastContainer []string
	file, err := os.Open("../pot/tells")
	if err != nil {
		panic(err)
	}
	defer file.Close()

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
			if len(newBroad.Payload.Message) > 500 {
				newBroad.Payload.Message = lines[i][:500]
			}
			if strings.Contains(lines[i], "!:::tick:::!") {
				continue
			}
			if len(newBroad.Payload.Message) > 1 {
				broadcastContainer = append(broadcastContainer, newBroad.Payload.Message)
			}
		}
	return broadcastContainer
}
func paintOver(twoBuilder *gtk.Builder) {
        rows := 7
        cols := 4
        count := 0
	var broadcastContainer []string
	file, err := os.Open("../pot/paintOver")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tells, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(tells), "\n")
	for i := 0;i < len(lines);i++ {
			broadcastContainer = append(broadcastContainer, lines[i])
	}
        for r := 0;r < rows;r++ {
                for c := 0;c < cols;c++ {
                        count++
                        if count >= len(broadcastContainer) {
                                count = len(broadcastContainer)-1
                        }
                        messageName := fmt.Sprint("message"+strconv.Itoa(count))
                        messageUncast, err := twoBuilder.GetObject(messageName)
                        if err != nil {
                                panic(err)
                        }
                        message := messageUncast.(*gtk.Label)
                        message.SetText(lines[count])
                }
        }

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
		broadcastContainer = drawBroadcasts(os.Args[1], play, broadcastContainer)
		//log.Print(string(contents))
	}
	for i := 0;i < len(broadcastContainer);i++ {
		fmt.Print(broadcastContainer[i])
	}
	fmt.Print("\033[26;53H\n")

	return ""
}
