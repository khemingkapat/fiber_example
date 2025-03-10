package auth

import (
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
		return "", err
	}

	// Compare the hashed password from the database with the provided password
	err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err // Return error if password is incorrect
	}

	// Generate JWT token after successful login
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = selectedUser.ID // Use selectedUser.ID here
	claims["role"] = selectedUser.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}
