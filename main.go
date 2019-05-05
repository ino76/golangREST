package main

import (
	"fmt"
	"net/http"

	"github.com/ino76/golangREST/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	localhost := "localhost:8080"
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	fmt.Println("server started at " + localhost)
	http.ListenAndServe(localhost, r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return s
}
