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
	// Why is client protocol udp4 not just udp
	conn, err := net.Dial("udp4", ServerAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Sending to server")
	_, err = conn.Write([]byte("Hello from Client"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Receiving from server")
	buffer := make([]byte, BufferSize)
	length, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Received: %s\n", string(buffer[:length]))
}
