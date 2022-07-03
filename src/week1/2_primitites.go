package main

import (
	"fmt"
)

func main() {

	var n bool = true
	fmt.Printf("%v, %T \n", n, n)

	//Boolean Init var with bool
	m := 1 == 1
	l := 1 == 2

	fmt.Printf("%v, %T \n", m, m)
	fmt.Printf("%v, %T \n", l, l)

	// unigned Integers
	// uint8 0-255 Vs. signed integers (int) -128-127

	var p uint8 = 2
	fmt.Printf("%v, %T \n", p, p)

	//bit Operations

	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  //0010 => 2
	fmt.Println(a | b)  // 1011 => 11 OR
	fmt.Println(a ^ b)  // 1001 => 9 (Exclusive OR)
	fmt.Println(a &^ b) // 0100 -> AND NOT, OPPOSITE TO OR

	//bit shifting

	g := 9              // 1001
	fmt.Println(g << 3) // 1001000 -> 72
	fmt.Println(g >> 3) // 1

	//complex numbers
	// float32 +  float32i
	// 1+2i = complex(1, 2)
	var x complex64 = 1 + 2i
	fmt.Printf("%v, %T \n", x, x)

	//real part
	fmt.Printf("Real part %v, %T \n", real(x), real(x))
	fmt.Printf("Img part %v, %T \n", imag(x), imag(x))

	//String as array, strings are unmutable
	s := "This is a string"
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	//slice of bytes

	y := []byte(s)
	fmt.Printf("%v, %T\n", y, y)

	// Rune '' (single quotes) u32 for strings

	testString := '♛'
	fmt.Printf("Rune 1: %c; Unicode: %U; ", testString, testString)
}
