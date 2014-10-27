package benchmarks

import (
	"testing"
)

func do_switch(N int, value int, output *int) {
	for i := 0; i < N; i++ {
		switch value {
		case 0:
			*output += 1
		case 1:
			*output += 2
		case 2:
			*output += 3
		case 3:
			*output += 4
		case 4:
			*output += 5
		case 5:
			*output += 6
		case 6:
			*output += 7
		case 7:
			*output += 8
		case 8:
			*output += 9
		case 9:
			*output += 10
		case 10:
			*output += 11
		case 11:
			*output += 12
		case 12:
			*output += 13
		case 13:
			*output += 14
		case 14:
			*output += 15
		case 15:
			*output += 16
		case 16:
			*output += 17
		case 17:
			*output += 18
		case 18:
			*output += 19
		case 19:
			*output += 20
		}
	}
}

func do_if_chain(N int, value int, output *int) {
	for i := 0; i < N; i++ {
		if value == 0 {
			*output += 1
		} else if value == 1 {
			*output += 2
		} else if value == 2 {
			*output += 3
		} else if value == 3 {
			*output += 4
		} else if value == 4 {
			*output += 5
		} else if value == 5 {
			*output += 6
		} else if value == 6 {
			*output += 7
		} else if value == 7 {
			*output += 8
		} else if value == 8 {
			*output += 9
		} else if value == 9 {
			*output += 10
		} else if value == 10 {
			*output += 11
		} else if value == 11 {
			*output += 12
		} else if value == 12 {
			*output += 13
		} else if value == 13 {
			*output += 14
		} else if value == 14 {
			*output += 15
		} else if value == 15 {
			*output += 16
		} else if value == 16 {
			*output += 17
		} else if value == 17 {
			*output += 18
		} else if value == 18 {
			*output += 19
		} else if value == 19 {
			*output += 20
		}
	}
}

var jumpTable = map[int]func(output *int){
	0:  func(output *int) { *output += 1 },
	1:  func(output *int) { *output += 2 },
	2:  func(output *int) { *output += 3 },
	3:  func(output *int) { *output += 4 },
	4:  func(output *int) { *output += 5 },
	5:  func(output *int) { *output += 6 },
	6:  func(output *int) { *output += 7 },
	7:  func(output *int) { *output += 8 },
	8:  func(output *int) { *output += 9 },
	9:  func(output *int) { *output += 10 },
	10: func(output *int) { *output += 11 },
	11: func(output *int) { *output += 12 },
	12: func(output *int) { *output += 13 },
	13: func(output *int) { *output += 14 },
	14: func(output *int) { *output += 15 },
	15: func(output *int) { *output += 16 },
	16: func(output *int) { *output += 17 },
	17: func(output *int) { *output += 18 },
	18: func(output *int) { *output += 19 },
	19: func(output *int) { *output += 20 },
}

func do_jump_table(N int, value int, output *int) {
	for i := 0; i < N; i++ {
		fn, ok := jumpTable[value]
		if ok {
			fn(output)
		}
	}
}

func BenchmarkSwitch(b *testing.B) {
	var value int
	do_switch(b.N, 0, &value)
	do_switch(b.N, 1, &value)
	do_switch(b.N, 2, &value)
	do_switch(b.N, 3, &value)
	do_switch(b.N, 4, &value)
	do_switch(b.N, 5, &value)
	do_switch(b.N, 6, &value)
	do_switch(b.N, 7, &value)
	do_switch(b.N, 8, &value)
	do_switch(b.N, 9, &value)
	do_switch(b.N, 10, &value)
	do_switch(b.N, 11, &value)
	do_switch(b.N, 12, &value)
	do_switch(b.N, 13, &value)
	do_switch(b.N, 14, &value)
	do_switch(b.N, 15, &value)
	do_switch(b.N, 16, &value)
	do_switch(b.N, 17, &value)
	do_switch(b.N, 18, &value)
	do_switch(b.N, 19, &value)
}

func BenchmarkIfChain(b *testing.B) {
	var value int
	do_if_chain(b.N, 0, &value)
	do_if_chain(b.N, 1, &value)
	do_if_chain(b.N, 2, &value)
	do_if_chain(b.N, 3, &value)
	do_if_chain(b.N, 4, &value)
	do_if_chain(b.N, 5, &value)
	do_if_chain(b.N, 6, &value)
	do_if_chain(b.N, 7, &value)
	do_if_chain(b.N, 8, &value)
	do_if_chain(b.N, 9, &value)
	do_if_chain(b.N, 10, &value)
	do_if_chain(b.N, 11, &value)
	do_if_chain(b.N, 12, &value)
	do_if_chain(b.N, 13, &value)
	do_if_chain(b.N, 14, &value)
	do_if_chain(b.N, 15, &value)
	do_if_chain(b.N, 16, &value)
	do_if_chain(b.N, 17, &value)
	do_if_chain(b.N, 18, &value)
	do_if_chain(b.N, 19, &value)
}

func BenchmarkJumpTable(b *testing.B) {
	var value int
	do_jump_table(b.N, 0, &value)
	do_jump_table(b.N, 1, &value)
	do_jump_table(b.N, 2, &value)
	do_jump_table(b.N, 3, &value)
	do_jump_table(b.N, 4, &value)
	do_jump_table(b.N, 5, &value)
	do_jump_table(b.N, 6, &value)
	do_jump_table(b.N, 7, &value)
	do_jump_table(b.N, 8, &value)
	do_jump_table(b.N, 9, &value)
	do_jump_table(b.N, 10, &value)
	do_jump_table(b.N, 11, &value)
	do_jump_table(b.N, 12, &value)
	do_jump_table(b.N, 13, &value)
	do_jump_table(b.N, 14, &value)
	do_jump_table(b.N, 15, &value)
	do_jump_table(b.N, 16, &value)
	do_jump_table(b.N, 17, &value)
	do_jump_table(b.N, 18, &value)
	do_jump_table(b.N, 19, &value)
}
