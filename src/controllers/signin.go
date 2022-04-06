package controllers

import (
	"encoding/json"
	"errors"
	"github.com/lucchesisp/go-dev-book/src/authentication"
	"github.com/lucchesisp/go-dev-book/src/database"
	"github.com/lucchesisp/go-dev-book/src/models"
	"github.com/lucchesisp/go-dev-book/src/repositories"
	"github.com/lucchesisp/go-dev-book/src/responses"
	"github.com/lucchesisp/go-dev-book/src/security"
	"io/ioutil"
	"net/http"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.GetConnection()
	defer db.Close()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.NewUsersRepository(db)
	userRegistred, err := userRepository.FindByEmail(user.Email)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.Verify(userRegistred.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("invalid email or password"))
		return
	}

	token, _ := authentication.CreateToken(userRegistred.ID)

	responses.JSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}
