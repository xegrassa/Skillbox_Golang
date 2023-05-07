/*
Научиться:

	работать с запросами POST, GET, PUT, DELETE;
	применять принципы написания обработчиков HTTP-запросов.

Что нужно сделать

	Напишите HTTP-сервис, который принимает входящие соединения с JSON-данными и обрабатывает их следующим образом:
*/
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"skillbox/practice_30/internal/controls"
	"skillbox/practice_30/internal/storage"

	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	sPointer := storage.NewMemStorage()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "storage", sPointer)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/create", controls.CreateUser)
	r.Post("/make_friends", controls.MakeFriends)
	r.Delete("/user", controls.DeleteUser)
	r.Get("/friends/{userId}", controls.GetFriends)
	r.Put("/{userId}", controls.UpdateUser)

	go func() {
		addr := "127.0.0.1:8080"
		log.Printf("server running on %v\n", addr)
		http.ListenAndServe(addr, r)
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, os.Interrupt)
	<-sigs

	log.Println("server shutdown")
}
