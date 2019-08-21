package main

import "testing"

// 初始化map
func TestMap(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m[3])
	t.Log(len(m))
	m2 := map[int]int{}
	m2[1] = 4
	t.Log(len(m2))
	m3 := make(map[int]int, 10)
	// 注意打印的区别:log 为正常打印,logf为格式打印;Print和Printf类似
	//格式化输出用%v,则会自动判断类型
	t.Logf("len m3=%v", len(m3))
}

// 判断map中是否存在key(go中map如果不存在该key,仍然会返回0)
func TestMapKeyExist(t *testing.T) {
	m := make(map[int]int, 10)
	t.Logf("map key is %v", m[3])
	if v, isExists := m[3]; isExists {
		t.Logf("map key is %v", v)
	} else {
		t.Logf("map value is not exist")
	}
	m[3] = 0
	t.Logf("map key is %v", m[3])
	if v, isExists := m[3]; isExists {
		t.Logf("map key is %v", v)
	} else {
		t.Logf("map value is not exist")
	}
}

// 遍历map
func TestTraveMap(t *testing.T) {
	m := map[int]string{1: "one", 2: "two", 3: "three"}
	for k, v := range m {
		t.Log(k, v)
	}
	for k := range m {
		t.Log(k)
	}
	for _, k := range m {
		t.Log(k)
	}
}

// map添加函数
func TestMapFactory(t *testing.T) {
	m:=map[int]func(number int)int{}
	m[1]= func(number int) int {
		return number
	}
	m[2]= func(number int) int {
		return number*number
	}
	t.Log(m[1](5),m[2](5))
}
