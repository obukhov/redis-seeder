package seeder

import "math/rand"

// EnumStringGenerator generates one of provided string randomly
type EnumStringGenerator struct {
	values []string
	count  int
}

// NewEnumStringGenerator creates EnumStringGenerator
func NewEnumStringGenerator(values ...string) EnumStringGenerator {
	return EnumStringGenerator{
		values,
		len(values),
	}
}

func (e EnumStringGenerator) generate() string {
	return e.values[rand.Intn(e.count)]
}
