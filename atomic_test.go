package benchmarks

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkAtomicSwap(b *testing.B) {
	var v uint32
	for i := 0; i < b.N; i++ {
		atomic.SwapUint32(&v, uint32(i))
	}
}

func BenchmarkAtomicCompareAndSwapSuccess(b *testing.B) {
	var v uint32
	for i := 1; i < b.N+1; i++ {
		atomic.CompareAndSwapUint32(&v, uint32(i-1), uint32(i))
	}
}

func BenchmarkAtomicCompareAndSwapFail(b *testing.B) {
	var v uint32
	for i := 0; i < b.N; i++ {
		atomic.CompareAndSwapUint32(&v, 1, uint32(i))
	}
}

func BenchmarkAtomicCompareAndSwapFailContended(b *testing.B) {
	var v uint32
	var wg sync.WaitGroup
	n := runtime.GOMAXPROCS(0)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < b.N/n; i++ {
				atomic.CompareAndSwapUint32(&v, 1, uint32(i))
			}
		}()
	}
	wg.Wait()
}

func BenchmarkAtomicLoad(b *testing.B) {
	var v uint32
	for i := 0; i < b.N; i++ {
		_ = atomic.LoadUint32(&v)
	}
}

func BenchmarkAtomicLoadContended(b *testing.B) {
	var v uint32
	var wg sync.WaitGroup
	n := runtime.GOMAXPROCS(0)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < b.N/n; i++ {
				_ = atomic.LoadUint32(&v)
			}
		}()
	}
	wg.Wait()
}
