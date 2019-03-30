package main

import "fmt"

// SpecialPythagoreanTripletProduct : find the product abc, where a < b < c, a² + b² = c² and a + b + c = sum
func SpecialPythagoreanTripletProduct(sum int) int {
	for a := 1; 3*a < sum; a++ {
		for b := a + 1; a+2*b < sum; b++ {
			c := sum - a - b
			if a*a+b*b == c*c {
				return a * b * c
			}
		}
	}
	return 0
}

func main() {
	result := SpecialPythagoreanTripletProduct(1000)
	fmt.Printf("Answer: %v\n", result)
}
