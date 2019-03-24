package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	maxDivisor := int(math.Sqrt(float64(num)))
	for i := 2; i <= maxDivisor;  i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func Primes() func() int {
	current := 1

	return func() int{
		current++
		for !isPrime(current) {
			current++
		}
		return current
	}
}

func LargestPrimeFactor(num int) int {
	p := Primes()
	result := 0

	maxDivisor := int(math.Sqrt(float64(num)))
	for i := p(); i < maxDivisor; i = p() {
		if num % i == 0 {
			result = i
		}
	}

	return result
}


func main() {
	result := LargestPrimeFactor(600851475143)
	fmt.Printf("Answer: %v\n", result)
}
