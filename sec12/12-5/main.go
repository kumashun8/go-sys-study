package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	listener "github.com/lestrrat-go/server-starter/listener"
)

func main() {
	// シグナル初期化
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)

	// Server::Startからもらったソケットを確認
	listeners, err := listener.ListenAll()
	if err != nil {
		panic(err)
	}

	// ウェブサーバーをgoruntineで起動
	server := http.Server{
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "server pid: %d %v\n", os.Getpid(), os.Environ())
		}),
	}
	go server.Serve(listeners[0])

	// SIGTERMを受けたら終了する
	<-signals
	server.Shutdown(context.Background())
}
