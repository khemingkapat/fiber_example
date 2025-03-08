package queries

import (
	"fmt"
	"log"

	object "github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

func CreatePerson(db *gorm.DB, person *object.Person) {
	result := db.Create(person)

	if err := result.Error; err != nil {
		log.Fatalf("Error Creating Person : %v", err)
	}

	fmt.Println("Person Created")
}
