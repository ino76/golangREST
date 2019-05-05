package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/ino76/golangREST/models"
	"github.com/julienschmidt/httprouter"

	"gopkg.in/mgo.v2"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Ahoj Dane .. go rest :)\n")
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	// fetch user
	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()

	uc.session.DB("go-web-dev-db").C("users").Insert(u)

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Write code to delete user\n")
}
