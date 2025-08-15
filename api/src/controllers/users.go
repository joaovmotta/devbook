package controllers

import (
	"devbook-api/src/database"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	requestBody, err := io.ReadAll(r.Body)

	if err != nil {

		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err := json.Unmarshal(requestBody, &user); err != nil {

		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare(); err != nil {

		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {

		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.NewUserRepository(db)

	user.ID, err = userRepository.Create(user)

	defer db.Close()

	if err != nil {

		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func FindUsers(w http.ResponseWriter, r *http.Request) {

	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()

	if err != nil {

		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUserRepository(db)

	users, err := repository.FindUser(nameOrNick)

	if err != nil {

		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
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
