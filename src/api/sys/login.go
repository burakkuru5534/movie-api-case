package sys

import (
	"database/sql"
	"encoding/json"
	"github.com/burakkuru5534/src/auth"
	"github.com/burakkuru5534/src/helper"
	"net/http"

	"log"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var (
		upassFromDb *string

		userID       int64
		userName     string
		userFullName string
		userEmail    string
	)

	var loginInfo = &struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	err := helper.BodyToJsonReq(r, &loginInfo)
	if err != nil {
		log.Println("Login body to json error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	qs := "select id, coalesce(code, '') code, coalesce(full_name, '') full_name, coalesce(email, '') email, upass from sysusr where code = $1 or email = $1 and is_active"
	err = helper.App.DB.QueryRowx(qs, loginInfo.Username).Scan(&userID, &userName, &userFullName, &userEmail, &upassFromDb)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Login user not found: ", err)
			http.Error(w, "{\"error\": \"Bad request\"}", http.StatusBadRequest)
		} else {
			log.Println("Login user query error: ", err)
			http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		}
		return
	}

	// compare password hashes
	loginResult := helper.CheckPass(*upassFromDb, loginInfo.Password)
	if !loginResult {
		log.Println("Login password not match: ", err)
		http.Error(w, "{\"error\": \"Bad request\"}", http.StatusBadRequest)
		return
	}

	tc, err := auth.NewTokenClaimsForUser(userID, userName, userEmail, userFullName)
	if err != nil {
		log.Println("Login token claims error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	var respStruct struct {
		UserID       int64
		UserName     string
		UserEmail    string
		UserFullName string
		AccessToken  string
	}
	respStruct.UserID = userID
	respStruct.UserName = userName
	respStruct.UserEmail = userEmail
	respStruct.UserFullName = userFullName
	respStruct.AccessToken = tc.Encode(helper.Conf.Auth.JWTAuth)

	json.NewEncoder(w).Encode(respStruct)

}
