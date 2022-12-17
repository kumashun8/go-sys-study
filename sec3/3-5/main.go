package main

import (
	"fmt"
	"io"
	"os"
)

func myCopyN(dst io.Writer, src io.Reader, n int64) (int64, error) {
	lr := io.LimitReader(src, n)
	w, err := io.Copy(dst, lr)
	if w == n {
		return w, nil
	}
	if w < n && err == nil {
		return w, io.EOF
	}
	return 0, err

}

func main() {
	file, err := os.Open("hoge.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// check real func
	io.CopyN(os.Stdout, file, 5)
	fmt.Println()
	myCopyN(os.Stdout, file, 5)
}
