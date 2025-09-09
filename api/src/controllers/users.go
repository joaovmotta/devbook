package controllers

import (
	"devbook-api/src/authorization"
	"devbook-api/src/database"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if err := user.Prepare("creation"); err != nil {

		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {

		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)

	user.ID, err = userRepository.Create(user)

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

	parameters := mux.Vars(r)

	userId, err := strconv.ParseInt(parameters["userId"], 10, 64)

	if err != nil {

		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {

		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)

	user, err := repository.FindUserById(uint64(userId))

	if err != nil {

		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)

	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)

	if err != nil {

		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIdFromToken, err := authorization.GetUserId(r)

	if err != nil {

		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userId != userIdFromToken {

		responses.Error(w, http.StatusForbidden, errors.New("not authorized to do it"))
		return
	}

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

	if err := user.Prepare("update"); err != nil {

		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()

	if err != nil {

		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	userRepository.Update(uint64(userId), user)

	responses.JSON(w, http.StatusNoContent, nil)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)

	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)

	if err != nil {

		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIdFromToken, err := authorization.GetUserId(r)

	if err != nil {

		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userId != userIdFromToken {

		responses.Error(w, http.StatusForbidden, errors.New("not authorized to do it"))
		return
	}

	db, err := database.Connect()

	if err != nil {

		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	userRepository := repositories.NewUserRepository(db)
	userRepository.Delete(uint64(userId))

	responses.JSON(w, http.StatusNoContent, nil)
}
