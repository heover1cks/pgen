package password

import (
	cryptoRand "crypto/rand"
	"gonum.org/v1/gonum/stat/combin"
	"idpassgen/pkg/utils"
	"log"
	"math"
	"math/big"
	"math/rand"

	"time"
)

const (
	alphabetUpperCaseBytes        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetLowerCaseBytes        = "abcdefghijklmnopqrstuvwxyz"
	numberBytes                   = "1234567890"
	specialCharacterBytes         = "!@#$%^&*<>?-=_+"
	extendedSpecialCharacterBytes = "~!@#$%^&*()_+`-={}|[]\\\\:\\\"<>?,./"
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

type Password struct {
	inserted  []bool
	generated []byte
	seed      int64
	length    int
}

func NewPassword(initial []byte, length int, seed int64) *Password {
	return &Password{
		inserted:  make([]bool, length),
		generated: initial,
		seed:      seed,
		length:    length,
	}
}
func (p *Password) Get() string {
	return string(p.generated)
}

// TODO: need to check what is better if Password contains seed value or generate seed everytime insertAtRandomPosition is called
func (p *Password) insertAtRandomPosition(target byte) {
	pos := rand.Int() % p.length
	if p.inserted[pos] {
		p.insertAtRandomPosition(target)
	} else {
		front := append(p.generated[:pos], target)
		if pos == p.length-1 {
			p.generated = front
		} else {
			back := p.generated[pos+1:]
			p.generated = append(front, back...)
		}
	}
}

func (p *Options) GeneratePasswordV3() (Pass string) {
	seed, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
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

	var others []byte
	var alphabets int
	var v int
	var pw *Password

	if p.includeNumbers {
		// Cap number count to max
		v, countForEach = utils.Shift(countForEach)
		if v > p.maxNumbers && p.maxNumbers >= 0 {
			countForEach[0] += v - p.maxNumbers
			v = p.maxNumbers
		}
		// append number bytes
		others = utils.RandStringBytes(v, []byte(numberBytes))
	}

	if p.includeSpecialCharacters {
		v, countForEach = utils.Shift(countForEach)

		// Cap special character count to max
		if v > p.maxSpecialCharacters && p.maxSpecialCharacters >= 0 {
			countForEach[0] += v - p.maxSpecialCharacters
			v = p.maxSpecialCharacters
		}
		// append special character bytes
		others = append(others, utils.RandStringBytes(v, []byte(specialCharacterBytes))...)
	}
	if p.includeUpperCase || p.includeLowerCase {
		var alphabetByte []byte
		switch {
		case p.includeUpperCase && p.includeLowerCase:
			alphabetByte = []byte(alphabetUpperCaseBytes + alphabetLowerCaseBytes)
			break
		case p.includeUpperCase && !p.includeLowerCase:
			alphabetByte = []byte(alphabetUpperCaseBytes)
			break
		case !p.includeUpperCase && p.includeLowerCase:
			alphabetByte = []byte(alphabetLowerCaseBytes)
			break
		default:
			panic("")
		}
		for _, i := range countForEach {
			alphabets += i
		}

		pw = NewPassword(utils.RandStringBytes(alphabets+len(others), alphabetByte), p.passLength, seed.Int64())
		for _, char := range others {
			pw.insertAtRandomPosition(char)
		}
	}
	return pw.Get()
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
