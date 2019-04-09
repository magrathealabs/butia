package butia

import (
	"testing"

	"github.com/magrathealabs/butia/spec"
)

type ApplicationControllerSuite struct {
	spec.ControllerSuite
}

func (suite *ApplicationControllerSuite) TestNewApplicationController() {
	suite.NotNil(NewApplicationController(NewBasicServer()))
}

func (suite *ApplicationControllerSuite) SetupTest() {
	suite.ControllerSuite.SetupTest()

	controller := NewApplicationController(NewBasicServer())
	controller.Setup()
	suite.Server = controller.Server
}

func (suite *ApplicationControllerSuite) TestSetup() {
	controller := NewApplicationController(NewBasicServer())
	controller.Setup()

	suite.Equal(8, len(controller.Server.Routes()))
}

func (suite *ApplicationControllerSuite) TestIndex() {
	response := suite.GET("/", "")

	suite.Equal(200, response.Code)
	suite.Equal("Index must be implemented", response.Body.String())

}

func (suite *ApplicationControllerSuite) TestCreate() {
	response := suite.POST("/create", "")

	suite.Equal(200, response.Code)
	suite.Equal("Create must be implemented", response.Body.String())
}

func (suite *ApplicationControllerSuite) TestNew() {
	response := suite.GET("/new", "")

	suite.Equal(200, response.Code)
	suite.Equal("New must be implemented", response.Body.String())
}

func (suite *ApplicationControllerSuite) TestEdit() {
	response := suite.GET("/edit/2", "")

	suite.Equal(200, response.Code)
	suite.Equal("Edit must be implemented for 2", response.Body.String())
}

func (suite *ApplicationControllerSuite) TestShow() {
	response := suite.GET("/show/2", "")

	suite.Equal(200, response.Code)
	suite.Equal("Show must be implemented for 2", response.Body.String())
}

func (suite *ApplicationControllerSuite) TestUpdate() {
	response := suite.PATCH("/update/2", "")

	suite.Equal(200, response.Code)
	suite.Equal("Update must be implemented for 2", response.Body.String())
}

func (suite *ApplicationControllerSuite) TestDestroy() {
	response := suite.DELETE("/destroy/2", "")

	suite.Equal(200, response.Code)
	suite.Equal("Destroy must be implemented for 2", response.Body.String())
}

func TestApplicationControllerSuite(t *testing.T) {
	spec.Run(t, new(ApplicationControllerSuite))
}
