package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"learn.go/GolandProjects/userserv/internal/user"
)

func main() {
	log.Println("create router")
	router := httprouter.New()
	log.Println("register user handler")
	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Println("start application")
	listener, err := net.Listen("tcp", ":1212")
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router, // Добавлено
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Println("server is listening port 1212")
	log.Fatalln(server.Serve(listener))
}
