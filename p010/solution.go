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

// Primes : primes sequence generator
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

// SumPrimes : sum of all primes below a value
func SumPrimes(below int) int {
	p := Primes()
	result := 0
	for i := p(); i < below; i = p(){
		result += i
	}
	return result
}

func main() {
	result := SumPrimes(2000000)
	fmt.Printf("Answer: %v\n", result)
}
