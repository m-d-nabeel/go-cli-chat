package client

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/m-d-nabeel/cmd-line-message-app/pkg/chat"
)

type Client struct {
	conn net.Conn
}

func NewClient() (*Client, error) {
	return &Client{}, nil
}

func (c *Client) Connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Client) Close() {
	c.conn.Write([]byte(chat.ExitMessage))
	c.conn.Close()
}

func (c *Client) ReceiveMessages() {
	for {
		msg, err := chat.ReadMessage(c.conn)
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}
		fmt.Println(msg)
	}
}
func (c *Client) SendMessages() {
	inputChan := make(chan string)

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			if len(input) == 1 {
				input = "[CAT FACT]: " + chat.GetCatFact()
			}
			fmt.Print(chat.MoveOneUp + chat.EraseLine + chat.MoveToStart + "You : " + input)
			inputChan <- input
		}
	}()

	for msg := range inputChan {
		err := chat.WriteMessage(c.conn, msg)
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}
