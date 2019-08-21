package part1

import (
	"fmt"
	"testing"
)

type Student struct {
	Id   string
	Name string
}
type Human struct {
	Name      string
	Funcation func(int) int
}

func add(i int) int {
	return i + 1
}

// 对象定义及方法使用
func TestCreateStudent(t *testing.T) {
	list := Human{Name: "hdd",Funcation:add}
	t.Logf("human is %v", list.Funcation(1))
	s := Student{"1", "hdd"}
	s.Id = "5"
	t.Log(s)
	s1 := Student{Id: "1", Name: "hdd"}
	t.Log(s1)
	s2 := new(Student)
	s2.Name = "hdd"
	s2.Id = "1"
	t.Log(s2.String())
}

//func(s Student)String()string{
//	return fmt.Sprintf("id:%v-name:%v",s.Id,s.Name)
//}
func (s *Student) String() string {
	return fmt.Sprintf("name is %v - value is %v", s.Name, s.Id)
}
