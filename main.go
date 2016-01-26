package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "hello")
}

func main() {
	router := httprouter.New()
	router.GET("/", index)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
