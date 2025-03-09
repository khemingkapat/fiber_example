package auth

import (
	"gorm.io/gorm"
)

var jwtSecretKey string = "mysecretkey"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}
