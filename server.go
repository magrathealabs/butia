package butia

import "github.com/gin-gonic/gin"

// Server can receive web requests
type Server struct {
	*gin.Engine
}

// NewServer create a new gin router without configurations
func NewServer() *Server {
	return &Server{gin.Default()}
}
