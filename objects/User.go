package object

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string
	Email        string
	PasswordHash string
	Role         string
	PersonID     uint   // Foreign key to People table
	Person       Person `gorm:"foreignKey:PersonID"` // One-to-one relationship
}
