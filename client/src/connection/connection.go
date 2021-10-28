package connection

import (
	"e2ee/client/src/encryption"
	"e2ee/client/src/io"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Message struct {
	Encrypted bool   `json:"encrypted"`
	Content   string `json:"content"`
}

func Connect(host, port string) (net.Conn, error) {
	connection, err := net.Dial("tcp", host+":"+port)

	if err != nil {
		return nil, errors.New("unable to connect")
	}

	return connection, nil
}

func HandleConnection(conn net.Conn, key string) {
	fmt.Print("\n")
	go HandleInput(conn, key)

	for {
		input := io.Prompt("> ")

		str, err := encryption.Encrypt(key, []byte(input))

		if err != nil {
			fmt.Print("\r(!) Message wasn't able to be encrypted, please ensure your keys are 32 bytes!\r\n> ")
			continue
		}

		conn.Write([]byte(str))
	}
}

func HandleInput(conn net.Conn, key string) {
	for {
		buf := make([]byte, 1024)

		size, err := conn.Read(buf)

		if err != nil {
			fmt.Println(err)
			return
		}

		var message Message
		json.Unmarshal(buf[:size], &message)

		if !message.Encrypted {
			fmt.Print("\r(!) Message: " + message.Content + "\r\n> ")
			continue
		}

		str, err := encryption.Decrypt(key, message.Content)

		if err != nil {
			fmt.Print("\r(!) Message wasn't able to be decrypted, please ensure your keys are the same!\r\n> ")
			continue
		}

		fmt.Print("\rMessage: " + str + "\r\n> ")
	}
}
