package repositories

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MongoRepoTestSuite struct {
	suite.Suite
	mock *mock.Mock
}

func SetupTestSuite(t *testing.T) *MongoRepoTestSuite {
	return &MongoRepoTestSuite{
		mock: &mock.Mock{},
	}
}

func TestMongoRepoTestSuite(t *testing.T) {
	suite.Run(t, SetupTestSuite(t))
}
