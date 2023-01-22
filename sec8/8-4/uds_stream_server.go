package main

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
)

const SocketFilename = "bench-unixdomainsocket-stream"

func UnixDomainSocketStreamServer() {
	path := filepath.Join(os.TempDir(), SocketFilename)
	os.Remove(path)
	listener, err := net.Listen("unix", path)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				panic(err)
			}
			_, err = httputil.DumpRequest(request, true)
			if err != nil {
				panic(err)
			}
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body:       io.NopCloser(strings.NewReader("Hello World\n")),
			}
			response.Write(conn)
			conn.Close()
		}()
	}
}
