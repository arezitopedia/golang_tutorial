package main

import (
	"fmt"
)

const myCost int64 = 32

const (
	a = iota //enumerated constant
	b = iota //enumerated constant
)

const (
	c = iota
	d = iota
)

func main() {
	const myCost int = 12
	fmt.Printf("%v,%T", myCost, myCost)

	// Constant must be defnied before COMPILE TIME, it cannot be defined by code executed at RUN time
	// and can be SHADOWED
	fmt.Printf("%v\n", myCost)
	// The compiler can infere the type of a constant per operation

	// iota consntant, can change values as asigned auto incremented by +1 stating from 0. And are specific to code blocks

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)

	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
}
