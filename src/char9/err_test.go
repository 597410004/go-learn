package char9

import (
	"errors"
	"testing"
)
// go错误机制处理
var LessThanError = errors.New("n should not less than 2")
var GreaterThanError = errors.New("n should not greater than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanError
	}
	if n > 100 {
		return nil, GreaterThanError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
func TestFib(t *testing.T) {
	if v, err := GetFibonacci(99);err != nil {
		t.Error(err)
	}else{
		t.Log(v)
	}
}
