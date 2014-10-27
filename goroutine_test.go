package benchmarks

import (
	"runtime"
	"sync/atomic"
	"testing"
)

func do_inc(n chan struct{}) {
	n <- struct{}{}
}

func BenchmarkGoroutineCreate(b *testing.B) {
	n := make(chan struct{})
	for i := 0; i < b.N; i++ {
		go do_inc(n)
		<-n
	}
}

func do_wait(val *int64, c chan interface{}) {
	x := int(0)
	atomic.AddInt64(val, 1)
	c <- &x
}

func BenchmarkGoroutinePersist(b *testing.B) {
	var val int64
	c := make(chan interface{})
	for i := 0; i < b.N; i++ {
		go do_wait(&val, c)
	}
	for atomic.LoadInt64(&val) != int64(b.N) {
		runtime.Gosched()
	}
}
