package chat

import (
	"bufio"
	"encoding/json"
	"io"
	"net"
	"net/http"
)

const (
	EraseLine   = "\x1b[2K"
	MoveToStart = "\r"
	MoveOneUp   = "\x1b[1A"
	ExitMessage = "/exit\n"
	GetLog      = "/getlog\n"
	GetUsers    = "/getusers\n"
	Private     = "/private"
)

func ReadMessage(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	buf := make([]byte, 1024)
	n, err := reader.Read(buf)
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func WriteMessage(conn net.Conn, msg string) error {
	_, err := conn.Write([]byte(msg))
	return err
}

func GetCatFact() string {
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	var fact struct {
		Fact   string `json:"fact"`
		Length int    `json:"length"`
	}
	err = json.Unmarshal(body, &fact)
	if err != nil {
		return ""
	}
	return fact.Fact
}
