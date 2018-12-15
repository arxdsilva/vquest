package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp4", ":3333")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		fmt.Println("<<<<<<<")
		netData, err := bufio.NewReader(c).ReadString('\n')
		fmt.Println(string(netData), err, "<<<<<<<")
		if (err != nil) && (err != io.EOF) {
			fmt.Println(err)
			break
		}
		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		result := fmt.Sprintf("%s voce esta conectado.\n", string(netData))
		c.Write([]byte(string(result)))
	}
	c.Close()
}
