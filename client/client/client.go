package client

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	conn net.Conn
}

func NewClient(network, address string) Client {
	conn, err := net.Dial(network, address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return Client{
		conn: conn,
	}
}

// Start запуск приложения
func (c Client) Start() {
	//var disconnect bool
	go c.clientReader()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "disconnect" {
			c.conn.Close()
			break
		}
		go c.serverWriter(scanner.Bytes())
	}
}

// serverWriter отправляет сообщение на сервер
func (c Client) serverWriter(msg []byte) {
	_, err := c.conn.Write(msg)

	if err != nil {
		fmt.Println(err)
	}
}

// clientReader выводит на экран все сообщения от сервера
func (c Client) clientReader() {
	buffer := make([]byte, 256)
	for {
		n, err := c.conn.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}
