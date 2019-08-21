package char16

import (
	"fmt"
	"sync"
	"testing"
)
// 每次Gc 会被清掉
// 适用于复用,协诚安全(会有锁的开销),需要自己管理生命周期的资源优化
func TestSysncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("创建一个新对象")
			return 10
		},
	}
	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(100)
	pool.Put(100)
	pool.Put(100)
	v1, _ := pool.Get().(int)
	fmt.Println(v1)

}
