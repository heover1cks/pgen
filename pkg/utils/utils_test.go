package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRandCombination(t *testing.T) {
	start := time.Now()
	for _, x := range RandCombination(10, 3) {
		print(x)
	}
	println()
	end := time.Since(start)
	println("elapsed: ", end.String())
}

func TestRandCombinationWithCombin(t *testing.T) {
	start := time.Now()
	for _, x := range RandCombinationWithCombin(10, 3) {
		print(x)
	}
	println()
	end := time.Since(start)
	println("elapsed: ", end.String())
}

func TestShift(t *testing.T) {
	testSlice1 := []int{1, 2, 3, 4, 5, 6}
	val, rest := Shift(testSlice1)
	assert.Equal(t, 1, val)
	assert.Equal(t, []int{2, 3, 4, 5, 6}, rest)
	println(val)
	PrintArrayValues(rest)

	testSlice2 := []int{7, 8}
	val, rest = Shift(testSlice2)
	assert.Equal(t, 7, val)
	assert.Equal(t, []int{8}, rest)
	println(val)
	PrintArrayValues(rest)

	val, rest = Shift(rest)
	assert.Equal(t, 8, val)
	assert.Equal(t, []int{}, rest)
	println(val)
	PrintArrayValues(rest)
}
