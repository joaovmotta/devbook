package controllers

import (
	"devbook-api/src/database"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	requestBody, err := io.ReadAll(r.Body)

	if err != nil {

		log.Fatal(err)
	}

	var user models.User

	if err := json.Unmarshal(requestBody, &user); err != nil {

		log.Fatal(err)
	}

	db, err := database.Connect()

	if err != nil {

		log.Fatal(err)
	}

	userRepository := repositories.NewUserRepository(db)

	userId, err := userRepository.Create(user)

	if err != nil {

		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("User created with id: %d", userId)))
}

func FindUsers(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Get all users"))
}

func FindUserById(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Get user by id"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Delete user"))
}
