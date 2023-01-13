package passphrase

import (
	"fmt"
	"idpassgen/pkg/words"
	"math"
	"math/rand"
	"strings"
	"time"
)

const (
	maximumDigits = 3
)

type Options struct {
	separator        string
	wordsCount       int
	isNumberIncluded bool
}

func defaultOptions() *Options {
	return &Options{
		separator:        "-",
		wordsCount:       3,
		isNumberIncluded: true,
	}
}

func NewOptions() *Options {
	return defaultOptions()
}

func (i *Options) WithSeparator(separator string) {
	i.separator = separator
}
func (i *Options) WithWordsCount(wordsCount int) {
	i.wordsCount = wordsCount
}
func (i *Options) WithNumber(isNumberIncluded bool) {
	i.isNumberIncluded = isNumberIncluded
}

func (i *Options) GeneratePassPhrase() (passPhrase string) {
	rand.Seed(time.Now().UnixNano())
	wordsList := words.RetrieveWordsFromOSDict()

	numberLocation := rand.Intn(i.wordsCount)

	for k := 0; k < i.wordsCount; k++ {
		passPhrase = passPhrase + strings.ToLower(wordsList[rand.Int63()%int64(len(wordsList))])
		if i.isNumberIncluded && k == numberLocation {
			max := int(math.Pow(float64(10), float64(maximumDigits)))
			passPhrase += fmt.Sprint(rand.Int() % max)
		}
		if k != i.wordsCount-1 {
			passPhrase = passPhrase + i.separator
		}
	}
	return
}
