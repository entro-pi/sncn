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

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your command")
	go watch()
	for scanner.Scan() {

		input := scanner.Text()
		//Should probably do some error checking before
		//passing it along
		doInput(input)
		fmt.Print("Enter your command")

	}
}

func doInput(input string) {
	connection := getConnectionString()
	conn, err := amqp.Dial(connection)

	failOnError(err, "Failed to connect to RabbitMQ")

	defer conn.Close()

	ch, err := conn.Channel()

	failOnError(err, "Failed to open a channel")

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"input", //name
		true, // durable
		true, //delete when used
		false, //exclusive
		false, //no-wait
		nil, //arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := input
	err = ch.Publish(
	"", //exchange
	q.Name, // routing key
	false, //mandatory
	false, //immediate
	amqp.Publishing {
		ContentType: "text/plain",
		Body: []byte(body),
	})
	log.Printf(" [x] Sent %s", body)
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

        q, err := ch.QueueDeclare(
                "input", //name
                true, // durable
                true, //delete when used
                false, //exclusive
                false, //no-wait
                nil, //arguments
        )
        failOnError(err, "Failed to declare a queue")
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
				log.Printf("Received a message: %s", d.Body)
				f, err := os.Open("../pot/broadcast")
				if err != nil {
					panic(err)
				}
				defer f.Close()
				f.WriteString("\n")
			}

		}()
		<-forever


	}
}

func watch() {
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
	            log.Println("event:", event)
	            if event.Op&fsnotify.Write == fsnotify.Write {
	                log.Println("\033[48:2:150:0:150mmodified file:", event.Name,"\033[0m")
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

	        case err, ok := <-watcher.Errors:
	            if !ok {
	                return
	            }
	            log.Println("error:", err)
		default:
			for i := 0;i < len(broadcastContainer);i++ {
				fmt.Print(broadcastContainer[i])
			}
	        }
	    }
	}()

	err = watcher.Add("../pot")
	if err != nil {
	    log.Fatal(err)
	}
	<-done
}
