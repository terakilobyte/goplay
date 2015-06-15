package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/", testroute)
	r.GET("/fibonacci/:which", fibon)
	r.GET("/:name", testParam)
	r.POST("/login", login)
	r.Run(":8080")
}
