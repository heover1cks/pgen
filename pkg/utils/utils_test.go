package utils

import (
	"testing"
	"time"
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
