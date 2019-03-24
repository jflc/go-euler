package main

import (
	"fmt"
)

func SumSquareDifference(num int) int {
	sumOfSquares := (2*num+1)*(num+1)*num/6
	squareOfSum := num*(num+1)/2
	squareOfSum *= squareOfSum

	return squareOfSum - sumOfSquares
}

func main() {
	result := SumSquareDifference(100)
	fmt.Printf("Answer: %v\n", result)
}
