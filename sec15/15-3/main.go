package main

import (
	"fmt"
	"sync"
)

// MyPool is a wrapper of sync.Pool
type MyPool struct {
	sync.Pool
}

// NewMyPool is a constructor of MyPool
func NewMyPool(new func() interface{}) *MyPool {
	return &MyPool{
		Pool: sync.Pool{
			New: new,
		},
	}
}

// Get is a wrapper of sync.Pool.Get
func (p *MyPool) Get() string {
	res := p.Pool.Get()
	switch v := res.(type) {
	case string:
		return v
	case *string:
		return *v
	default:
		return "default"
	}
}

func main() {
	// Poolを作成。Newで新規作成時のコードを実装
	var count int
	pool := NewMyPool(func() interface{} {
		count++
		return fmt.Sprintf("created: %d", count)
	})

	// 追加した要素から受け取れる
	// プールが空だと新規作成
	// プールに追加する要素はポインタが推奨される
	// https://staticcheck.io/docs/checks#SA6002
	str1, str2 := "manualy added: 1", "manualy added: 2"
	pool.Put(&str1)
	pool.Put(&str2)
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}
