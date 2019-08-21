package part3

import "fmt"
// 结构体中有相同字段
type Humans struct {
	name   string
	age    int
	weight int
}
type Students struct {
	name string
	int
	Humans
}

func main() {
	s:=new(Students)
	s.name="dong"
	s.int=10
	s.Humans.name="han"
	fmt.Println(s.Humans.name)
	fmt.Println(s.name)
}