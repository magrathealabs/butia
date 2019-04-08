package boot

import (
	"os"
	"testing"

	"github.com/magrathealabs/butia/spec"
)

type EnvSuite struct {
	spec.Suite
}

func (suite *EnvSuite) SetupTest() {
	suite.Suite.SetupTest()
}

func (suite *EnvSuite) TestEnv() {
	suite.Equal("", os.Getenv("ENV_TEST_SUITE"))
	suite.NotPanics(Env)
	suite.Equal("butia", os.Getenv("ENV_TEST_SUITE"))
}

func TestSpecTestSuite(t *testing.T) {
	spec.Run(t, new(EnvSuite))
}
