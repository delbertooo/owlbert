package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"./pinlike"
	"./webhook"
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

	http.HandleFunc(webhooksMapping, webhooksHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func webhooksHandler(w http.ResponseWriter, r *http.Request) {
	configName := r.URL.Path[len(webhooksMapping):]
	metaData, err := webhook.ReadGitlabHookMetaData(r)
	if err != nil {
		log.Println("Could not read request body.")
		return
	}
	action, err := webhook.LoadWebhookAction(configName, metaData.ObjectKind)
	if err != nil {
		log.Println("Error while loading action for "+configName, err)
		return
	}
	if action == nil {
		log.Printf("Could not find action '%v' in config '%v'.", metaData.ObjectKind, configName)
		return
	}
	if !action.IsAuthorized(r) {
		log.Printf("Denied unauthorized access to config '%v'.", configName)
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	action.Run(pin)
}
