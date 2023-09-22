package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golamrabbiazad/blogs-w-go/src/handlers"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.GET("/blogs", handlers.GetBlogs)
	r.POST("/blogs", handlers.PostBlog)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
