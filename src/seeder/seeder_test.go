package seeder

import (
	"context"
	"github.com/mediocregopher/radix/v4"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net"
	"testing"
)

type SeederTestSuite struct {
	suite.Suite
}

type RedisClientMock struct {
	mock.Mock
}

func (m *RedisClientMock) Addr() net.Addr {
	args := m.Called()
	return args.Get(0).(net.Addr)
}

func (m *RedisClientMock) Do(ctx context.Context, action radix.Action) error {
	args := m.Called(ctx, action)
	return args.Error(0)
}

func (m *RedisClientMock) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (suite *SeederTestSuite) TestScan() {
	clientMock := &RedisClientMock{}
	clientMock.
		On(
			"Do",
			mock.Anything,
			radix.FlatCmd(nil, "SET", "key", "value"),
		).Times(5).Return(nil).
		On(
			"Do",
			mock.Anything,
			radix.FlatCmd(nil, "EXPIRE", "key", 100),
		).Times(5).Return(nil)

	seeder := NewSeeder(clientMock, zerolog.Nop())

	seeder.Seed(NewGenericRecordGenerator(
		5,
		NewEnumStringGenerator("key"),
		NewEnumStringGenerator("value"),
		NewIntRangeGenerator(100, 101),
	))

}

func TestSeederTestSuite(t *testing.T) {
	suite.Run(t, new(SeederTestSuite))
}
