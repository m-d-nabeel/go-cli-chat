package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/m-d-nabeel/cmd-line-message-app/client"
)

func main() {
	port := flag.String("port", "8000", "Port number for the server")
	flag.Parse()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	client, err := client.NewClient()
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}
	defer client.Close()

	err = client.Connect("localhost" + ":" + *port)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	fmt.Println("Connected to server on port", *port)
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
