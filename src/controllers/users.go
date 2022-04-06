package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lucchesisp/go-dev-book/src/authentication"
	"github.com/lucchesisp/go-dev-book/src/database"
	"github.com/lucchesisp/go-dev-book/src/models"
	"github.com/lucchesisp/go-dev-book/src/repositories"
	"github.com/lucchesisp/go-dev-book/src/responses"
	"github.com/lucchesisp/go-dev-book/src/security"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNickname := strings.ToLower(r.URL.Query().Get("q"))

	db, err := database.GetConnection()
	defer db.Close()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.NewUsersRepository(db)
	users, err := userRepository.FindByNameOrNickname(nameOrNickname)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)

	if err != nil {
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
	user, err := userRepository.FindByID(userID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	// TODO: transformar 'register' em um ENUM

	if err = user.Prepare("register"); err != nil {
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
	user, err = userRepository.Create(user)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := authentication.ExtractUserID(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != tokenUserID {
		responses.Error(w, http.StatusForbidden, fmt.Errorf("user id does not match token user id"))
		return
	}

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

	if err = user.Prepare("update"); err != nil {
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

	user, err = userRepository.Update(userID, user)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := authentication.ExtractUserID(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != tokenUserID {
		responses.Error(w, http.StatusForbidden, fmt.Errorf("user id does not match token user id"))
		return
	}

	db, err := database.GetConnection()
	defer db.Close()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.NewUsersRepository(db)

	err = userRepository.Delete(userID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, map[string]string{"message": "user deleted"})
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := authentication.ExtractUserID(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != tokenUserID {
		responses.Error(w, http.StatusForbidden, fmt.Errorf("user id does not match token user id"))
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var password models.Password

	if err = json.Unmarshal(bodyRequest, &password); err != nil {
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

	savedPassword, err := userRepository.FindPassword(userID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.Verify(savedPassword, password.CurrentPassword); err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("current password is incorrect"))
		return
	}

	hashedPassword, err := security.Hash(password.NewPassword)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = userRepository.UpdatePassword(userID, string(hashedPassword)); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, map[string]string{"message": "password updated"})
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := authentication.ExtractUserID(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID == tokenUserID {
		responses.Error(w, http.StatusForbidden, fmt.Errorf("not allowed to follow yourself"))
		return
	}

	db, err := database.GetConnection()
	defer db.Close()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.NewUsersRepository(db)

	err = userRepository.Follow(tokenUserID, userID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, map[string]string{"message": "user followed"})
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	tokenUserID, err := authentication.ExtractUserID(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID == tokenUserID {
		responses.Error(w, http.StatusForbidden, fmt.Errorf("not allowed to unfollow yourself"))
		return
	}

	db, err := database.GetConnection()
	defer db.Close()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userRepository := repositories.NewUsersRepository(db)

	err = userRepository.Unfollow(tokenUserID, userID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, map[string]string{"message": "user unfollowed"})
}

func FindFollowers(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)

	if err != nil {
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
	followers, err := userRepository.FindFollowers(userID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

func FindFollowing(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)

	if err != nil {
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
	following, err := userRepository.FindFollowing(userID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, following)
}
