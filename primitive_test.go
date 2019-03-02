package benchmarks

import (
	"testing"
)

type Opaque struct {
	val int
}

func (o Opaque) Add(other Opaque) Opaque {
	return Opaque{o.val + other.val}
}

func (o *Opaque) AddPtr(other *Opaque) Opaque {
	return Opaque{o.val + other.val}
}

func BenchmarkStructOpByValue(b *testing.B) {
	c := Opaque{0}
	inc := Opaque{1}
	for i := 0; i < b.N; i++ {
		c = c.Add(inc)
	}
}

func BenchmarkStructOpByPtr(b *testing.B) {
	c := Opaque{0}
	inc := &Opaque{1}
	for i := 0; i < b.N; i++ {
		c = c.AddPtr(inc)
	}
}

func BenchmarkPrimtiveOp(b *testing.B) {
	c := 0
	inc := 1
	for i := 0; i < b.N; i++ {
		c = c + inc
	}
}
