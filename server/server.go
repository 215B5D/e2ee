package main

import (
	"e2ee/server/src/socket"
)

func main() {
	socket.Listen("localhost", "1337")
}
