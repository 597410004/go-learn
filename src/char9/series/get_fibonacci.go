package series

import (
	"errors"
	"fmt"
)

var LessThanError = errors.New("n should not less than 2")
var GreaterThanError = errors.New("n should not greater than 100")
//测试导出,首字符必须大写
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
//init方法可以多个,按顺序执行,每次在执行方法前执行
func init() {
	fmt.Print("init")
}
func init() {
	fmt.Println("init2")
}