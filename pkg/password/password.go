package password

import (
	"gonum.org/v1/gonum/stat/combin"
	"idpassgen/pkg/utils"
	"log"
	"math/rand"
	"time"
)

const (
	alphabetUpperCaseBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetLowerCaseBytes = "abcdefghijklmnopqrstuvwxyz"
	numberBytes            = "1234567890"
	specialCharacterBytes  = "!@#$%^&*"
)

type Options struct {
	includeUpperCase         bool
	includeLowerCase         bool
	includeNumbers           bool
	includeSpecialCharacters bool

	maxNumbers           int
	maxSpecialCharacters int

	passLength int
}

func NewOptions() *Options {
	return &Options{
		includeUpperCase:         true,
		includeLowerCase:         true,
		includeNumbers:           true,
		includeSpecialCharacters: true,
		passLength:               16,
		maxNumbers:               3,
		maxSpecialCharacters:     3,
	}
}

func (p *Options) WithPassLength(length int) {
	p.passLength = length
}
func (p *Options) WithUpperCase(includeUpperCase bool) {
	p.includeUpperCase = includeUpperCase
}
func (p *Options) WithLowerCase(includeLowerCase bool) {
	p.includeLowerCase = includeLowerCase
}
func (p *Options) WithNumbers(includeNumbers bool) {
	p.includeNumbers = includeNumbers
}
func (p *Options) WithSpecialCharacters(includeSpecialCharacters bool) {
	p.includeSpecialCharacters = includeSpecialCharacters
}
func (p *Options) WithMaxNumbers(maxNumbers int) {
	p.maxNumbers = maxNumbers
}
func (p *Options) WithMaxSpecialCharacters(maxSpecialCharacters int) {
	p.maxSpecialCharacters = maxSpecialCharacters
}

func (p *Options) GeneratePassword() (Pass string) {
	rand.Seed(time.Now().UnixNano())
	if !p.includeUpperCase && !p.includeLowerCase && !p.includeNumbers && !p.includeSpecialCharacters {
		log.Fatalf("all options are disabled")
	}
	s := p.SeparationCounter()
	c := combin.Combinations(p.passLength, s-1)

	var countForEach []int
	prev := 0
	selected := c[rand.Intn(len(c))]
	for _, v := range selected {
		countForEach = append(countForEach, v-prev+1)
		prev = v + 1
	}
	countForEach = append(countForEach, p.passLength-prev+1)

	var ret []byte
	var v int
	if p.includeNumbers {
		v, countForEach = utils.Shift(countForEach)
		if v > p.maxNumbers && p.maxNumbers >= 0 {
			countForEach[0] += v - p.maxNumbers
			v = p.maxNumbers
		}
		ret = append(ret, utils.RandStringBytes(v, []byte(numberBytes))...)
	}
	if p.includeSpecialCharacters {
		v, countForEach = utils.Shift(countForEach)
		if v > p.maxSpecialCharacters && p.maxSpecialCharacters >= 0 {
			countForEach[0] += v - p.maxSpecialCharacters
			v = p.maxSpecialCharacters
		}
		ret = append(ret, utils.RandStringBytes(v, []byte(specialCharacterBytes))...)
	}
	if p.includeUpperCase && p.includeLowerCase {
		var tot int
		for _, v := range countForEach {
			tot += v
		}
		mid := rand.Int() % tot
		ret = append(ret, utils.RandStringBytes(mid, []byte(alphabetUpperCaseBytes))...)
		ret = append(ret, utils.RandStringBytes(tot-mid+1, []byte(alphabetLowerCaseBytes))...)

	} else if p.includeUpperCase {
		v, countForEach = utils.Shift(countForEach)
		ret = append(ret, utils.RandStringBytes(v, []byte(alphabetUpperCaseBytes))...)
	} else if p.includeLowerCase {
		v, countForEach = utils.Shift(countForEach)
		ret = append(ret, utils.RandStringBytes(v, []byte(alphabetLowerCaseBytes))...)
	}
	rand.Shuffle(len(ret), func(i, j int) {
		ret[i], ret[j] = ret[j], ret[i]
	})
	return string(ret)
}

func (p *Options) SeparationCounter() int {
	return utils.BoolCounter(p.includeSpecialCharacters, p.includeNumbers, p.includeUpperCase, p.includeLowerCase)
}
