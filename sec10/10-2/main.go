package main

import (
	"fmt"
	"sync"
	"syscall"
	"time"
)

type Filelock struct {
	l  sync.Mutex
	fd int
}

func NewFileLock(filename string) *Filelock {
	if filename == "" {
		panic("filename edded")
	}
	fd, err := syscall.Open(filename, syscall.O_CREAT|syscall.O_RDONLY, 0750)
	if err != nil {
		panic(err)
	}
	return &Filelock{fd: fd}
}

func (m *Filelock) Lock() {
	m.l.Lock()
	if err := syscall.Flock(m.fd, syscall.LOCK_EX); err != nil {
		panic(err)
	}
}

func (m *Filelock) Unlock() {
	if err := syscall.Flock(m.fd, syscall.LOCK_UN); err != nil {
		panic(err)
	}
	m.l.Unlock()
}

func main() {
	l := NewFileLock("main.go")
	fmt.Println("try locking...")
	l.Lock()
	fmt.Println("locked!", time.Now())
	time.Sleep(10 * time.Second)
	l.Unlock()
	fmt.Println("unlock.", time.Now())
}
