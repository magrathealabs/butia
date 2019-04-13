package butia

import (
	"testing"

	"github.com/magrathealabs/butia/spec"
)

type ServerSuite struct {
	spec.Suite
}

func (suite *ServerSuite) TestServerRootDirectoryPath() {
	server := NewCompleteServer()
	suite.Contains(server.RootDirectoryPath(), "github.com/magrathealabs/web")
}

func (suite *ServerSuite) TestServerStaticDirectoryPath() {
	server := NewCompleteServer()
	suite.Contains(server.StaticDirectoryPath(), "static")
	suite.Contains(server.StaticDirectoryPath(), server.RootDirectoryPath())
}

func (suite *ServerSuite) TestServerViewDirectoryPath() {
	server := NewCompleteServer()
	suite.Contains(server.ViewDirectoryPath(), "views")
	suite.Contains(server.ViewDirectoryPath(), server.RootDirectoryPath())
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
