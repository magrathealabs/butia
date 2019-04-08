package boot

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magrathealabs/butia/spec"
)

type GinSuite struct {
	spec.Suite
}

func (suite *GinSuite) TestGin() {
	suite.NotPanics(Gin)
	suite.Equal("release", gin.Mode())
}

func TestGinSuite(t *testing.T) {
	spec.Run(t, new(GinSuite))
}
