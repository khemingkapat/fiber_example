package queries

import (
	"log"

	"gorm.io/gorm"

	"github.com/khemingkapat/fiber_example/objects"
)

func GetPeople(db *gorm.DB) []object.Person {
	var people []object.Person

	result := db.Find(&people)

	if err := result.Error; err != nil {
		log.Printf("Error Getting People : %v", err)
		return nil
	}

	return people
}
