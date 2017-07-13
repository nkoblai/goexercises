package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func (a *App) UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	if user.Name == "" || user.Password == "" {
		http.Error(w, "Bad arguments", http.StatusBadRequest)
		return
	}
	var userid int
	err = a.DB.QueryRow(`INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`, user.Name, user.Email, user.Password).Scan(&userid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.ID = userid
	js, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(js)
}

func (a *App) ShowUsersHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := a.DB.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	js, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
