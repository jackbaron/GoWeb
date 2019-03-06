package dataservice

import (
	"fmt"
	"log"
	"projects/blog/helpers"
	"projects/blog/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

var hmacSecret = []byte{97, 48, 97, 50, 97, 98, 105, 49, 99, 102, 83, 53, 57, 98, 52, 54, 97, 102, 99, 12, 12, 13, 56, 34, 23, 16, 78, 67, 54, 34, 32, 21}

// UserRepo struct
type UserRepo struct {
	db *gorm.DB
}

// NewUserRepo is create Connect DB
func NewUserRepo() *UserRepo {
	return &UserRepo{db: GetDBConection()}
}

// Register User
func (repo *UserRepo) Register(obj *models.User) int {
	// Check username unique
	user := repo.db.Where("user_name = ? ", obj.UserName).Find(&obj)
	if !user.RecordNotFound() {
		// Exists username
		helpers.WiteLog("debug", "Exists username")
		return -1
	}
	// check email unique
	// user = repo.db.Where("email = ?", obj.Email, &obj)
	// if !user.RecordNotFound() {
	// 	// Exists Email
	// 	return -2
	// }

	tx := repo.db.Begin()
	obj.Token = []byte(encodeToken(obj.UserName))
	err := tx.Create(obj).Error

	if err != nil {
		log.Fatal(err)
		helpers.WiteLog("debug", "Error insert user")
	}

	tx.Commit()

	return 1
}

// GetUser get user
func (repo *UserRepo) GetUser(username string) (user *models.User) {
	user = &models.User{}
	err := repo.db.Where("user_name = ?", username).Find(user)
	if err.RecordNotFound() {
		log.Println("Username not correct")
		return nil
	}
	return
}

func encodeToken(email string) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"nbf":   time.Now(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecret)

	fmt.Println(tokenString, err)

	return tokenString
}
