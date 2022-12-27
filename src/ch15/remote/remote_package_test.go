package remote

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestRemoteConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("Key"), 10)
	t.Log(m.Get(cm.StrKey("Key")))
}
