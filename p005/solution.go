package main

import (
	"fmt"
)


func isMultiple(num, min, max int) bool {
	for i := min; i <= max; i++ {
		if num % i != 0 {
			return false
		}
	}

	return true
}

func SmallestMultiple(min, max int) int {
	result := 0

	for i := max; result == 0; i += max {
		if isMultiple(i, min, max) {
			result = i
		}
	}

	return result
}



func main() {
	result := SmallestMultiple(1, 20)
	fmt.Printf("Answer: %v\n", result)
}
