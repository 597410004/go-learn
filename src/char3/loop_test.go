package char3

import "testing"

// while true
func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 15 {
		n++
		t.Log(n)
	}
}

// if 包含多个条件,第一个为变量,第二个为condition(必须为布尔值)
func TestCondition(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1**1")
	}
}

// switch测试多个常量
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 3:
			t.Log("0,3")
		case 1, 2:
			t.Log("1,2")
		default:
			t.Log("default")
		}
	}
}

// switch条件测试
func TestSwitchMultiCase2(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("0,3")
		case i%3 == 0:
			t.Log("1,2")
		default:
			t.Log("default")
		}
	}
}
