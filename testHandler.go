package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func testroute(c *gin.Context) {
	c.String(http.StatusOK, "Hello world!")
}

func testParam(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func login(c *gin.Context) {
	var json LoginJSON
	c.Bind(&json)
	if json.User == "Nathan" && json.Password == "hellyeah" {
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
	}
}

func fibon(c *gin.Context) {
	param := c.Param("which")

	which, err := strconv.Atoi(param)
	if err != nil {
		println(err)
	}
	a := 0
	b := 1
	for i := 0; i < which; i++ {
		tmp := b
		b = a + b
		a = tmp
	}

	c.String(http.StatusOK, "Your fibonnaci is %v", b)

}
