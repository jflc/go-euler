package main

import (
	"fmt"
	"math"
)

func reverse(num int) int {
	result := 0

	for num != 0 {
		result *= 10
		result += num % 10
		num /= 10
	}

	return result
}

func isPalindrome(num int) bool {
	return num == reverse(num)
}

func LargestPalindromeProduct(digits int) int{
	result := 0

	min := math.Pow10(digits-1)
	max := math.Pow10(digits)-1

	for i := min; i < max; i++ {
		for j := max; j >= i; j-- {
			product := int(i*j)
			if product < result {
				break
			}
			if isPalindrome(product) {
				result = product
				break
			}
		}
	}

	return result
}


func main() {
	result := LargestPalindromeProduct(3)
	fmt.Printf("Answer: %v\n", result)
}
