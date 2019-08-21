package char15

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func runTask(id int) string {
	return fmt.Sprintf("The result is from %d", id)
}

func FistResponce() string {
	numOfRunner := 10
	var group sync.WaitGroup
	ch := make(chan string,numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		group.Add(1)
		go func(i int) {
			defer func() {
				group.Done()
			}()
			ret := runTask(i)
			ch <- ret
			fmt.Println(i)
		}(i)
	}
	group.Wait()
	return <-ch
}
func TestResponse(t *testing.T) {
	fmt.Printf("befroe :%v\n",runtime.NumGoroutine())
	id:=FistResponce()
	fmt.Printf("%s\n",id)
	fmt.Printf("after :%v\n",runtime.NumGoroutine())
}
