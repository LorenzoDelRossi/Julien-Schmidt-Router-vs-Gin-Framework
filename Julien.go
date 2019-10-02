package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"net/url"
)

func main() {
	router := httprouter.New()
	router.POST("/v1/add", posthandler)
	http.ListenAndServe(":9001", router)

}
//posthandler function parses the POST request and prints the "name" value
func posthandler(w http.ResponseWriter, r *http.Request, c httprouter.Params) {
	r.ParseForm()
	fmt.Println("Response: ", r.PostForm.Get("name"))
}
