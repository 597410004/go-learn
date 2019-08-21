package char17

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFuctions(t*testing.T)  {
	fmt.Println(square(15))
	assert.Equal(t,225,square(15),"should equal")

}

func TestFailtInCode(t*testing.T){
	t.Log("start")
	t.Error("error")
	t.Log("end1")
	t.FailNow()
	t.Log("end2")
}
