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
	router := gin.Default()
	router.GET("v1/julien", func(c *gin.Context) {
		go juliengo(500)
	})
	router.GET("v1/gin", func(c *gin.Context) {
		go gingo(500)
	})
	http.ListenAndServe("localhost:8080", router)
}

//juliengo sends N POST requests to the Julien server on port :9001 and takes the time spent in the handling process
func juliengo(N int) {
	var i int
	starts := time.Now()
	for i = 0; i < N; i++ {
		go func() {
			_, err := http.PostForm("http://localhost:9001/v1/add", url.Values{"name": {"OK"}})
			if err != nil {
				log.Fatalln(err)
			}
		}()
	}
	fmt.Println("CHIAMATE CONCORRENTI GESTITE:", i)
	fmt.Println("TEMPO IN JULIENGO: ", time.Since(starts))
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
