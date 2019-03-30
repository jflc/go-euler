package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpecialPythagoreanTripletProductOf12(t *testing.T) {
	// given
	sum := 12

	// when
	result := SpecialPythagoreanTripletProduct(sum)

	// then
	assert.Equal(t, 60, result)
}

func TestSpecialPythagoreanTripletProductOf1000(t *testing.T) {
	// given
	sum := 1000

	// when
	result := SpecialPythagoreanTripletProduct(sum)

	// then
	assert.Equal(t, 31875000, result)
}
