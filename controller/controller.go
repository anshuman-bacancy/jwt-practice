package controller

import (
	"encoding/json"
	"jwt/models"
	"jwt/services"
	"jwt/utils"
	"net/http"
)

// HomeHandler handles "/"
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hi, Welcome to home"))
}

// SignUpHandler handles "/signup"
func SignUpHandler(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	email := req.FormValue("email")
	password := req.FormValue("password")
	role := req.FormValue("role")

	services.SaveUser(name, email, password, role)
}

// SignInHandler handles "/signin"
func SignInHandler(res http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")

	match, user := services.SignInUser(email, password)
	if match {
		token, _ := utils.GenerateJWTToken(user.Email, user.Role)
		userResp := struct {
			User  models.User
			Token string
		}{user, token}

		res.Header().Set("Content-Type", "appliction/json")
		json.NewEncoder(res).Encode(userResp)
	} else {
		res.Write([]byte("Not Valid user"))
	}
}

// AdminHandler handles admin routes
func AdminHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("welcome, admin"))
}
