package main

import (
	"e2ee/client/src/connection"
	"e2ee/client/src/io"
	"e2ee/client/src/util"
	"fmt"
	"strings"
)

var (
	CONNECTED = false
	COMMANDS  = map[string]interface{}{
		"key":     key,
		"connect": connect,
	}
	KEY = ""
)

func main() {
	fmt.Print("         ,_---~~~~~----._         \n  _,,_,*^____      _____``*g*\"*, \n / __/ /'     ^.  /      \\ ^@q   f \n[  @f | @))    |  | @))   l  0 _/  \n \\`/   \\~____ / __ \\_____/    \\   \n  |           _l__l_           I   \n  }          [______]           I  \n  ]            | | |            |  \n  ]             ~ ~             |  \n  |           go chat          |   \n   |       github: 215B5D      |   \n\n")

	for !CONNECTED {
		args := strings.Split(io.Prompt("~> "), " ")

		if len(args) == 0 {
			continue
		}

		_, ok := COMMANDS[strings.ToLower(args[0])]

		if !ok {
			fmt.Printf("%s: unknown command\n", strings.ToLower(args[0]))
			continue
		}

		COMMANDS[strings.ToLower(args[0])].(func([]string))(args)
	}
}

func key(args []string) {
	if len(args) < 2 || (strings.ToLower(args[1]) != "get" && strings.ToLower(args[1]) != "set") {
		fmt.Println("key: invalid usage, please try `key <action (get/set)> <action>`")
		return
	}

	switch strings.ToLower(args[1]) {
	case "get":
		if len(KEY) == 0 {
			fmt.Println("key: please set a key first")
			return
		}

		fmt.Printf("key: your current key is: %s\n", KEY)
	case "set":
		if len(args) > 2 {
			KEY = args[2]
		} else {
			KEY = util.GenerateString(32)
		}

		fmt.Printf("key: your current key is: %s\n", KEY)
	}
}

func connect(args []string) {
	if len(args) < 2 {
		fmt.Println("connect: invalid usage, please try `connect <host> <port (1337)>`")
		return
	}

	if len(KEY) == 0 {
		fmt.Println("connect: please set a key before attempting to initiate a connection")
		return
	}

	host := args[1]
	port := "1337"

	if len(args) > 2 {
		port = args[2]
	}

	fmt.Printf("connect: connecting to %s:%s\n", host, port)

	conn, err := connection.Connect(host, port)

	if err != nil {
		fmt.Println("connect: unable to connect")
		return
	}

	fmt.Printf("connect: connected to %s:%s\n", host, port)
	CONNECTED = true

	connection.HandleConnection(conn, KEY)
}
