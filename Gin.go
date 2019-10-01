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
	go gingo(500)
	r.Run(":9002") // listen on port :9002
}

//gingo sends N POST requests to the Gin server on port :9002 and takes the time spent in the handling process
func gingo(N int) {
	var i int
	starts := time.Now()
	for i = 0; i < N; i++ {
		go func() {
			_, err := http.PostForm("http://localhost:9002/v1/add", url.Values{"name": {"canovaccio"}})
			if err != nil {
				//handle postform error
				log.Fatalln(err)
			}
		}()
	}
	fmt.Println("CHIAMATE CONCORRENTI GESTITE: ", i)
	fmt.Println("TEMPO IN GINGO: ", time.Since(starts))
}
