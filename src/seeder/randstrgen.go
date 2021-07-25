package seeder

import "math/rand"

// RandStringGenerator generates random string of given length range with given symbols
type RandStringGenerator struct {
	minLength    int
	maxLength    int
	symbols      []rune
	symbolsCount int
}

// NewRandStringGenerator creates RandStringGenerator
func NewRandStringGenerator(minLength, maxLength int, symbols ...rune) RandStringGenerator {
	return RandStringGenerator{
		minLength,
		maxLength,
		symbols,
		len(symbols),
	}
}

func (s RandStringGenerator) generate() string {
	length := rand.Intn(s.maxLength-s.minLength) + s.minLength
	result := make([]rune, 0, length)
	for i := 0; i < length; i++ {
		result = append(result, s.symbols[rand.Intn(s.symbolsCount)])
	}

	return string(result)
}
