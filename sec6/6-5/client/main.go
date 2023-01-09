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
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	var conn net.Conn = nil
	var err error
	requests := make([]*http.Request, 0, len(sendMessages))
	conn, err = net.Dial("tcp", ServerAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// リクエストだけ先に送る
	for i := 0; i < len(sendMessages); i++ {
		lastMessage := i == len(sendMessages)-1
		request, err := http.NewRequest("GER", "http://"+ServerAddr+"?message="+sendMessages[i], nil)
		if lastMessage {
			request.Header.Add("Connection", "close")
		} else {
			request.Header.Add("Connection", "keep-alive")
		}
		if err != nil {
			panic(err)
		}
		err = request.Write(conn)
		if err != nil {
			panic(err)
		}
		fmt.Println("send: ", sendMessages)
		requests = append(requests, request)
	}

	// レスポンスをまとめて受信
	reader := bufio.NewReader(conn)
	for _, request := range requests {
		response, err := http.ReadResponse(reader, request)
		if err != nil {
			panic(err)
		}
		dump, err := httputil.DumpResponse(response, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))
	}
}
