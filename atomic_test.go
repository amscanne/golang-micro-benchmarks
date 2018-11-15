package benchmarks

import (
	"sync/atomic"
	"testing"
)

func BenchmarkAtomicSwap(b *testing.B) {
	var v uint32
	for i := 0; i < b.N; i++ {
		atomic.SwapUint32(&v, uint32(i))
	}
}

func BenchmarkAtomicCompareAndSwap(b *testing.B) {
	var v uint32
	for i := 0; i < b.N; i++ {
		atomic.CompareAndSwapUint32(&v, 0, uint32(i))
	}
}
