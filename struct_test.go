package benchmarks

import (
	"testing"
)

func do_struct(N int, value bool, output *int) {
	if value {
		mut := &Adder{}
		for i := 0; i < N; i++ {
			mut.Mutate(output)
		}
	} else {
		mut := &Subtractor{}
		for i := 0; i < N; i++ {
			mut.Mutate(output)
		}
	}
}

func BenchmarkStruct(b *testing.B) {
	var value int
	do_struct(b.N, true, &value)
	do_struct(b.N, false, &value)
}
