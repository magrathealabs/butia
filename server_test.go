package butia

import (
	"testing"

	"github.com/magrathealabs/butia/spec"
)

type ServerSuite struct {
	spec.Suite
}

func (suite *ServerSuite) TestNewBasicServer() {
	suite.NotNil(NewBasicServer())
	suite.Equal(0, len(NewBasicServer().Routes()))
}

func (suite *ServerSuite) TestNewServerWithStatic() {
	suite.NotNil(NewServerWithStatic())
	suite.Equal(2, len(NewServerWithStatic().Routes()))
}

func (suite *ServerSuite) TestNewCompleteServer() {
	suite.NotNil(NewCompleteServer())
	suite.Equal(2, len(NewCompleteServer().Routes()))
}

func TestServerSuite(t *testing.T) {
	spec.Run(t, new(ServerSuite))
}
