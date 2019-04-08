package butia

import "github.com/gin-gonic/gin"

// Context store params and can submit a web response
type Context struct {
	*gin.Context
}

// NewContext initialize a butia context
func NewContext(context *gin.Context) *Context {
	return &Context{context}
}
