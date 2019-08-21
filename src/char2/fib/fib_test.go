package fib

import "testing"

//常量定义
func TestFibList(t *testing.T) {
	//var a int=1
	// var a=1
	//var a,b=1
	a := 1
	b := 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		tmp := a
		t.Log(" ", b)
		a = b
		b = tmp + a
	}
}

// 一行定义多个变量
func TestExchange(t *testing.T) {
	a := 1
	b := 5
	a, b = b, a
	t.Log(a, b)
	v, g, h, k := 1, 2, 3, 4
	t.Log(v, g, h, k)
}
