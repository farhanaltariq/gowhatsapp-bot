package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"hanz/ai/utils"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
	_ "modernc.org/sqlite"
)

func main() {
	dbLog := waLog.Stdout("Database", "DEBUG", false)
	store, err := sqlstore.New("sqlite", "file:examplestore.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}

	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := store.GetFirstDevice()
	if err != nil {
		panic(err)
	}

	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)

	// Use a channel to receive signals for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start the database connection concurrently
	go func() {
		err = utils.ConnectDB(client)
		if err != nil {
			panic(err)
		}
	}()

	// Start the event handling concurrently
	go func() {
		client.AddEventHandler(func(evt interface{}) {
			// set to true to not send result to client
			utils.EventHandler(client, evt, false)
		})
	}()

	// run healthcheck endpoint
	go func() {
		http.HandleFunc("/", utils.HealthCheck)
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err)
		}
		fmt.Println("Healthcheck endpoint running at port :8080")
	}()

	<-stop // Wait for the stop signal

	client.Disconnect()
}
