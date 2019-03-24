package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestProductWith4(t *testing.T) {
	// given
	adjacentDigits := 4

	// when
	result := LargestProduct(adjacentDigits)

	// then
	assert.Equal(t, 5832, result)
}

func TestLargestProductWith13(t *testing.T) {
	// given
	adjacentDigits := 13

	// when
	result := LargestProduct(adjacentDigits)

	// then
	assert.Equal(t, 23514624000, result)
}