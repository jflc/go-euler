package main

import "fmt"

func SumMultiples(max int, nums... int) int {
	items := make(map[int]bool)
	result := 0
	for _, n := range nums {
		for i := n; i < max; i+= n {
			if _, exists := items[i]; !exists {
				result += i
				items[i] = true
			}

		}
	}
	return result
}

func main() {
	result := SumMultiples(1000, 3, 5)
	fmt.Printf("Answer: %v\n", result)
}
