package queries

import (
	"log"

	"github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

func GetPerson(db *gorm.DB, id uint) *object.Person {
	var person object.Person
	result := db.First(&person, id)

	if err := result.Error; err != nil {
		log.Fatalf("Error Creating Person : %v", err)
		return nil
	}

	return &person
}
