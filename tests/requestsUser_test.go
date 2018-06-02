package test

import (
	"testing"

	"../models"
	"../requests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RequestsUserTestSuite struct {
	suite.Suite
	allUsers []models.User
}

func TestRequestsUserTestSuite(t *testing.T) {
	suite.Run(t, new(RequestsUserTestSuite))
}

func (suite *RequestsUserTestSuite) SetupTest() {
	suite.allUsers = requests.GetAllUsers()
}
func (suite *RequestsUserTestSuite) TestLengthUSers() {
	assert.NotEqual(suite.T(), 1, len(suite.allUsers))
}
