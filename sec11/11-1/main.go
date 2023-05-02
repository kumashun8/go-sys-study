package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	path, _ := os.Executable()
	fmt.Printf("実行ファイル名: %s\n", os.Args[0])
	fmt.Printf("実行ファイルパス: %s\n", path)

	fmt.Printf("プロセスID: %d\n", os.Getpid())
	fmt.Printf("親プロセスID: %d\n", os.Getppid())

	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Printf("グループID: %d\nセッションID: %d\n", syscall.Getpgrp(), sid)

	fmt.Printf("ユーザーID: %d\n", os.Getuid())
	fmt.Printf("グループID: %d\n", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Printf("サブグループID %v\n", groups)

	fmt.Printf("実効ユーザーID: %d\n", os.Geteuid())
	fmt.Printf("実効グループID: %d\n", os.Getegid())

	wd, _ := os.Getwd()
	fmt.Println(wd)
}
