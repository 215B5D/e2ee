package socket

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

var CONNS []net.Conn

type Message struct {
	Encrypted bool   `json:"encrypted"`
	Content   string `json:"content"`
}

func Listen(host, port string) {
	listener, err := net.Listen("tcp", host+":"+port)

	if err != nil {
		fmt.Printf("(socket) Unable to listen (%s:%s)\n", host, port)
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("(socket) Succesfully bound, awaiting incoming connections!")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("(socket) Unable to accept incoming connection!")
			continue
		}

		go Handle(conn)
	}
}

func Handle(conn net.Conn) {
	fmt.Println("(socket) Succesfully accepted incoming connection!")

	CONNS = append(CONNS, conn)

	fmt.Printf("(socket) There are currently %d active connection(s)\n", len(CONNS))

	Broadcast(conn, Message{
		Encrypted: false,
		Content:   "User has joined the chat room!",
	})

	for {
		buf := make([]byte, 1024)

		size, err := conn.Read(buf)

		if err != nil {
			fmt.Println("(socket) Succesfully deleted outgoing connection!")

			CONNS = RemoveConn(conn, CONNS)

			Broadcast(conn, Message{
				Encrypted: false,
				Content:   "User has left the chat room!",
			})

			return
		}

		Broadcast(conn, Message{
			Encrypted: true,
			Content:   string(buf[:size]),
		})
	}
}

func Broadcast(current net.Conn, message Message) {
	for _, conn := range CONNS {
		if conn.RemoteAddr() == current.RemoteAddr() {
			continue
		}

		data, err := json.Marshal(message)

		if err != nil {
			fmt.Println("(socket) Unable to marshal JSON!")
			return
		}

		conn.Write(data)
	}
}

func RemoveConn(conn net.Conn, conns []net.Conn) []net.Conn {
	for i, item := range conns {
		if item.RemoteAddr() != conn.RemoteAddr() {
			continue
		}

		conns = append(CONNS[:i], CONNS[i+1:]...)
	}

	return conns
}
