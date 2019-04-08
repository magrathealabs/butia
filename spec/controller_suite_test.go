package spec

import (
	"testing"

	"github.com/gin-gonic/gin"
)

type SpecControllerTestSuite struct {
	ControllerSuite
}

func (suite *SpecControllerTestSuite) SetupTest() {
	suite.Suite.SetupTest()

	suite.Server = gin.Default()
}

func (suite *SpecControllerTestSuite) TestGET() {
	req := suite.GET("/", "")
	suite.Equal(404, req.Code)
	suite.Equal("404 page not found", req.Body.String())
}

func (suite *SpecControllerTestSuite) TestPOST() {
	req := suite.POST("/", "")
	suite.Equal(404, req.Code)
	suite.Equal("404 page not found", req.Body.String())
}

func (suite *SpecControllerTestSuite) TestPUT() {
	req := suite.PUT("/", "")
	suite.Equal(404, req.Code)
	suite.Equal("404 page not found", req.Body.String())
}

func (suite *SpecControllerTestSuite) TestPATCH() {
	req := suite.PATCH("/", "")
	suite.Equal(404, req.Code)
	suite.Equal("404 page not found", req.Body.String())
}

func (suite *SpecControllerTestSuite) TestDELETE() {
	req := suite.DELETE("/", "")
	suite.Equal(404, req.Code)
	suite.Equal("404 page not found", req.Body.String())
}

func TestSpecControllerTestSuite(t *testing.T) {
	Run(t, new(SpecControllerTestSuite))
}
