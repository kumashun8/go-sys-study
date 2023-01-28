package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const Filename = "textfile.txt"

func open() {
	file, err := os.Create(Filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.WriteString(file, "New file content\n")
}

func read() {
	file, err := os.Open(Filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("Read file:")
	io.Copy(os.Stdout, file)
}

func main() {
	open()
	read()

	f, _ := os.Create("file.txt")
	a := time.Now()
	f.Write([]byte("緑の怪獣"))
	b := time.Now()
	f.Sync()
	c := time.Now()
	f.Close()
	d := time.Now()
	fmt.Printf("Write: %v\n", b.Sub(a))
	fmt.Printf("Sync: %v\n", c.Sub(b))
	fmt.Printf("Close: %v\n", d.Sub(c))

	if len(os.Args) == 1 {
		fmt.Println("%s  [exec file name]", os.Args[0])
		os.Exit(1)
	}
	info, err := os.Stat(os.Args[1])
	if err == os.ErrNotExist {
		fmt.Printf("file not found: %s\n", os.Args[1])
	} else if err != nil {
		panic(err)
	}
	fmt.Println("FileInfo")
	fmt.Printf("  ファイル名: %v\n", info.Name())
	fmt.Printf("  サイズ: %v\n", info.Size())
	fmt.Printf("  変更日時: %v\n", info.ModTime())
	fmt.Println("Mode()")
	fmt.Printf("  ディレクトリ?: %v\n", info.Mode().IsDir())
	fmt.Printf("  読み書き可能な通常ファイル?: %v\n", info.Mode().IsRegular())
	fmt.Printf("  Unixのファイルアクセス権限ビット: %o\n", info.Mode().Perm())
	fmt.Printf("  モードのテキスト表現: %v\n", info.Mode().String())

	os.Remove(Filename)
}
