package boot

import "github.com/gin-gonic/gin"

// Gin bootstrap
func Gin() {
	gin.SetMode(gin.ReleaseMode)
}
