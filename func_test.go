package benchmarks

import (
	"testing"
)

func do_func(N int, value bool, output *int) {
	if value {
		for i := 0; i < N; i++ {
			add(output)
		}
	} else {
		for i := 0; i < N; i++ {
			subtract(output)
		}
	}
}

func BenchmarkFunc(b *testing.B) {
	var value int
	do_struct(b.N, true, &value)
	do_struct(b.N, false, &value)
}
