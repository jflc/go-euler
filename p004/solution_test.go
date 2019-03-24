package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindromeTrue(t *testing.T) {
	// given
	num := 9009

	// when
	result := isPalindrome(num)

	// then
	assert.True(t, result)
}

func TestIsPalindromeFalse(t *testing.T) {
	// given
	num := 9090

	// when
	result := isPalindrome(num)

	// then
	assert.False(t, result)
}

func TestLargestPalindromeProduct2(t *testing.T) {
	// given
	digits := 2

	// when
	result := LargestPalindromeProduct(digits)

	// then
	assert.Equal(t, 9009, result)
}

func TestLargestPalindromeProduct3(t *testing.T) {
	// given
	digits := 3

	// when
	result := LargestPalindromeProduct(digits)

	// then
	assert.Equal(t, 906609, result)
}