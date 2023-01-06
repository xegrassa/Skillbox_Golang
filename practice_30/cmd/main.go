/*
Научиться:

	работать с запросами POST, GET, PUT, DELETE;
	применять принципы написания обработчиков HTTP-запросов.

Что нужно сделать

	Напишите HTTP-сервис, который принимает входящие соединения с JSON-данными и обрабатывает их следующим образом:
*/
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"skillbox/practice_30/pkg/controls"
	"syscall"

	"github.com/go-chi/chi"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, os.Interrupt)

	r := chi.NewRouter()
	r.Post("/create", controls.CreateUser)
	r.Post("/make_friends", controls.MakeFriends)
	r.Delete("/user", controls.DeleteUser)
	r.Get("/friends/{userId}", controls.GetFriends)

	r.Put("/{userId}", controls.UpdateUser)

	go func() {
		fmt.Println("server running")
		http.ListenAndServe(":8080", r)
	}()

	<-sigs
	fmt.Println("server shutdown")
}
