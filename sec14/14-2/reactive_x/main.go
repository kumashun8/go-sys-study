package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	// observableを作成
	emitter := make(chan rxgo.Item)
	observable := rxgo.FromChannel(emitter, rxgo.WithErrorStrategy(rxgo.ContinueOnError))
	observable.Map(func(_ context.Context, i interface{}) (interface{}, error) {
		line := i.(string)
		if strings.HasPrefix(line, "func ") {
			fmt.Println(line)
			return line, nil
		}
		return nil, nil
	}).DoOnError(func(err error) {
		fmt.Printf("Encounted error: %v\n", err)
	})

	// subscribe
	sub := observable.Observe()

	// observableに値を送信
	go func() {
		fmt.Println("call go func")
		content, err := os.ReadFile("main.go")
		if err != nil {
			emitter <- rxgo.Error(err)
		} else {
			for _, line := range strings.Split(string(content), "\n") {
				emitter <- rxgo.Item{V: line}
			}
		}
		close(emitter)
	}()

	<-sub
}
