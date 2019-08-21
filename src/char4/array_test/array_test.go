package array_test

import "testing"

//数组定义
func TestArray(t *testing.T) {
	var arr [3]int
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [...]int{1, 2, 3, 4, 5, 8, 9}
	t.Log(arr[2], arr1[4], arr2[6])
}

//数组遍历
func TestArrayTravel(t *testing.T) {
	// 不确定长度 用...
	arr2 := [...]int{1, 2, 3, 4, 5, 8, 9}
	for idx, e := range arr2 {
		t.Log(idx, e)
	}
	// 使用占位符标记index,如果不使用 只用一个变量,则只能遍历index
	for _, e := range arr2 {
		t.Log(e)
	}
	for e := range arr2 {
		t.Log(e)
	}
}

//数组切片
func TestArraySection(t *testing.T) {
	arr2 := [...]int{1, 2, 3, 4, 5, 8, 9}
	//不可以使用负数,只包含结尾,不包含开头
	arr2_sec1 := arr2[:4]
	arr2_sec2 := arr2[2:]
	arr2_sec3 := arr2[:]
	arr2_sec4 := arr2[2:3]
	t.Log(arr2_sec1, arr2_sec2, arr2_sec3, arr2_sec4)

}
