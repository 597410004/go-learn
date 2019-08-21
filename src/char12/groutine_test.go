package char12

import (
	"sync"
	"testing"
	"time"
)

// 协诚
func TestGroutine(t *testing.T) {
	for i := 1; i < 10; i++ {
		go func(i int) {
			t.Log(i)
		}(i)
		time.Sleep(50 * time.Millisecond)
	}
}

// 并发机制,不安全
func TestCount(t *testing.T) {
	count := 0
	for i := 0; i < 500; i++ {
		go func() {
			count++
		}()
	}
	t.Logf("count is :%v", count)
}

//加锁操作(设置固定时间等待)
func TestSafeCount(t *testing.T) {
	// 同步锁
	var mul sync.Mutex
	count := 0
	for i := 0; i < 900; i++ {
		go func() {
			defer func() {
				// 最后释放锁
				mul.Unlock()
			}()
			// 加锁
			mul.Lock()
			count++
		}()
	}
	//如果去掉等待时间则不会显示正常结果
	time.Sleep(1 * time.Second)
	t.Logf("count is :%v", count)
}

// waitGroup 所有任务完成才执行下面
func TestSafeCountWait(t *testing.T) {
	// 同步锁
	var mul sync.Mutex
	var group sync.WaitGroup
	count := 0
	for i := 0; i < 900; i++ {
		group.Add(1)
		go func() {
			defer func() {
				// 最后释放锁
				mul.Unlock()
			}()
			// 加锁
			mul.Lock()
			group.Done()
			count++
		}()
	}
	group.Wait()
	t.Logf("count is :%v", count)
}
