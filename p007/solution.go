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

func PrimeNumber(pos int) int {
	p := Primes()
	for i := 1; i < pos; i++{
		p()
	}
	return p()
}

func main() {
	result := PrimeNumber(10001)
	fmt.Printf("Answer: %v\n", result)
}
