package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	_ "modernc.org/sqlite"
)

func eventHandler(client *whatsmeow.Client, evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		// if not from me and not empty
		msg := v.Message.ExtendedTextMessage.GetText()
		if msg == "" {
			msg = v.Message.GetConversation()
		}

		sender := v.Info.MessageSource.Chat
		senderName := v.Info.PushName
		if v.Info.IsFromMe || msg == "" {
			return
		}

		fmt.Println("\033[32mSender\t:", senderName, " | ", sender, "\033[0m")
		fmt.Println("\033[32mMessage\t:", msg, "\033[0m")
		// reply message
		protoMsg := &proto.Message{
			ExtendedTextMessage: &proto.ExtendedTextMessage{
				// text to be send to sender
				Text: &msg,
			},
		}
		client.SendMessage(context.Background(), sender, protoMsg)
	}
}

func connect(client *whatsmeow.Client) error {
	var err error
	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				// Render the QR code here
				// e.g. qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
				fmt.Println("QR code:", evt.Code)

				// render the QR code and wait for the user to scan it
				// then continue
				config := qrterminal.Config{
					Level:     qrterminal.L,
					Writer:    os.Stdout,
					BlackChar: qrterminal.WHITE,
					WhiteChar: qrterminal.BLACK,
					QuietZone: 1,
				}
				qrterminal.GenerateWithConfig(evt.Code, config)
				fmt.Println("Scan the QR code above")

			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err = client.Connect()
		if err != nil {
			return err
		}
	}
	return nil
}

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

	client.AddEventHandler(func(evt interface{}) {
		eventHandler(client, evt)
	})

	err = connect(client)
	if err != nil {
		panic(err)
	}

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}
