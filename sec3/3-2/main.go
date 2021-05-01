package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func makeRandFile() int64 {
	writer, err := os.Create("rand.txt")
	if err != nil {
		panic(err)
	}
	writeSize, err := io.CopyN(writer, rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	defer writer.Close()
	return writeSize
}

func main() {
	fileSize := makeRandFile()
	fmt.Fprintf(os.Stdout, "file size are %d bytes!\n", fileSize)
}
