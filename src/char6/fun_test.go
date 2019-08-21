package main

import (
	"fmt"
	"testing"
	"time"
)

//别名
type Func func(op int) int

func timeSpend(in Func) Func {
	return func(op int) int {
		start := time.Now()
		// 类似装饰模式
		ret := in(op)
		fmt.Println("time spend:", time.Since(start).Seconds())
		return ret
	}
}
func Fun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}
func TestFun(t *testing.T) {
	tF := timeSpend(Fun)
	t.Log(tF(5))
}

// 可变参数
func Sum(ops ...int) int {
	ref := 0
	for _, op := range ops {
		ref += op
	}
	return ref
}
func TestSum(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5))
}

//延迟执行,defer 相当于final,在函数执行完毕前执行,即使函数出错也会执行
func Clear() {
	fmt.Println("清空资源")
}
// 相当于try catch
func Recover() {
	if err := recover(); err != nil {
		fmt.Println("revocer",err)
	}
}

func TestDef(t *testing.T) {
	defer func() {
		fmt.Println("清空资源")
	}()
	fmt.Println("--------")
	panic("err")
}
func TestRecover(t *testing.T) {
	defer Recover()
	fmt.Println("--------")
	panic("err")
}
