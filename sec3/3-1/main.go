package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func copyFile(dst string, src string) int64 {
	oldFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	newFile, err := os.Create(dst)
	if err != nil {
		panic(err)
	}

	fileSize, err := io.Copy(newFile, oldFile)
	if err != nil {
		panic(err)
	}
	defer oldFile.Close()
	defer newFile.Close()
	return fileSize
}

func main() {
	defaultDstName := "new-" + time.Now().Format(time.UnixDate) + ".txt"
	src := flag.String("src", "", "source file name")
	dst := flag.String("dst", defaultDstName, "destination file name")
	flag.Parse()

	fileSize := copyFile(*dst, *src)
	fmt.Fprintf(os.Stdout, "%d bytes copied!\n", fileSize)
}
