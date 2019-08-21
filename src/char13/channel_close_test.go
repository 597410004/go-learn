package char13

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

}
func dataRecevier(ch chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
	}()

}
// 通道关闭继续添加值会报异常,通道关闭继续取值会再正常接受结束后再接受第一个值后退出
func TestChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataRecevier(ch, &wg)
	wg.Add(1)
	dataRecevier(ch, &wg)
	wg.Wait()
}
