package benchmarks

type Mutator interface {
    Mutate(output *int)
}

type Adder struct {
}

type Subtractor struct {
}

func add(output *int) {
    *output += 1
}

func (add *Adder) Mutate(output *int) {
    *output += 1
}

func subtract(output *int) {
    *output -= 1
}

func (sub *Subtractor) Mutate(output *int) {
    *output -= 1
}
