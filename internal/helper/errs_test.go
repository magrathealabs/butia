package helper

import (
	"errors"
	"testing"

	"github.com/magrathealabs/butia/spec"
)

type ErrsSuite struct {
	spec.Suite
}

func (suite *ErrsSuite) SetupTest() {
	suite.Suite.SetupTest()
}

func (suite *ErrsSuite) TestRaiseErr() {
	suite.NotPanics(func() {
		RaiseErr(nil)
	})

	suite.Panics(func() {
		RaiseErr(errors.New("Testing"))
	})
}

func TestSpecTestSuite(t *testing.T) {
	spec.Run(t, new(ErrsSuite))
}
