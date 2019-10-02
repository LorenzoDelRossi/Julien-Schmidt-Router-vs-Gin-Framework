package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"time"
	"net/url"
)

func main() {
	router := httprouter.New()
	router.POST("/v1/add", posthandler)
	go juliengo(500)
	http.ListenAndServe(":9001", router)

}
//posthandler function parses the POST request and prints the "name" value
func posthandler(w http.ResponseWriter, r *http.Request, c httprouter.Params) {
	r.ParseForm()
	fmt.Println("Response: ", r.PostForm.Get("name"))
}

//juliengo sends N POST requests to the Julien server on port :9001 and takes the time spent in the handling process
func juliengo(N int) {
	var i int
	starts := time.Now()
	for i = 0; i < N; i++ {
		go func() {
			_, err := http.PostForm("http://localhost:9001/v1/add", url.Values{"name": {"prova"}})
			if err != nil {
				log.Fatalln(err)
			}
		}()
	}
	fmt.Println("CHIAMATE CONCORRENTI GESTITE:", i)
	fmt.Println("TEMPO IN JULIENGO: ", time.Since(starts))
}
