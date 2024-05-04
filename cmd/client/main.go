package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/m-d-nabeel/cmd-line-message-app/client"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	client, err := client.NewClient()
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}
	defer client.Close()

	err = client.Connect("localhost:8000")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	fmt.Println("Connected to server")
	go client.ReceiveMessages()
	go client.SendMessages()

	for range signalChan {
		fmt.Println("Closing connection")
		return
	}
}
