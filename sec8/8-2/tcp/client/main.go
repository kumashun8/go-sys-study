package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
)

const ServerAddr = "http://localhost:8888"

func main() {
	conn, err := net.Dial("unix", filepath.Join(os.TempDir(), "unixdomainsocket-sample"))
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("get", ServerAddr, nil)
	if err != nil {
		panic(err)
	}
	request.Write(conn)
	reqsponse, err := http.ReadResponse(bufio.NewReader(conn), request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(reqsponse, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dump))
}
