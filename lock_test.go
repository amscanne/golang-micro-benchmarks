package benchmarks

import (
	"sync"
	"testing"
)

func lock(mutex *sync.Mutex, count int, d chan struct{}) {
	for count > 0 {
		mutex.Lock()
		count--
		mutex.Unlock()
	}
	var notice struct{}
	d <- notice
}

func BenchmarkLock(b *testing.B) {
	m := &sync.Mutex{}
	d := make(chan struct{})
	go lock(m, b.N, d)
	<-d
}

func BenchmarkMultiLock(b *testing.B) {
	m := &sync.Mutex{}
	d := make(chan struct{})
	for i := 0; i < 10; i++ {
		go lock(m, b.N/10, d)
	}
	for i := 0; i < 10; i++ {
		<-d
	}
}
