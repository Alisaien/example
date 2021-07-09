package routes

import (
	"github.com/Alisaien/example/pkg/middleware"
	"github.com/Alisaien/example/pkg/routes/forum"
	"github.com/gin-gonic/gin"
)

var Mux *gin.Engine

func init() {
	Mux = gin.Default()

	// users can view posts without being authenticated
	Mux.GET("/post", forum.GetPost)

	// require auth for put/deletes
	reqAuth := Mux.Group("/post")
	reqAuth.Use(middleware.Authenticate)

	reqAuth.PUT("/post", forum.PutPost)
}
