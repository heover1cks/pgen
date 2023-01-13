package utils

import "math/rand"

func BoolCounter(inputs ...bool) (ret int) {
	for _, input := range inputs {
		if input == true {
			ret += 1
		}
	}
	return
}

func RandStringBytes(n int, characterSet []byte) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = characterSet[rand.Int63()%int64(len(characterSet))]
	}
	return b
}

func Shift[T any](s []T) (T, []T) {
	return s[0], s[1:]
}

func Pop[T any](s []T) (T, []T) {
	return s[len(s)-1], s[:len(s)-1]
}

func PrintArrayValues[T any](vals []T) {
	for _, v := range vals {
		print(v, " ")
	}
	println()
}
