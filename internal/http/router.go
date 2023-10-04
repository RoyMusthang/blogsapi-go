package http

import "github.com/gin-gonic/gin"

func handler() {
	router := gin.Default()

	router.GET("/user", func(c *gin.Context) {})
	router.GET("/login", func(c *gin.Context) {})
	router.GET("/categories", func(c *gin.Context) {})
	router.GET("/post", func(c *gin.Context) {})

}
