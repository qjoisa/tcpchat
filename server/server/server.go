package server

import (
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

type Server struct {
	Listener net.Listener
}

// NewServer принимает протокол и адресс и создает сервер
func NewServer(network, address string) Server {
	l, err := net.Listen(network, address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return Server{
		Listener: l,
	}
}

// Start запускает сервер
func (s Server) Start() error {
	go s.broadcaster()
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			return err
		}
		go s.handleConn(conn)
	}
}

// broadcaster рассылает входящие сообщения всем клиентам
// следит за подключением и отключением клиентов
func (s Server) broadcaster() {
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
					go s.writeMessages(k.conn, []byte(msg))
				}
			}
		}
	}
}

// handleConn обрабатывает действия клиента
func (s Server) handleConn(conn net.Conn) {
	ch := make(chan string)
	go s.clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{conn: conn, name: who, ch: ch}

	entering <- cli
	ch <- "You are " + who
	messages <- who + " has arrived"

	s.clientReader(conn)

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

// clientReader читает входящие сообщения от клиента
func (s Server) clientReader(conn net.Conn) {
	buffer := make([]byte, 256)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String()+":", err)
			return
		}
		messages <- conn.RemoteAddr().String() + ": " + string(buffer[:n])
	}
}

// clientWriter отправляет сообщения текущему клиенту
func (s Server) clientWriter(conn net.Conn, ch <-chan string) {
	for {
		_, err := conn.Write([]byte(<-ch))
		if err != nil {
			fmt.Println(err)
		}
	}
}

// writeMessages откправляет данные в чат
func (s Server) writeMessages(conn net.Conn, msg []byte) {
	_, err := conn.Write(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
}
