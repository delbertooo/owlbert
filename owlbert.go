package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"./config"
	"./pinlike"
)

var (
	pin   pinlike.PinLike
	mutex *sync.Mutex
)

const webhooksMapping = "/webhooks/"

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: owlbert <port>")
		os.Exit(1)
	}
	port := os.Args[1]

	var conn pinlike.Connection
	//conn = (pinlike.GpioConnection)
	conn = new(pinlike.StubConnection)
	if err := conn.Open(); err != nil {
		panic(err)
	}
	// Unmap gpio memory when done
	defer conn.Close()

	// Pin 10 here is exposed on the pin header as physical pin 19
	//pin = pinlike.NewGpioPinLike(10)
	pin = new(pinlike.StubPinLike)
	mutex = &sync.Mutex{}

	http.HandleFunc("/db-casino/push", handler)
	http.HandleFunc(webhooksMapping, webhooksHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	pin.High()
	time.Sleep(10 * time.Second)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func webhooksHandler(w http.ResponseWriter, r *http.Request) {
	configName := r.URL.Path[len(webhooksMapping):]
	action, err := config.LoadWebhookAction(configName, "push")
	if err != nil {
		log.Println("Unknown config: " + configName)
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	action.Run(pin)
}
