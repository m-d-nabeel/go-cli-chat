package server

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/m-d-nabeel/cmd-line-message-app/pkg/chat"
)

type Message struct {
	Time     string
	Username string
	Content  string
}
type Client struct {
	Username   string
	Connection net.Conn
	Messages   []Message
}

type Server struct {
	listener net.Listener
	clients  []*Client
	mutex    sync.Mutex
}

func (s *Server) GetClients() []*Client {
	return s.clients
}

func (s *Server) GetClient(username string) *Client {
	for _, c := range s.clients {
		if strings.EqualFold(c.Username, username) {
			return c
		}
	}
	return nil
}

func (s *Server) AddClient(clnt *Client) {
	s.mutex.Lock()
	s.clients = append(s.clients, clnt)
	s.mutex.Unlock()
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
		go s.handleClient(conn)
	}
}

func (s *Server) handleClient(conn net.Conn) {
	defer conn.Close()
	// Ask for username
	chat.WriteMessage(conn, "Welcome to the chat app!\nPlease enter your username:")
	username, err := chat.ReadMessage(conn)
	if err != nil {
		fmt.Println("Error reading username:", err)
		return
	}
	username = strings.TrimSpace(username)
	fmt.Println("New user joined:", username)
	chat.WriteMessage(conn, "Welcome, "+username+"!!!")
	chat.WriteMessage(conn, "\n--------------------------")
	clnt := Client{
		Username:   strings.TrimSpace(username),
		Connection: conn,
		Messages:   []Message{},
	}
	s.mutex.Lock()
	s.clients = append(s.clients, &clnt)
	s.mutex.Unlock()
	// Handle messages
	for {
		msg, err := chat.ReadMessage(clnt.Connection)
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}
		switch msg {
		case chat.ExitMessage:
			fmt.Println("User left:", clnt.Username)
			s.RemoveClient(&clnt)
			return
		case chat.GetLog:
			chat.WriteMessage(clnt.Connection, "\n--------Chat log--------\n")
			for _, msg := range clnt.Messages {
				chat.WriteMessage(clnt.Connection, msg.Time+": "+msg.Username+": "+msg.Content+"\n")
			}
			chat.WriteMessage(clnt.Connection, "--------------------------\n")
			continue
		case chat.GetUsers:
			chat.WriteMessage(clnt.Connection, "\n--------Users in chat--------\n")
			s.mutex.Lock()
			for _, c := range s.clients {
				chat.WriteMessage(clnt.Connection, c.Username+"\n")
			}
			s.mutex.Unlock()
			chat.WriteMessage(clnt.Connection, "------------------------------\n")
			continue
		}

		if strings.HasPrefix(msg, chat.Private) {
			parts := strings.Split(msg, " ")
			if len(parts) < 2 {
				chat.WriteMessage(clnt.Connection, "Invalid private message format. Please use: /private username message\n")
				continue
			}
			otherUser := parts[1]
			privateMessage := strings.Join(parts[2:], " ")
			privateMessage = strings.TrimSpace(privateMessage)
			s.sendPrivateMessage(&clnt, otherUser, privateMessage)
			continue
		}
		fmtMsg := Message{
			Username: clnt.Username,
			Content:  msg[:len(msg)-1],
			Time:     time.Now().Format("2006-01-02 15:04:05"),
		}
		clnt.Messages = append(clnt.Messages, fmtMsg)
		s.broadcastMessage(msg, &clnt)
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
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for i, client := range s.clients {
		if client.Connection == clnt.Connection {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
			break
		}
	}
}

func (s *Server) sendPrivateMessage(sender *Client, receiver string, msg string) {
	receiver = strings.TrimSpace(receiver)
	for _, c := range s.clients {
		if strings.EqualFold(c.Username, receiver) {
			chat.WriteMessage(c.Connection, sender.Username+": "+msg)
			return
		}
	}
	chat.WriteMessage(sender.Connection, "User not found\n")
}
