package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lucchesisp/go-dev-book/src/authentication"
	"github.com/lucchesisp/go-dev-book/src/database"
	"github.com/lucchesisp/go-dev-book/src/models"
	"github.com/lucchesisp/go-dev-book/src/repositories"
	"github.com/lucchesisp/go-dev-book/src/responses"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserID(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var publication models.Publications

	if err = json.Unmarshal(bodyRequest, &publication); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = publication.Prepare(); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.GetConnection()
	defer db.Close()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	publicationRepository := repositories.NewPublicationsRepository(db)
	publication, err = publicationRepository.Create(publication, userID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publication)
}

func GetPublication(w http.ResponseWriter, r *http.Request) {
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

	publicationRepository := repositories.NewPublicationsRepository(db)
	publication, err := publicationRepository.FindByID(userID)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, publication)
}

func GetPublications(w http.ResponseWriter, r *http.Request) {
	tokenUserID, err := authentication.ExtractUserID(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.GetConnection()
	defer db.Close()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	publicationRepository := repositories.NewPublicationsRepository(db)
	publications, err := publicationRepository.FindAll(tokenUserID)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publications)
}

func UpdatePublication(w http.ResponseWriter, r *http.Request) {}

func DeletePublication(w http.ResponseWriter, r *http.Request) {}
