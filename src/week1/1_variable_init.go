package main

import (
	"fmt"
	"strconv"
)

// Group rated variables in same var ()
var (
	actorName string = "Actor 1"
	season    int    = 11
)

var (
	i int
)

func block_scope() float32 {
	var j float32 = 3.1415
	return j
}

// variables get shadowed in main program
// UPPER CASE VARIABLES ARE SCOPED GLOBALY
// LOWER CASE VARIABLES ARE SCOPED TO THE PACKAGE, ANYTHING THAT CONSUMES THE PACKAGE CANT SEE IT (ANY FILE IN THE SAME PACKAGE CAN SE IT)
// BLOCK SCOPE VARIABLES SCOPED IN THE BLOCK {} OF FUNCTIONS ARE NEVER VISIBLE TO OTHERS
// THIS FUNCTION CAN ACCESS j float32 variable unless called

func main() {

	i = 2
	fmt.Println(i)

	called_func := block_scope()

	fmt.Println(called_func)

	// Convertions operator
	fmt.Printf("Convertion Operator \n")
	var l int = 32

	fmt.Printf("We want to convert an interger type to a float type: %v, %T \n", l, l)

	var j float64
	j = float64(l)

	fmt.Printf("Interger type converted to a float type: %v, %T \n", j, j)

	// It is needed to make an explicit conversion while changing types

	// Convertions different types to strings

	var conv string
	conv = strconv.FormatFloat(j, 'E', -1, 64)
	fmt.Printf("Convertion Operator: %s \n", conv)	

	
}
