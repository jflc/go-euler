package main

import "fmt"

func Fibonacci() func() int {
	current := 1
	next := 2

	return func() int{
		defer func() {
			current, next = next, current + next
		}()
		return current
	}
}

func EvenFibonacciSum(max int) int {
	f := Fibonacci()
	result := 0

	for num := f(); num < max; num = f() {
		if num % 2 == 0 {
			result += num
		}
	}

	return result
}

func main() {
	result := EvenFibonacciSum(4000000)
	fmt.Printf("Answer: %v\n", result)
}
