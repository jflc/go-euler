package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumSquareDifferenceOf10(t *testing.T) {
	// given
	num := 10

	// when
	result := SumSquareDifference(num)

	// then
	assert.Equal(t, 2640, result)
}

func TestSumSquareDifferenceOf100(t *testing.T) {
	// given
	num := 100

	// when
	result := SumSquareDifference(num)

	// then
	assert.Equal(t, 25164150, result)
}