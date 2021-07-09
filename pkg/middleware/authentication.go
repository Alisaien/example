package middleware

import "github.com/gin-gonic/gin"

func Authenticate(ctx *gin.Context) {
	// If not authenticated
	// write error response & return

	// else
	ctx.Next()
}
