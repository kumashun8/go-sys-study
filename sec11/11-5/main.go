package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 0 {
		return
	}
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	state := cmd.ProcessState
	fmt.Printf("%s\n", state.String())
	fmt.Printf("  Pid: %d\n", state.Pid())
	fmt.Printf("  System: %v\n", state.SystemTime())
	fmt.Printf("  User: %v\n", state.UserTime())

	count := exec.Command("./bin/count")
	stdout, _ := count.StdoutPipe()
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Printf("(stdout) %s\n", scanner.Text())
		}
	}()
	err = count.Run()
	if err != nil {
		panic(err)
	}
}
