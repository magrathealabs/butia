package spec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type SpecTestSuite struct {
	Suite
}

func (suite *SpecTestSuite) TestSetupTest() {
	assert.NotNil(suite.T(), suite.Expect)
}

func TestSpecTestSuite(t *testing.T) {
	Run(t, new(SpecTestSuite))
}
