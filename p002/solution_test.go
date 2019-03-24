package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci10(t *testing.T) {
	// given
	const seqSize = 10
	f := Fibonacci()

	// when
	var result [seqSize]int
	for i := 0; i < seqSize; i++ {
		result[i] = f()
	}

	// then
	expected := [seqSize]int {1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	assert.Equal(t, expected, result)
}

func TestEvenFibonacciSum10(t *testing.T) {
	// given
	max := 10

	// when
	result := EvenFibonacciSum(max)

	// then
	assert.Equal(t, 10, result)
}
