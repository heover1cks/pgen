package utils

import (
	cryptoRand "crypto/rand"
	"gonum.org/v1/gonum/stat/combin"
	"math"
	"math/big"
	"math/rand"
)

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

func RandCombination(n, k int) (ret []int) {
	seed, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	cur, prev := 0, 0
	for x := 0; x < k; x++ {
		cur = rand.Int()%(n-prev-k+x+1) + 1
		println(n-prev-k+x, cur, prev)
		ret = append(ret, cur+prev)
		prev = cur + prev
	}
	return
}

func RandCombinationWithCombin(n, k int) []int {
	seed, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	end := combin.Binomial(n, k)
	idx := rand.Int() % end
	cg := combin.NewCombinationGenerator(n, k)
	for i := 0; i < idx; i++ {
		cg.Next()
	}
	return cg.Combination(nil)
}
