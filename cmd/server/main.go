package main

import (
	"flag"
	"fmt"

	"github.com/m-d-nabeel/cmd-line-message-app/server"
)

func main() {
	port := flag.String("port", "8000", "Port number for the server")
	flag.Parse()
	server, err := server.NewServer(":" + *port)
	if err != nil {
		fmt.Println("Error creating server:", err)
		return
	}
	defer server.Close()

	fmt.Println("Server listening on ", *port)
	server.Start()
}
