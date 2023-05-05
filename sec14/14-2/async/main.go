package main

import (
	"os"
)

func main() {
	inputs := make(chan []byte)

	go func() {
		a, _ := os.ReadFile("a.txt")
		inputs <- a
	}()

	go func() {
		b, _ := os.ReadFile("b.txt")
		inputs <- b
	}()
}
