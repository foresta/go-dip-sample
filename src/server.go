// server.go

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/foresta/go-dip-sample/src/memory"
	"github.com/foresta/go-dip-sample/src/user"
	"github.com/gorilla/mux"
)

var user_repository user.Repository

func main() {

	user_repository = memory.NewUserRepository()

	r := mux.NewRouter()
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	r.HandleFunc("/users", listUsersHandler).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	type User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	type response struct {
		Users []*User `json:"users"`
	}

	// user一覧取得
	users := user_repository.FindAll()

	responseUsers := make([]*User, 0, len(users))
	for _, user := range users {
		u := &User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
		responseUsers = append(responseUsers, u)
	}
	res := &response{Users: responseUsers}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)

}

func createUserHandler(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	var err error
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"msg": "error. user not created.",
		})
		return
	}

	u := user.NewUser(body.Name, body.Email)
	err = user_repository.Store(u)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"msg": "user created.",
	})
}
