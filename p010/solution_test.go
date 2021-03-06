package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrimesSequence(t *testing.T) {
	// given
	const seqSize = 20
	p := Primes()

	// when
	var result [seqSize]int
	for i := 0; i < seqSize; i++ {
		result[i] = p()
	}

	// then
	expected := [seqSize]int {2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71}
	assert.Equal(t, expected, result)
}

func TestSumPrimesBelow10(t *testing.T) {
	// given
	below := 10

	// when
	result := SumPrimes(below)

	// then
	assert.Equal(t, 17, result)
}