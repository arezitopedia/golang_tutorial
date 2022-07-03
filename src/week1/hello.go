package main

import (
	"fmt"
	"os"
)

func main() {

	var responseVar string
	
	responseVar = fmt.Sprintf("Hello User %s", os.Getenv("USER"))

	fmt.Println(responseVar)
	fmt.Println("Hello, world!")
}
