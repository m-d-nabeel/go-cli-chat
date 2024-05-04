package main

import (
	"fmt"

	"github.com/m-d-nabeel/cmd-line-message-app/server"
)

func main() {
    server, err := server.NewServer(":8000")
    if err != nil {
        fmt.Println("Error creating server:", err)
        return
    }
    defer server.Close()

    fmt.Println("Server listening on :8000")
    server.Start()
}