package char10

import (
	"char9/series"
	"testing"
)
//测试包引用,
// 注意:
// 1.引用的包必须是正常go文件不可以是测试文件
// 2.必须将go文件添加到环境变量并重启idea
func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(9))
}
