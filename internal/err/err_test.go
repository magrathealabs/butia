package err

import (
	"errors"
	"testing"

	"github.com/magrathealabs/butia/spec"
)

type ErrSuite struct {
	spec.Suite
}

func (suite *ErrSuite) SetupTest() {
	suite.Suite.SetupTest()
}

func (suite *ErrSuite) TestRaiseErr() {
	suite.NotPanics(func() {
		RaiseErr(nil)
	})

	suite.Panics(func() {
		RaiseErr(errors.New("Testing"))
	})
}

func TestErrSuite(t *testing.T) {
	spec.Run(t, new(ErrSuite))
}
