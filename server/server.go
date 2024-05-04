package server

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/m-d-nabeel/cmd-line-message-app/pkg/chat"
)

type Client struct {
	Username   string
	Connection net.Conn
}

type Server struct {
	listener net.Listener
	clients  []Client
}

func NewServer(addr string) (*Server, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Server{listener: listener}, nil
}

func (s *Server) Close() {
	s.listener.Close()
	for _, conn := range s.clients {
		conn.Connection.Close()
	}
}

func (s *Server) Start() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		clnt := Client{Connection: conn}
		s.clients = append(s.clients, clnt)
		go s.handleClient(&clnt)
	}
}

func (s *Server) handleClient(clnt *Client) {
	defer clnt.Connection.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		chat.WriteMessage(clnt.Connection, "Welcome to the chat app!\nPlease enter your username:")
		username, err := chat.ReadMessage(clnt.Connection)
		if err != nil {
			fmt.Println("Error reading username:", err)
			return
		}
		clnt.Username = username[:len(username)-1]
		fmt.Println("New user joined:", clnt.Username)
		chat.WriteMessage(clnt.Connection, "Welcome, "+clnt.Username+"!!!")
		chat.WriteMessage(clnt.Connection, "\n--------------------------")
		wg.Done()
	}()
	wg.Wait()
	for {
		msg, err := chat.ReadMessage(clnt.Connection)
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}
		if msg == chat.ExitMessage {
			fmt.Println("User left:", clnt.Username)
			s.RemoveClient(clnt)
			return
		}
		s.broadcastMessage(msg, clnt)
	}
}

func (s *Server) broadcastMessage(msg string, sender *Client) {
	log.Println("Broadcasting message:", msg[:len(msg)-1])
	for _, clnt := range s.clients {
		if clnt.Connection != sender.Connection {
			_, err := clnt.Connection.Write([]byte(sender.Username + ": " + msg[:len(msg)-1]))
			if err != nil {
				fmt.Println("Error broadcasting message:", err)
			}
		}
	}
}

func (s *Server) RemoveClient(clnt *Client) {
	for i, client := range s.clients {
		if client.Connection == clnt.Connection {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
			break
		}
	}
}
