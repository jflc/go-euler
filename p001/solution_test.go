package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumMultiples1(t *testing.T) {
	// given
	max := 10
	nums := []int{3,5}

	// when
	result := SumMultiples(max, nums...)

	// then
	assert.Equal(t, 23, result)
}

func TestSumMultiples2(t *testing.T) {
	// given
	max := 10
	nums := []int{1,2,3}

	// when
	result := SumMultiples(max, nums...)

	// then
	assert.Equal(t, 45, result)
}

func TestSumMultiples3(t *testing.T) {
	// given
	max := 1000
	nums := []int{3,5}

	// when
	result := SumMultiples(max, nums...)

	// then
	assert.Equal(t, 233168, result)
}
