package char16

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type ReusableObj struct {
}
type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(num int) *ObjPool {
	pool := new(ObjPool)
	pool.bufChan = make(chan *ReusableObj, num)
	for i := 0; i < num; i++ {
		pool.bufChan <- &ReusableObj{}
	}
	return pool
}
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")

	}
}
func (p *ObjPool) RealseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("over flow")


	}
}
func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 12 ; i++ {
		if v, err := pool.GetObj(1); err != nil {
			t.Log(err)
		} else {
			fmt.Printf("%T\n",v)
			//if err := pool.RealseObj(v); err != nil {
			//	t.Error(err)
			//}
		}
	}
	t.Log("Done")
}
