package benchmarks

import (
    "testing"
)

func do_if(N int, value bool, output *int) {
    for i := 0; i < N; i++ {
        if value {
            *output += 1
        } else {
            *output -= 1
        }
    }
}

func BenchmarkIf(b *testing.B) {
    var value int
    do_if(b.N, true, &value)
    do_if(b.N, false, &value)
}
