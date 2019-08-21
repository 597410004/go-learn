package char11

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 15)
	t.Log(m.Get(cm.StrKey("key")))
}
