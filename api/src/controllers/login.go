package controllers

import (
	"devbook-api/src/authorization"
	"devbook-api/src/database"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"
	"devbook-api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	request, err := ioutil.ReadAll(r.Body)

	if err != nil {

		responses.Error(w, http.StatusUnprocessableEntity, err)

		return
	}

	var user models.User

	if err := json.Unmarshal(request, &user); err != nil {

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

	userDB, err := repository.FindUserByEmail(user.Email)

	if err != nil {

		responses.Error(w, http.StatusBadRequest, err)
	}

	if err := security.CheckPassword(user.Password, userDB.Password); err != nil {

		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authorization.CreateToken(userDB.ID)

	if err != nil {

		responses.Error(w, http.StatusInternalServerError, err)
	}

	w.Write([]byte(token))
}
