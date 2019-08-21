package array

import (
	"testing"
)

/*
切片与数组区别:
1. 切片可以无限伸缩
2. 切片无法相互比较
3. 初始化时是否指定长度,指定长度为数组
*/
//切片声明 len()切片长度,cap()切片容量
func TestSlice(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	s2 = append(s2, 1)
	t.Log(s2, len(s2), cap(s2))
}

//切片增长容量变化
func TestGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 5; i++ {
		s = append(s, i)
		// len相当于数组默认元素个数,cap相当于数组容量个数,可无限延长,每次延长为上一次2倍容量
		t.Log(len(s), cap(s))
	}
}

//切片共享结构测试
func TestSliceShareMemory(t *testing.T) {
	year := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	year2 := year[1:5]
	// 容量为从1开始到结尾的共享容量
	t.Log(len(year2), cap(year2), year2)
	year2 = append(year2, 0)
	t.Log(len(year2), cap(year2), year2)
	t.Log(year)
	year2[2] = 0
	t.Log(year)
}
