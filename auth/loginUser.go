package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/khemingkapat/fiber_example/objects"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginUser(db *gorm.DB, user *object.UserLogin) (string, error) {
	// Find the user by email in the database
	selectedUser := new(object.User)
	result := db.Where("email = ?", user.Email).First(selectedUser)

	if err := result.Error; err != nil {
		log.Println("Couldn't find user with email:", user.Email)
		return "", err
	}

	// Compare the hashed password from the database with the provided password
	err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
	if err != nil {
		log.Println("Wrong Password for user:", user.Email)
		log.Println("Entered password:", user.Password)               // This is for debugging
		log.Println("Stored hashed password:", selectedUser.Password) // This is for debugging
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		log.Println("Entered Hashed password:", string(hashedPassword))

		return "", err // Return error if password is incorrect
	}

	// Generate JWT token after successful login
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = selectedUser.ID // Use selectedUser.ID here
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		log.Println("Error creating token")
		return "", err
	}

	return t, nil
}
