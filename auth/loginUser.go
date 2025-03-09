package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginUser(db *gorm.DB, user *User) (string, error) { // as returning token as string
	selectedUser := new(User)
	result := db.Where("email = ?", user.Email).First(selectedUser)

	if err := result.Error; err != nil {
		log.Println("Couldn't parse user")
		return "", err
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(selectedUser.Password),
		[]byte(user.Password),
	)
	if err != nil {
		log.Println("Wrong Password")
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		log.Println("Error create token")
		return "", err
	}

	return t, nil
}
