package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	r := gin.Default()
	r.POST("/v1/add", func(c *gin.Context) {
		c.Request.ParseForm()
		message := c.Request.PostFormValue("name")
		fmt.Println("Response: ", message)
	})
	r.Run(":9002") // listen on port :9002
}

