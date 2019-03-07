package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"projects/blog/dataservice"
	"projects/blog/helpers"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
)

// store will hold all session data
var store *sessions.CookieStore

// LoginGet Get Login Form
func LoginGet(w http.ResponseWriter, r *http.Request) {
	//  Check is exists session login
	sess := helpers.Instance(r)
	if sess.Values["id"] != nil {
		http.Redirect(w, r, "/admin", http.StatusFound)
		return
	}

	v := map[string]interface{}{
		"FlashedMessages": sess.Flashes(),
		csrf.TemplateTag:  csrf.TemplateField(r),
	}
	t, _ := template.ParseFiles("views/Admin/Login.html")
	t.Execute(w, v)
}

const (
	// Name of the session variable that tracks login attempts
	sessLoginAttempt = "login_attempt"
)

// loginAttempt increments the number of login attempts in sessions variable
func loginAttempt(sess *sessions.Session) {
	// Log the attempt
	if sess.Values[sessLoginAttempt] == nil {
		sess.Values[sessLoginAttempt] = 1
	} else {
		sess.Values[sessLoginAttempt] = sess.Values[sessLoginAttempt].(int) + 1
	}
}

// LoginPost Post Login Form
func LoginPost(w http.ResponseWriter, r *http.Request) {
	//create new session
	sess := helpers.Instance(r)

	// Check user submit post deveice
	if sess.Values[sessLoginAttempt] != nil && sess.Values[sessLoginAttempt].(int) >= 5 {
		log.Println("Brute force login prevented")
		LoginGet(w, r)
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
			helpers.Empty(sess)
			sess.Values["id"] = user.ID
			sess.Values["name"] = user.Name
			sess.Save(r, w)
			http.Redirect(w, r, "/admin", http.StatusFound)
		} else {
			sess.AddFlash(helpers.Flash{Message: "Username and password is not correct", Class: helpers.FlashError})
			sess.Save(r, w)
			// password is not coreect
		}
	} else {
		// username is not correct
		sess.AddFlash(helpers.Flash{Message: "Username and password is not correct", Class: helpers.FlashError})
		sess.Save(r, w)
	}
	LoginGet(w, r)
}

// LogOut user
func LogOut(w http.ResponseWriter, r *http.Request) {
	// get session
	sess := helpers.Instance(r)
	if sess.Values["id"] != nil {
		helpers.Empty(sess)
		fmt.Println(sess)
		sess.AddFlash(helpers.Flash{Message: "Goodbye !", Class: helpers.FlashNotice})
		sess.Save(r, w)
		fmt.Println("User Logout")
		go LoginGet(w, r)
	}
	go AdminHome(w, r)
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
