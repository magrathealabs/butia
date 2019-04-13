package controllers

import (
	"github.com/magrathealabs/butia"
)

// ApplicationController is an generic controller
type ApplicationController struct {
	butia.ApplicationController
}

// NewApplicationController init the main controller
func NewApplicationController() *ApplicationController {
	return &ApplicationController{}
}

// BuildRoutes from this application
func (controller *ApplicationController) BuildRoutes() *butia.Server {
	server := butia.NewCompleteServer()

	NewHomeController().BuildRoutes(server)

	return server
}
