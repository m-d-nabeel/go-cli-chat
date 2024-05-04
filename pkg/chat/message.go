package chat

import (
	"bufio"
	"net"
)

const (
	EraseLine   = "\x1b[2K"
	MoveToStart = "\r"
	MoveOneUp   = "\x1b[1A"
	ExitMessage = "/exit\n"
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
