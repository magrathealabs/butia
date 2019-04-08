package butia

import "github.com/gin-gonic/gin"

// GinHandleFunc create a gin func over a butia func
func GinHandleFunc(context func(*Context)) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		context(NewContext(ginContext))
	}
}

// GinHandleFuncs create a slice of gin func over a butia func
func GinHandleFuncs(contexts []func(*Context)) (funcs []gin.HandlerFunc) {
	for _, context := range contexts {
		funcs = append(funcs, GinHandleFunc(context))
	}
	return
}
