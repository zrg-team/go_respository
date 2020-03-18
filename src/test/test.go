package main

import (
	"fmt"
)

// DefaultCriteria for chain function
func DefaultCriteria(input string, inputFunc func(string) string) interface{} {
	input = "hello"
	fmt.Printf("DefaultCriteria %s\n", input)
	return inputFunc(input)
}

func log(x string) string {
	fmt.Printf(">>>>>>>> %s\n", x)
	return x
}

func main() {
	DefaultCriteria("1", log)
	// api.Run()
}
