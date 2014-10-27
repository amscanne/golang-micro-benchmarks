package benchmarks

import (
	"testing"
)

const ChannelBuffer = (64 * 1024)

func produce(x chan int, count int) {
	for count > 0 {
		x <- count
		count--
	}
}

func consume(x chan int, d chan struct{}, count int) {
	for count > 0 {
		<-x
		count--
	}
	var notice struct{}
	d <- notice
}

func BenchmarkChannel(b *testing.B) {
	x := make(chan int, ChannelBuffer)
	d := make(chan struct{})
	go produce(x, b.N)
	go consume(x, d, b.N)
	<-d
}

func BenchmarkMultiChannel(b *testing.B) {
	x := make(chan int, ChannelBuffer)
	d := make(chan struct{})
	for i := 0; i < 10; i++ {
		go produce(x, b.N/10)
	}
	for i := 0; i < 10; i++ {
		go consume(x, d, b.N/10)
	}
	for i := 0; i < 10; i++ {
		<-d
	}
}
