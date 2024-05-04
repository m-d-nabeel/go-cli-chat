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
	showOptions()
	go client.ReceiveMessages()
	go client.SendMessages()

	for range signalChan {
		fmt.Println("Closing connection")
		return
	}
}

func showOptions() {
	fmt.Println("Options:")
	fmt.Println("1. \"/exit\" to exit")
	fmt.Println("2. \"/getlog\" to get chat log")
	fmt.Println("3. \"/getusers\" to get users list")
	fmt.Println("4. \"/private [username] [message]\" to send private message")
	fmt.Println("5. Type message to send to all users")
	fmt.Println("6. Press Ctrl+C to exit")
	fmt.Println()
}
