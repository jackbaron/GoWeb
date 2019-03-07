package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"projects/blog/helpers"

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

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		sess := helpers.Instance(r)
		if sess.Values["id"] == nil {
			http.Redirect(w, r, "/login", http.StatusFound)
		}
		// log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
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
	r.HandleFunc("/logout", controllers.LogOut).Methods("GET")

	// Admin
	adminPrefix := r.PathPrefix("/admin").Subrouter()
	adminPrefix.HandleFunc("/", controllers.AdminHome).Methods("GET")      // /admin
	adminPrefix.HandleFunc("/about", controllers.AdminHome).Methods("GET") // /admin/about
	// Check authentication login page admin
	adminPrefix.Use(loggingMiddleware)

	fmt.Printf("Server up and running . Running PORT: %s\n", port)

	fs := http.FileServer(http.Dir("./assets"))
	r.PathPrefix("/js/").Handler(fs)
	r.PathPrefix("/css/").Handler(fs)
	r.PathPrefix("/img/").Handler(fs)
	r.PathPrefix("/fonts/").Handler(fs)

	err := http.ListenAndServe(port, csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false))(r))

	if err != nil {
		log.Fatal("Error listening and server", err)
		return false
	}
	return true
}
