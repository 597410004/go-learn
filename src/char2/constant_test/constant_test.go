package constant_test

import "testing"

const (
	Moday = 2 + iota
	Tuesday
	Wednesday
)

// var s string 初始化为空字符串
func TestConstant(t *testing.T) {
	t.Log(Moday, Tuesday, Wednesday)
}
