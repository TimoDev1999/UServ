package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"learn.go/GolandProjects/userserv/internal/handlers"
)

const (
	usersURL = "/users"
	usserURL = "/user/:uuid"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.POST(usserURL, h.CreateUser)
	router.GET(usserURL, h.GetUserByUUID)
	router.PUT(usserURL, h.UpdateUser)
	router.PATCH(usserURL, h.PartiallyUpdateUser)
	router.DELETE(usserURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is list of users"))

}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("this is create user"))

}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is user by uuid"))

}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is update user"))

}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is delete user"))

}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is part.update user"))

}
