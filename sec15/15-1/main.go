package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Page Size: ", os.Getpagesize())
}
