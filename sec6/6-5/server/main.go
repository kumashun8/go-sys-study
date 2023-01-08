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

// 青空文庫: 羅生門より
// https://www.aozora.gr.jp/cards/000879/files/127_15260.html
var contents = []string{
	"ある日の暮方の事である。一人の下人が、羅生門の下で雨やみを待っていた。",
	"広い門の下には、この男のほかに誰もいない。ただ、所々丹塗にぬりの剥はげた、大きな円柱に、蟋蟀が一匹とまっている。",
	"羅生門が、朱雀大路にある以上は、この男のほかにも、雨やみをする市女笠や揉烏帽子が、もう二三人はありそうなものである。それが、この男のほかには誰もいない。",
	"何故かと云うと、この二三年、京都には、地震とか辻風とか火事とか饑饉とか云う災がつづいて起った。",
	"そこで洛中のさびれ方は一通りではない。",
	"旧記によると、仏像や仏具を打砕いて、その丹にがついたり、金銀の箔がついたりした木を、路ばたにつみ重ねて、薪の料しろに売っていたと云う事である。",
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

		fmt.Fprint(conn, strings.Join([]string{
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
