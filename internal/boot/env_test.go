package boot

import (
	"os"
	"testing"

	"github.com/magrathealabs/butia/spec"
)

type EnvSuite struct {
	spec.Suite
}

func (suite *EnvSuite) TestEnv() {
	suite.Equal("", os.Getenv("ENV_SUITE_TEST_ENV"))
	suite.NotPanics(Env)
	suite.Equal("butia", os.Getenv("ENV_SUITE_TEST_ENV"))
}

func (suite *EnvSuite) TestGetenv() {
	suite.NotPanics(Env)

	suite.Equal("OUPS", Getenv("ENV_SUITE_TEST_GETENV", "OUPS"))
	suite.Equal("butia", Getenv("ENV_SUITE_TEST_ENV", "OUPS"))
}

func TestEnvSuite(t *testing.T) {
	spec.Run(t, new(EnvSuite))
}
