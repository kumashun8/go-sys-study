package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

const BufferSize = 1500

func main() {
	clientPath := filepath.Join(os.TempDir(), "unixdomainsocket-client")
	os.Remove(clientPath)
	fmt.Println("Server is running at " + clientPath)
	conn, err := net.ListenPacket("unixgram", clientPath)
	if err != nil {
		panic(err)
	}
	unixServerAddr, err := net.ResolveUnixAddr("unixgram", filepath.Join(os.TempDir(), "unixdomainsocket-server"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	log.Println("Sending to server")
	_, err = conn.WriteTo([]byte("Hello from Client"), unixServerAddr)
	if err != nil {
		panic(err)
	}
	log.Println("Receiving from server")
	buffer := make([]byte, BufferSize)
	length, _, err := conn.ReadFrom(buffer)
	if err != nil {
		panic(err)
	}
	log.Printf("Received: %s\n", string(buffer[:length]))
}
