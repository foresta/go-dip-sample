package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// /users/
	users := r.PathPrefix("/users").Subrouter()
	users.HandleFunc("/", UserListHandler).Methods("GET")
	users.HandleFunc("/", CreateUserHandler).Methods("POST")
	users.HandleFunc("/{user_id}", ShowUserHandler).Methods("GET")
	users.HandleFunc("/{user_id}", UpdateUserHandler).Methods("PUT")
	users.HandleFunc("/{user_id}", DeleteUserHandler).Methods("DELETE")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func UserListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UserListHandler")
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateUserHandler")

}

func ShowUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ShowUserHandler")

}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UpdateUserHandler")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeleteUserHandler")

}
