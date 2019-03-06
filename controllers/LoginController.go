package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"projects/blog/dataservice"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
)

// store will hold all session data
var store *sessions.CookieStore

// LoginGet Get Login Form
func LoginGet(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/Admin/Login.html")
	t.Execute(w, map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	})
}

// LoginPost Post Login Form
func LoginPost(w http.ResponseWriter, r *http.Request) {
	session, errorSession := store.Get(r, "cockie-name")
	if errorSession != nil {
		fmt.Println("StatusInternalServerError")
		http.Redirect(w, r, "/admin", http.StatusInternalServerError)
		return
	}
	repo := dataservice.NewUserRepo()
	r.ParseForm()
	userName := r.FormValue("UserName")
	passWord := r.FormValue("PassWord")

	// get user
	user := repo.GetUser(userName)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.PassWord, []byte(passWord))
		if err == nil {
			session.Values["username"] = userName
			err = session.Save(r, w)
			if err != nil {
				fmt.Println("Error save session user")
				http.Redirect(w, r, "/admin", http.StatusInternalServerError)
			}
			http.Redirect(w, r, "/admin", http.StatusFound)
		} else {
			// password is not coreect
			LoginGet(w, r)
		}
	} else {
		// username is not correct
		LoginGet(w, r)
	}
}

// func RegisterPost(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	var user models.User
// 	user.UserName = r.FormValue("UserName")
// 	passWord := r.FormValue("PassWord")
// 	user.PassWord, _ = bcrypt.GenerateFromPassword([]byte(passWord), bcrypt.DefaultCost)
// 	repo := dataservice.NewUserRepo()
// 	err := repo.Login(&user)
// 	fmt.Println(err)
// 	if err == 1 {
// 		fmt.Println("Username is exists")
// 	}

// 	LoginGet(w, r)
// }
