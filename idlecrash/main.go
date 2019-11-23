package main

import (
	"os"
	"io/ioutil"
	"log"
	"github.com/fsnotify/fsnotify"
)

func main() {


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
