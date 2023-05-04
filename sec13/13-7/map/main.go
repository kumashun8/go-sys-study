package main

import (
	"fmt"
	"sync"
)

func main() {
	smap := sync.Map{}

	smap.Store("hello", "world")
	smap.Store(1, 2)

	smap.Delete("test")

	// これは既に存在するので何もしない
	smap.LoadOrStore(1, 3)
	// これは存在しないので挿入
	smap.LoadOrStore(2, 4)

	smap.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %v, value: %v\n", key, value)
		return true
	})
}
