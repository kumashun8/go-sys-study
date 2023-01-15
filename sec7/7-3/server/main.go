package main

import (
	"fmt"
	"net"
	"time"
)

const (
	Interval      = 10 * time.Second
	MulticastAddr = "224.0.0.1:9999"
)

func main() {
	fmt.Println("Start tick server at", MulticastAddr)
	conn, err := net.Dial("udp", MulticastAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	start := time.Now()
	// Truncate: 時刻を10秒区切りに
	// Add+Sub: 10秒単位の端数を導出, その分waitすればちょうどの時刻からtickerが始まる=>時報
	wait := start.Truncate(Interval).Add(Interval).Sub(start)
	time.Sleep(wait)
	ticker := time.Tick(Interval)
	for now := range ticker {
		conn.Write([]byte(now.String()))
		fmt.Println("Tick:", now.String())
	}
}
