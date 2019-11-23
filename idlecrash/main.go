package main

import (
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
		false, // durable
		false, //delete when used
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


func watch() {


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
			file, err := os.Open(event.Name)
			if err != nil {
				panic(err)
			}
			contents, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}
			log.Print(string(contents))


	        case err, ok := <-watcher.Errors:
	            if !ok {
	                return
	            }
	            log.Println("error:", err)
	        }
	    }
	}()

	err = watcher.Add("../pot")
	if err != nil {
	    log.Fatal(err)
	}
	<-done
}
