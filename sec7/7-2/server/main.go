package main

import (
	"fmt"
	"net"
)

const (
	ServerAddr = "localhost:8888"
	BufferSize = 1500
)

func main() {
	fmt.Print("Server is running", ServerAddr)
	conn, err := net.ListenPacket("udp", ServerAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, BufferSize)
	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received from %v: %v\n", remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from Server"), remoteAddress)
		if err != nil {
			panic(err)
		}
	}
}
