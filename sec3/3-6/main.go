package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader
	// make a string "ASCII" by only 3 vars above.
	ammi := io.NewSectionReader(programming, 5, 4)
	c := io.LimitReader(computer, 1)
	s := io.LimitReader(system, 1)

	b := make([]byte, 6)
	n, _ := ammi.ReadAt(b, 0)
	m, _ := s.Read(b[n:])
	// AMMISC
	c.Read(b[n+m:])

	// ASCIIC
	b[1], b[2], b[4] = b[n+m-1], b[n+m], b[3]
	stream = strings.NewReader(string(b[:5]))
	io.Copy(os.Stdout, stream)
}
