package benchmarks

import (
	"testing"
)

func varargs(args ...uintptr) {
	args[0] = args[1] + args[2]
}

func slice(args []uintptr) {
	args[0] = args[1] + args[2]
}

func BenchmarkArgsVar(b *testing.B) {
	args := make([]uintptr, 3)
	for i := 0; i < b.N; i++ {
		varargs(args...)
	}
}

func BenchmarkArgsSlice(b *testing.B) {
	args := make([]uintptr, 3)
	for i := 0; i < b.N; i++ {
		slice(args)
	}
}
