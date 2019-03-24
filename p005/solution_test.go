package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestIsMultipleTrue(t *testing.T) {
	// given
	num := 2520
	min := 1
	max := 10

	// when
	result := isMultiple(num, min, max)

	// then
	assert.True(t, result)
}

func TestIsMultipleFalse(t *testing.T) {
	// given
	num := 12345
	min := 2
	max := 8

	// when
	result := isMultiple(num, min, max)

	// then
	assert.False(t, result)
}

func TestSmallestMultipleFrom1To10(t *testing.T) {
	// given
	min := 1
	max := 10

	// when
	result := SmallestMultiple(min, max)

	// then
	assert.Equal(t,2520, result)
}

func TestSmallestMultipleFrom1To20(t *testing.T) {
	// given
	min := 1
	max := 20

	// when
	result := SmallestMultiple(min, max)

	// then
	assert.Equal(t,232792560, result)
}
