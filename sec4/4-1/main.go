package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

func main() {
	f := flag.String("wait", "1", "wait time[sec]")
	flag.Parse()
	t, err := strconv.Atoi(*f)
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Now())
	timer := time.After(time.Duration(t) * time.Second)
	end := <-timer
	fmt.Println(end)
}
