package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("waiting 5 seconds")
	after := time.After(5 * time.Second)
	<-after
	fmt.Println("done")

	fmt.Println("report each seconds")
	for now := range time.Tick(1 * time.Second) {
		fmt.Println("now: ", now.Round(time.Second))
	}
}
