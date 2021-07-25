package seeder

import (
	"math/rand"
	"strconv"
)

// IntRangeStringGenerator generates a random integer in provided range as a string
type IntRangeStringGenerator struct {
	min int
	max int
}

// NewIntRangeStringGenerator creates IntRangeStringGenerator
func NewIntRangeStringGenerator(min, max int) IntRangeStringGenerator {
	return IntRangeStringGenerator{min, max}
}

func (i IntRangeStringGenerator) generate() string {
	return strconv.Itoa(rand.Intn(i.max-i.min) + i.min)
}
