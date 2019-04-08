package butia

import (
	"testing"

	"github.com/magrathealabs/butia/spec"
)

type ContextSuite struct {
	spec.Suite
}

func (suite *ContextSuite) TestNewContext() {
	suite.NotNil(NewContext(nil))
}

func TestContextSuite(t *testing.T) {
	spec.Run(t, new(ContextSuite))
}
