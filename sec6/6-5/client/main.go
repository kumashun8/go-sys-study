package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
)

const ServerAddr = "localhost:8888"

func main() {
	conn, err := net.Dial("tcp", ServerAddr)
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("GET", "http://"+ServerAddr, nil)
	if err != nil {
		panic(err)
	}
	request.Write(conn)
	response, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
