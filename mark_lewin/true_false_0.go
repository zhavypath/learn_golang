package main

import (
	"fmt"
)

func main() {
	var trueValue bool
	trueValue = true
	falseValue := false

	fmt.Println(trueValue && trueValue)
	fmt.Println(trueValue && falseValue)
	fmt.Println(trueValue || trueValue)
	fmt.Println(trueValue || falseValue)
	fmt.Println(!trueValue)
}
