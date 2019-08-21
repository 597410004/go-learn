package char13

import (
	"fmt"
	"testing"
	"time"
)

func SomeThing() string {
	time.Sleep(time.Second * 1)
	return "Done!"
}
func OtherThing() {
	fmt.Println("do otherthing")
	time.Sleep(time.Second * 3)
	fmt.Println("otherThing Done!")
}
func TestThing(t *testing.T) {
	fmt.Println(SomeThing())
	OtherThing()
}

// channel如果不设定buffer,则必须得到返回值才往下执行,即task is running 必须等待1秒以后才能执行taskexit
func AsyncService() chan string {
	retCh := make(chan string)
	//retCh := make(chan string,1)
	go func() {
		ret := SomeThing()
		fmt.Println("task is running...")
		// 将返回值放入channel中
		retCh <- ret
		fmt.Println("task exit")
	}()
	return retCh
}

// channel如果设定buffer,task is running立马执行taskexit
func AsyncServiceWithBuffer() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := SomeThing()
		fmt.Println("task is running...")
		// 将返回值放入channel中
		retCh <- ret
		fmt.Println("task exit")
	}()
	return retCh
}
func TestAsync(t *testing.T) {
	retCh := AsyncService()
	OtherThing()
	// 取出channel值
	fmt.Println(<-retCh)
}
func TestAsyncWithBuffer(t *testing.T) {
	retCh := AsyncServiceWithBuffer()
	OtherThing()
	// 取出channel值
	fmt.Println(<-retCh)
}
// 超时机制
func TestSelect(t *testing.T) {
	select {
	case retCh := <-AsyncServiceWithBuffer():
		fmt.Println("返回结果",retCh)
	case <-time.After(time.Second*5):
		fmt.Println("超时")

	}

	}

