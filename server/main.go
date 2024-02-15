package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type client struct {
	conn net.Conn
	name string
	ch   chan<- string
}

var (
	// канал для всех входящих клиентов
	entering = make(chan client)
	// канал для сообщения о выходе клиента
	leaving = make(chan client)
	// канал для всех сообщений
	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// broadcaster рассылает входящие сообщения всем клиентам
// следит за подключением и отключением клиентов
func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case v := <-entering:
			clients[v] = true
		case v := <-leaving:
			clients[v] = false
		case msg := <-messages:
			for k, v := range clients {
				if v {
					writeMessages(k.conn, []byte(msg))
				}
			}
		}
	}
}

// handleConn обрабатывает действия клиента
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	cli := client{conn, who, ch}

	entering <- cli
	ch <- "You are " + who
	messages <- who + " has arrived"

	clientReader(conn)

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

// clientReader читает входящие сообщения от клиента
func clientReader(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- conn.RemoteAddr().String() + ": " + input.Text()
	}
}

// clientWriter отправляет сообщения текущему клиенту
func clientWriter(conn net.Conn, ch <-chan string) {
	for {
		_, err := conn.Write([]byte(<-ch))
		if err != nil {
			fmt.Println(err)
		}
	}
}

// writeMessages откправляет данные в чат
func writeMessages(conn net.Conn, msg []byte) {
	_, err := conn.Write(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
}
