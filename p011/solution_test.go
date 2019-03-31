package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntGridSuccess(t *testing.T) {
	// given
	grid := String(`1 2 3
	4 5 6
	7 8 9`)

	// when
	result, err := grid.IntGrid()

	// then
	expected := IntGrid{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestIntGridError(t *testing.T) {
	// given
	grid := String(`a b c
	d e f
	g h i`)

	// when
	_, err := grid.IntGrid()

	// then
	assert.Error(t, err)
}


func TestProductDownSuccess(t *testing.T) {
	// given
	grid := IntGrid([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	row := 0
	column := 0
	adjacentDigits := 3
	direction := Down

	// when
	result, err := grid.Product(row, column, adjacentDigits, direction)

	// then
	assert.NoError(t, err)
	assert.Equal(t, 28, result)
}

func TestProductDownInvalidAdjacentDigits(t *testing.T) {
	// given
	grid := IntGrid([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	row := 0
	column := 0
	adjacentDigits := 4
	direction := Down

	// when
	_, err := grid.Product(row, column, adjacentDigits, direction)

	// then
	assert.Error(t, err)
}

func TestProductRightSuccess(t *testing.T) {
	// given
	grid := IntGrid([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	row := 0
	column := 0
	adjacentDigits := 3
	direction := Right

	// when
	result, err := grid.Product(row, column, adjacentDigits, direction)

	// then
	assert.NoError(t, err)
	assert.Equal(t, 6, result)
}

func TestProductRightInvalidAdjacentDigits(t *testing.T) {
	// given
	grid := IntGrid([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	row := 0
	column := 0
	adjacentDigits := 4
	direction := Right

	// when
	_, err := grid.Product(row, column, adjacentDigits, direction)

	// then
	assert.Error(t, err)
}

func TestProductUpLeftSuccess(t *testing.T) {
	// given
	grid := IntGrid([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	row := 2
	column := 2
	adjacentDigits := 3
	directions := []Direction{Up, Left}

	// when
	result, err := grid.Product(row, column, adjacentDigits, directions...)

	// then
	assert.NoError(t, err)
	assert.Equal(t, 45, result)
}

func TestProductUpLeftInvalidAdjacentDigits(t *testing.T) {
	// given
	grid := IntGrid([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	row := 1
	column := 1
	adjacentDigits := 3
	directions := []Direction{Up, Left}

	// when
	_, err := grid.Product(row, column, adjacentDigits, directions...)

	// then
	assert.Error(t, err)
}

func TestProductLeftInvalidIndexes(t *testing.T) {
	// given
	grid := IntGrid([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	row := 3
	column := 3
	adjacentDigits := 1
	directions := []Direction{Left}

	// when
	_, err := grid.Product(row, column, adjacentDigits, directions...)

	// then
	assert.Error(t, err)
}

func TestLargestProductSmallGrid(t *testing.T) {
	// given
	grid := IntGrid([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	adjacentDigits := 2

	// when
	result := LargestProduct(grid, adjacentDigits)

	// then
	assert.Equal(t, 72, result)
}

func TestLargestProductDefaultGrid(t *testing.T) {
	// given
	grid := defaultGrid
	adjacentDigits := 4

	// when
	result := LargestProduct(grid, adjacentDigits)

	// then
	assert.Equal(t, 70600674, result)
}