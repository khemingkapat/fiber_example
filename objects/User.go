package object

import (
	"gorm.io/gorm"
)

// User Model
type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"` // Hide password from JSON response
	Role     string `gorm:"not null" json:"role"`     // "tenant" or "manager"
	Room     *Room  `gorm:"foreignKey:TenantID" json:"room,omitempty"`
}

type UserLogin struct {
	gorm.Model
	Email    string
	Password string
}
