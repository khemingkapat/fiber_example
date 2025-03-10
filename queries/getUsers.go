package queries

import (
	"log"

	"gorm.io/gorm"

	"github.com/khemingkapat/fiber_example/objects"
)

func GetUsers(db *gorm.DB) []object.User {
	var people []object.User

	result := db.Find(&people)

	if err := result.Error; err != nil {
		log.Printf("Error Getting People : %v", err)
		return nil
	}

	return people
}
