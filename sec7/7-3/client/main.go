package main

import (
	"fmt"
	"net"
)

const (
	MulticastAddr = "224.0.0.1:9999"
	BufferSize    = 1500
)

func main() {
	fmt.Println("Listen tick server at", MulticastAddr)
	address, err := net.ResolveUDPAddr("udp", MulticastAddr)
	if err != nil {
		panic(err)
	}
	listener, err := net.ListenMulticastUDP("udp", nil, address)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := listener.Close(); err != nil {
			panic(err)
		}
	}()

	buffer := make([]byte, BufferSize)
	for {
		length, remoteAddress, err := listener.ReadFromUDP(buffer)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Server %v\n", remoteAddress)
		fmt.Printf("Now    %s\n", string(buffer[:length]))
	}
}
