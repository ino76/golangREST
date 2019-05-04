package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	localhost := "localhost:8080"
	r := httprouter.New()
	r.GET("/", index)
	fmt.Println("server started at " + localhost)
	http.ListenAndServe(localhost, r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Ahoj Dane .. go rest :)\n")
}
