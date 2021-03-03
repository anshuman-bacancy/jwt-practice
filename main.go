package main

import (
	. "jwt/middlewares"
	"log"
	"net/http"
	"os"

	"jwt/controller"

	"github.com/gorilla/mux"
)

// func init() {
// 	// OpenDb to make migrations
// 	_ = database.OpenDb()
// }

func main() {
	router := mux.NewRouter()

	var port string
	if len(os.Args) >= 1 {
		port = string(":" + os.Args[1])
	} else {
		port = ":8000" // default port
	}

	// handing routes
	router.HandleFunc("/", controller.HomeHandler).Methods("GET")
	router.HandleFunc("/signup", controller.SignUpHandler).Methods("POST")
	router.HandleFunc("/signin", controller.SignInHandler).Methods("POST")
	router.HandleFunc("/admin", AuthorizeAdmin(controller.AdminHandler)).Methods("GET")

	log.Println("Server started at", port)
	http.ListenAndServe(port, router)
}
