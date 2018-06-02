package test

import (
	"testing"

	"../models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	user models.User
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (suite *UserTestSuite) SetupTest() {
	suite.user = models.User{1, "tom"}
}
func (suite *UserTestSuite) TestUserId() {
	assert.Equal(suite.T(), 1, suite.user.Id)
}

func (suite *UserTestSuite) TestUserFirstname() {
	assert.Equal(suite.T(), "tom", suite.user.Firstname)
}
