package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

const ServerAddr = "localhost:8888"

// func isGZipAcceptable(request *http.Request) bool {
// 	return strings.Contains(strings.Join(request.Header["Accept-Encoding"], ","), "gzip")
// }

var contents = []string{
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("Timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))
		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")
		// 	response := http.Response{
		// 		StatusCode: 200,
		// 		ProtoMajor: 1,
		// 		ProtoMinor: 1,
		// 		Header:     make(http.Header),
		// 	}
		// 	if isGZipAcceptable(request) {
		// 		content := "Hello World (gzipped)\n"
		// 		var buffer bytes.Buffer
		// 		writer := gzip.NewWriter(&buffer)
		// 		io.WriteString(writer, content)
		// 		writer.Close()
		// 		response.Body = io.NopCloser(&buffer)
		// 		response.ContentLength = int64(buffer.Len())
		// 		response.Header.Set("Content-Encoding", "gzip")
		// 	} else {
		// 		content := "Hello World\n"
		// 		response.Body = io.NopCloser(strings.NewReader(content))
		// 		response.ContentLength = int64(len(content))
		// 	}
		// 	response.Write(conn)
	}
}

func main() {
	listener, err := net.Listen("tcp", ServerAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at", ServerAddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}
