package char17

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatStringByAdd(t *testing.T) {
	arr := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, value := range arr {
		ret += value
	}
	assert.Equal(t, "12345", ret)
}
func TestConcatStringByBuffer(t *testing.T) {
	arr := []string{"1", "2", "3", "4", "5"}
	var buf bytes.Buffer
	for _, value := range arr {
		buf.WriteString(value)
	}
	assert.Equal(t, "12345", buf.String())
}

// 用来测试性能,必须以Benchmark开头
func BenchmarkConcatStringByAdd(b *testing.B) {
	arr := []string{"1", "2", "3", "4", "5"}
	//开始计时,性能无关代码放在外面
	b.ResetTimer()
	//循环测试
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, value := range arr {
			ret += value
		}
	}
	//结束并计时
	b.StopTimer()
}

func BenchmarkConcatStringByBuffer(b *testing.B) {
	arr := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		for _, value := range arr {
			buf.WriteString(value)
		}
	}
	b.StopTimer()
}
