package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	name := "Hayato Okuma"
	age := 23
	height := 174.5

	fmt.Fprintf(os.Stdout, "Write with os.Stdout at %v\n", time.Now())
	fmt.Fprintf(os.Stdout, "My name is %s, %d years old, and my height is %f cm.\n", name, age, height)
}
