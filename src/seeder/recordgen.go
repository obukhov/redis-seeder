package seeder

// RecordGenerator generates records in async way to seed
type RecordGenerator interface {
	generate() <-chan Record
}

// StringGenerator implementations generate random strings with their given strategy
type StringGenerator interface {
	generate() string
}

// IntGenerator implementations generate random int with their given strategy
type IntGenerator interface {
	generate() int
}

// Record represents simple redis record to seed
type Record struct {
	key   string
	ttl   int
	value string
}

// NewGenericRecordGenerator creates GenericRecordGenerator
func NewGenericRecordGenerator(repeatCount int, keyGenerator, valueGenerator StringGenerator, ttlGenerator IntGenerator) GenericRecordGenerator {
	return GenericRecordGenerator{
		repeatCount,
		keyGenerator,
		valueGenerator,
		ttlGenerator,
	}
}

// GenericRecordGenerator generates <repeatCount> records with given key, value and ttl strategy
type GenericRecordGenerator struct {
	repeatCount    int
	keyGenerator   StringGenerator
	valueGenerator StringGenerator
	ttlGenerator   IntGenerator
}

func (g GenericRecordGenerator) generate() <-chan Record {
	result := make(chan Record)
	go func() {
		defer close(result)
		for i := 0; i < g.repeatCount; i++ {
			result <- Record{
				g.keyGenerator.generate(),
				g.ttlGenerator.generate(),
				g.valueGenerator.generate(),
			}

		}
	}()

	return result
}
