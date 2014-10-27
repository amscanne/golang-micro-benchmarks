package benchmarks

import (
	"testing"
)

func do_interface(N int, value bool, output *int) {
	var mut Mutator
	if value {
		mut = &Adder{}
	} else {
		mut = &Subtractor{}
	}
	for i := 0; i < N; i++ {
		mut.Mutate(output)
	}
}

func BenchmarkInterface(b *testing.B) {
	var value int
	do_interface(b.N, true, &value)
	do_interface(b.N, false, &value)
}
