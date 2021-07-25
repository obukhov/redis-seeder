package seeder

import "math/rand"

// IntRangeGenerator generates a random integer in provided range
type IntRangeGenerator struct {
	min int
	max int
}

// NewIntRangeGenerator creates IntRangeGenerator
func NewIntRangeGenerator(min, max int) IntRangeGenerator {
	return IntRangeGenerator{min, max}
}

func (i IntRangeGenerator) generate() int {
	return rand.Intn(i.max-i.min) + i.min
}
