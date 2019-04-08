package butia

import (
	"testing"

	"github.com/magrathealabs/butia/spec"
)

type ServerSuite struct {
	spec.Suite
}

func (suite *ServerSuite) TestNewServer() {
	suite.NotNil(NewServer())
	suite.Equal(0, len(NewServer().Routes()))
}

func TestServerSuite(t *testing.T) {
	spec.Run(t, new(ServerSuite))
}
