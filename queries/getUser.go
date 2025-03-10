package queries

import (
	"github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB, id uint) (*object.User, error) {
	var person object.User
	result := db.First(&person, id)

	if err := result.Error; err != nil {
		return nil, err
	}

	return &person, nil
}
