package benchmarks

import (
	"testing"
)

func do_closure(N int, value bool, output *int) {
	var fn func()
	if value {
		fn = func() { *output += 1 }
	} else {
		fn = func() { *output -= 1 }
	}
	for i := 0; i < N; i++ {
		fn()
	}
}

func BenchmarkClosure(b *testing.B) {
	var value int
	do_closure(b.N, true, &value)
	do_closure(b.N, false, &value)
}
