package controllers

import (
	"net/http"
	"testing"

	"github.com/magrathealabs/butia/spec"
)

type HomeControllerSuite struct {
	spec.ControllerSuite
}

func (suite *HomeControllerSuite) SetupTest() {
	suite.ControllerSuite.SetupTest()

	suite.Server = NewApplicationController().BuildRoutes()
}

func (suite *HomeControllerSuite) TestRoot() {
	res := suite.GET("/", "")

	suite.Equal(http.StatusOK, res.Code)

	body := res.Body.String()

	suite.Contains(body, "Butia App")
	suite.Contains(body, "Fork this project!")
}

func TestHomeControllerSuite(t *testing.T) {
	spec.Run(t, new(HomeControllerSuite))
}
