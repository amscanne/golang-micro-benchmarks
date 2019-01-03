package benchmarks

import (
	"testing"
)

func simple_func(output *int, value bool) {
	*output++
	if value { // Never run.
		simple_func(output, value)
	}
}

//go:nosplit
func nosplit_func(output *int, value bool) {
	*output++
	if value { // Never run.
		simple_func(output, value)
	}
}

func baseline(output *int, value bool) {
	*output++
	if value {
		// Never run. We want the cost of the comparison, but don't
		// make this recursive because we still want it to be inlined.
		*output++
	}
}

func BenchmarkBaselineFunc(b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		baseline(&value, false)
	}
}

func BenchmarkNormalFunc(b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		simple_func(&value, false)
	}
}

func BenchmarkNosplitFunc(b *testing.B) {
	var value int
	for i := 0; i < b.N; i++ {
		nosplit_func(&value, false)
	}
}
