package main

import (
	"io"
	"os"
	"os/exec"

	"github.com/kr/pty"
)

func main() {
	cmd := exec.Command("./bin/check")
	if len(os.Args) > 1 && os.Args[1] == "fake" {
		stdpty, stdtty, _ := pty.Open()
		defer stdtty.Close()
		cmd.Stdin = stdpty
		cmd.Stdout = stdpty
		errpty, errtty, _ := pty.Open()
		cmd.Stderr = errpty
		defer errtty.Close()
		go func() {
			io.Copy(os.Stdout, stdtty)
		}()
		go func() {
			io.Copy(os.Stderr, errtty)
		}()
	}
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
