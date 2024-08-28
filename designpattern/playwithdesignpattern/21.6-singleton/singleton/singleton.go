package singleton

import (
	"sync"
	"sync/atomic"
)

type singleton struct{}

var instance *singleton
var mu sync.Mutex

var initialized uint32

func GetInstance() *singleton {

	if atomic.LoadUint32(&initialized) == 1 { // 原子操作
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}
