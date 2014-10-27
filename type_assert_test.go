package benchmarks

import (
	"testing"
)

func do_type_assert(N int, value interface{}, output *int) {
	for i := 0; i < N; i++ {
		v := value.(int)
		*output += v
	}
}

func BenchmarkTypeAssert(b *testing.B) {
	var value int
	do_type_assert(b.N, int(1), &value)
	do_type_assert(b.N, int(-1), &value)
}
