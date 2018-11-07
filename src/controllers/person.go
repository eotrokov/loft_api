package controllers

import (
	"net/http"
	"github.com/google/uuid"
	"../repositories"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)


func Routes() *chi.Mux  {
	router:= chi.NewRouter()
	router.Get("/{id}", GetPerson)
	router.Delete("/{id}", DeletePerson)
	router.Post("/", CreatePerson)
	router.Get("/", GetAllPerson)
	return router
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	id:= uuid.MustParse( chi.URLParam(r, "id"))
	person := personRepo.GetPerson(id)
	render.JSON(w,r,person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	response:= make(map[string]string)
	response["message"]="Delete successfully"
	render.JSON(w,r,response)
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	response:= make(map[string]string)
	response["message"]="Create successfully"
	render.JSON(w,r,response)
}
func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	person := personRepo.GetPersons()
	render.JSON(w,r,person)
}
