package auth

import (
	object "github.com/khemingkapat/fiber_example/objects"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecretKey string = "mysecretkey"

func CreateUser(db *gorm.DB, user *object.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	result := db.Create(&user)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
