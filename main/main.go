package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	ApiwithGoChi "ApiWithGoChi.com"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
)

var mh *ApiwithGoChi.MongoHandler

func registerRoutes() http.Handler {
	r := chi.NewRouter()
	r.Route("/person", func(r chi.Router) {
		r.Get("/", getAllPerson)
		r.Post("/", addPerson)
	})
	return r
}

func getAllPerson(w http.ResponseWriter, r *http.Request) {
	persons := mh.Get(bson.M{})
	json.NewEncoder(w).Encode(persons)
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	existingPerson := &ApiwithGoChi.Person{}

	var person ApiwithGoChi.Person
	json.NewDecoder(r.Body).Decode(&person)
	person.CreatedOn = time.Now()
	err := mh.GetOne(existingPerson, bson.M{"phone_number": person.PhoneNumber})
	if err == nil {
		http.Error(w, fmt.Sprintf("Contact with phone number %s already exist", person.PhoneNumber), 400)
		return
	}
	_, err = mh.AddOne(&person)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	w.Write([]byte("Contact created successfully"))
	w.WriteHeader(201)
}

func main() {
	mongoDbConnection := "mongodb://localhost:27017"
	mh = ApiwithGoChi.NewHandler(mongoDbConnection)
	r := registerRoutes()
	log.Fatal(http.ListenAndServe(":3060", r))

}
