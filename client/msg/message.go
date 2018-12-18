package msg

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	Conn     net.Conn
	Response string
	Active   bool
}

func (s *Client) Connect() (err error) {
	s.Conn, err = net.Dial("tcp", "127.0.0.1:3333")
	s.Active = true
	return
}

func (s *Client) Communicate(msg string) (err error) {
	fmt.Fprintf(s.Conn, msg+"\n")
	resp, err := bufio.NewReader(s.Conn).ReadString('\n')
	if err != nil {
		return
	}
	s.Response = resp[:strings.IndexByte(resp, '\n')]
	return
}
