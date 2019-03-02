package benchmarks

import (
	"testing"
)

func produce(x chan struct{}, count int) {
	for i := 0; i < count; i++ {
		x <- struct{}{}
	}
}

func produceNonBlocking(x chan struct{}, count int) {
	for i := 0; i < count; i++ {
		select {
		case x <- struct{}{}:
		default:
		}
	}
}

func consume(x chan struct{}, count int) {
	for i := 0; i < count; i++ {
		<-x
	}
}

func consumeNonBlocking(x chan struct{}, count int) {
	for i := 0; i < count; i++ {
		select {
		case <-x:
		default:
		}
	}
}

func BenchmarkChannelReceive(b *testing.B) {
	b.StopTimer()
	d := make(chan struct{}, b.N)
	produce(d, b.N)
	b.StartTimer()
	consume(d, b.N)
}

func BenchmarkChannelSend(b *testing.B) {
	b.StopTimer()
	d := make(chan struct{}, b.N)
	b.StartTimer()
	produce(d, b.N)
}

func BenchmarkChannelNonBlockingReceive(b *testing.B) {
	b.StopTimer()
	d := make(chan struct{}, b.N)
	produce(d, b.N)
	b.StartTimer()
	consumeNonBlocking(d, b.N)
}

func BenchmarkChannelNonBlockingReceiveEmpty(b *testing.B) {
	b.StopTimer()
	d := make(chan struct{}, b.N)
	b.StartTimer()
	consumeNonBlocking(d, b.N)
}

func BenchmarkChannelNonBlockingSend(b *testing.B) {
	b.StopTimer()
	d := make(chan struct{}, b.N)
	b.StartTimer()
	produceNonBlocking(d, b.N)
}

func BenchmarkChannelNonBlockingSendFull(b *testing.B) {
	b.StopTimer()
	d := make(chan struct{}, b.N)
	produce(d, b.N)
	b.StartTimer()
	produceNonBlocking(d, b.N)
}
