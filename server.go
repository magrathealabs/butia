package butia

import (
	"path/filepath"

	"github.com/krakenlab/scopo/libs/env"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

// Server can receive web requests
type Server struct {
	*gin.Engine
}

// NewBasicServer create a new gin router without configurations.
// Use this for API applications
func NewBasicServer() *Server {
	return &Server{gin.Default()}
}

// NewServerWithStatic create a new gin router without configurations.
// Use this for API applications with assets (REACT)
func NewServerWithStatic() (server *Server) {
	server = NewBasicServer()
	staticFilePath := filepath.Join(env.GetString("GOPATH", "/go"), "src", env.GetString("GIT_SRV", "github.com"), env.GetString("GIT_ORG", "magrathealabs"), env.GetString("APP_NAME", "web"), "static")
	server.Static("/static", staticFilePath)
	return
}

// NewCompleteServer create a new gin with gin-template configuration.
// Use this for web applications.
func NewCompleteServer() (server *Server) {
	server = NewServerWithStatic()
	server.HTMLRender = gintemplate.Default()
	return
}

// Controller register a basic controller into a path
func (server *Server) Controller(path string, controller Controller) {
	server.GET(path, controller.Index)
	server.GET(path+"/new", controller.New)
	server.POST(path+"/create", controller.Create)
	server.GET(path+"/show/:id", controller.Show)
	server.GET(path+"/edit/:id", controller.Edit)
	server.PUT(path+"/update/:id", controller.Update)
	server.PATCH(path+"/update/:id", controller.Update)
	server.DELETE(path+"/destroy/:id", controller.Destroy)
}

// GET endpint
func (server *Server) GET(path string, funcs ...func(*Context)) {
	server.Engine.GET(path, GinHandleFuncs(funcs)...)
}

// POST endpint
func (server *Server) POST(path string, funcs ...func(*Context)) {
	server.Engine.POST(path, GinHandleFuncs(funcs)...)
}

// PUT endpint
func (server *Server) PUT(path string, funcs ...func(*Context)) {
	server.Engine.PUT(path, GinHandleFuncs(funcs)...)
}

// PATCH endpint
func (server *Server) PATCH(path string, funcs ...func(*Context)) {
	server.Engine.PATCH(path, GinHandleFuncs(funcs)...)
}

// DELETE endpint
func (server *Server) DELETE(path string, funcs ...func(*Context)) {
	server.Engine.DELETE(path, GinHandleFuncs(funcs)...)
}
