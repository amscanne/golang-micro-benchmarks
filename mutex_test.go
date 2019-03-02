package benchmarks

import (
	"sync"
	"testing"
)

func BenchmarkMutexLock(b *testing.B) {
	var m sync.Mutex
	for i := 0; i < b.N; i++ {
		m.Lock()
		m.Unlock()
	}
}

func BenchmarkRWMutexLock(b *testing.B) {
	var m sync.RWMutex
	for i := 0; i < b.N; i++ {
		m.Lock()
		m.Unlock()
	}
}

func BenchmarkRWMutexRLock(b *testing.B) {
	var m sync.RWMutex
	for i := 0; i < b.N; i++ {
		m.RLock()
		m.RUnlock()
	}
}
