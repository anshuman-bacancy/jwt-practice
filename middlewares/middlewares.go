package middlewares

import (
	"fmt"
	"jwt/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

// AuthorizeAdmin authorizes admin
func AuthorizeAdmin(handler http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Header["token"])
		if req.Header["Token"] == nil {
			fmt.Println("here 1")
			http.Redirect(res, req, "/", 303)
			return
		}

		config, _ := utils.LoadConfig(".")
		configByte := []byte(config.SigningKey)

		token, err := jwt.Parse(req.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return configByte, nil
		})
		if err != nil {
			fmt.Fprintf(res, err.Error())
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "Admin" {
				fmt.Println("claimsss", claims)
				// http.Redirect(res, req, "/admin", 303)
				handler.ServeHTTP(res, req)
			}
		}
		// http.Redirect(res, req, "/", 303)
	}
}
