package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"projects/blog/controllers"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

/*
* get port running app
 */
func getPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return ":3500", fmt.Errorf("$PORT not set")
	}

	return ":" + port, nil
}

/*
* get router running
 */
func GetRouter() bool {
	r := mux.NewRouter()

	port, _ := getPort()
	r.HandleFunc("/", controllers.Home).Methods("GET")

	r.HandleFunc("/login", controllers.LoginGet).Methods("GET")
	r.HandleFunc("/login", controllers.LoginPost).Methods("POST")

	fmt.Printf("Server up and running . Running PORT: %s\n", port)

	fs := http.FileServer(http.Dir("./assets"))
	r.PathPrefix("/js/").Handler(fs)
	r.PathPrefix("/css/").Handler(fs)
	r.PathPrefix("/img/").Handler(fs)
	r.PathPrefix("/fonts/").Handler(fs)

	err := http.ListenAndServe(port, csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(true))(r))

	if err != nil {
		log.Fatal("Error listening and server", err)
		return false
	}
	return true
}
