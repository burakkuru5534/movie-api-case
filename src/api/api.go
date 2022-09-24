package api

import (
	"encoding/json"
	"github.com/Shyp/go-dberror"
	"github.com/burakkuru5534/src/helper"
	"github.com/burakkuru5534/src/model"
	_ "github.com/letsencrypt/boulder/db"
	"net/http"

	"log"
)

func MovieCreate(w http.ResponseWriter, r *http.Request) {

	var Movie model.Movie

	err := helper.BodyToJsonReq(r, &Movie)
	if err != nil {
		log.Println("movie create body to json error: ", err)
		http.Error(w, "{\"error\": \"Bad request\"}", http.StatusBadRequest)
		return
	}

	err = Movie.Create()
	if err != nil {
		dberr := dberror.GetError(err)
		switch e := dberr.(type) {
		case *dberror.Error:
			if e.Code == "23505" {
				log.Println("movie already exist error: ", err)
				http.Error(w, "{\"error\": \"Movie with that name already exists\"}", http.StatusForbidden)
				return
			}
		}
		log.Println("movie create error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	respBody := struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Typ         string `json:"typ"`
	}{
		ID:          Movie.ID,
		Name:        Movie.Name,
		Description: Movie.Description,
		Typ:         Movie.Typ,
	}

	json.NewEncoder(w).Encode(respBody)

}

func MovieUpdate(w http.ResponseWriter, r *http.Request) {

	var Movie model.Movie

	//id := helper.StrToInt64(chi.URLParam(r, "id"))
	id := helper.StrToInt64(r.URL.Query().Get("id"))

	movieName, err := helper.GetMovieName(id)
	if err != nil {
		log.Println("movie update get movie name error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}
	isExists, err := helper.CheckIfMovieExists(movieName)
	if err != nil {
		log.Println("movie update check if movie exists error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	if !isExists {
		log.Println("movie update movie does not exist error: ", err)
		http.Error(w, "{\"error\": \"Movie with that name does not exist\"}", http.StatusNotFound)
		return
	}

	err = helper.BodyToJsonReq(r, &Movie)
	if err != nil {
		log.Println("movie update body to json error: ", err)
		http.Error(w, "{\"error\": \"Bad request\"}", http.StatusBadRequest)
		return
	}

	err = Movie.Update(id)
	if err != nil {
		dberr := dberror.GetError(err)
		switch e := dberr.(type) {
		case *dberror.Error:
			if e.Code == "23505" {
				log.Println("movie update already exist error: ", err)
				http.Error(w, "{\"error\": \"Movie with that name already exists\"}", http.StatusForbidden)
				return
			}
		}

		log.Println("movie update error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	respBody := struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Typ         string `json:"typ"`
	}{
		ID:          id,
		Name:        Movie.Name,
		Description: Movie.Description,
		Typ:         Movie.Typ,
	}
	json.NewEncoder(w).Encode(respBody)

}

func MovieDelete(w http.ResponseWriter, r *http.Request) {

	var Movie model.Movie

	//id := helper.StrToInt64(chi.URLParam(r, "id"))
	id := helper.StrToInt64(r.URL.Query().Get("id"))

	movieName, err := helper.GetMovieName(id)
	if err != nil {
		log.Println("movie delete get movie name error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}
	isExists, err := helper.CheckIfMovieExists(movieName)
	if err != nil {
		log.Println("movie delete check if movie exists error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	if !isExists {
		log.Println("movie delete movie does not exist error: ", err)
		http.Error(w, "{\"error\": \"Movie with that name does not exist\"}", http.StatusNotFound)
		return
	}

	err = Movie.Delete(id)
	if err != nil {
		log.Println("movie delete error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Movie deleted")

}

func MovieGet(w http.ResponseWriter, r *http.Request) {

	var Movie model.Movie

	id := helper.StrToInt64(r.URL.Query().Get("id"))
	//id := helper.StrToInt64(chi.URLParam(r, "id"))

	movieName, err := helper.GetMovieName(id)
	if err != nil {
		log.Println("movie get get movie name error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	isExists, err := helper.CheckIfMovieExists(movieName)
	if err != nil {
		log.Println("movie get check if movie exists error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	if !isExists {
		log.Println("movie get movie does not exist error: ", err)
		http.Error(w, "{\"error\": \"Movie with that name does not exist\"}", http.StatusNotFound)
		return
	}

	err = Movie.Get(id)
	if err != nil {
		log.Println("movie get error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	respBody := struct {
		ID          int64  `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Typ         string `json:"typ"`
	}{
		ID:          id,
		Name:        Movie.Name,
		Description: Movie.Description,
		Typ:         Movie.Typ,
	}
	json.NewEncoder(w).Encode(respBody)

}

func MovieList(w http.ResponseWriter, r *http.Request) {

	var Movie model.Movie
	//for pagination
	limit := helper.StrToInt64(r.URL.Query().Get("limit"))
	offset := helper.StrToInt64(r.URL.Query().Get("offset"))

	//initial pagination values if user won't send it in request
	if limit == 0 {
		limit = 10
	}
	if offset == 0 {
		offset = 0
	}
	MovieList, err := Movie.GetAll(limit, offset)
	if err != nil {
		log.Println("movie list error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(MovieList)

}
