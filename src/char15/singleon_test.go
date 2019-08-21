package char15

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

//单例
type Singleton struct {
}

var singleton *Singleton
var once sync.Once

func getSinglon() *Singleton {
	once.Do(func() {
		fmt.Println("Create obj")
		singleton = new(Singleton)
	})
	return singleton
}
func TestSinglon(t *testing.T) {
	var group sync.WaitGroup
	for i := 1; i < 10; i++ {
		group.Add(1)
		go func() {
			obj := getSinglon()
			defer func() {
				group.Done()
			}()
			// 输出内存地址
			fmt.Printf("%x\n",unsafe.Pointer(obj))
		}()
	}
	group.Wait()
}
