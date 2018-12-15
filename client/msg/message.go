package msg

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type Client struct {
	Conn     net.Conn
	Response string
}

func (s *Client) Conn() (err error) {
	s.Conn, err = net.Dial("tcp4", "127.0.0.1:3333")
	return
}

func (s *Client) Communicate(msg string) (err error) {
	fmt.Fprintf(s.Conn, msg)
	resp, err := bufio.NewReader(conn).ReadString('\n')
	if (err != nil) && (err != io.EOF) {
		return
	}
	s.Response = resp
}
