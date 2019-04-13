package controllers

import (
	"net/http"

	"github.com/magrathealabs/butia"
)

// HomeController handle requests at the root
type HomeController struct {
	ApplicationController
}

// NewHomeController init a controller
func NewHomeController() *HomeController {
	return &HomeController{}
}

// BuildRoutes of HomeController
func (controller *HomeController) BuildRoutes(server *butia.Server) {
	server.GET("/", controller.Root)
}

// Root GET /
func (controller *HomeController) Root(c *butia.Context) {
	c.HTML(http.StatusOK, "root", butia.H{})
}
