package queries

import (
	"fmt"
	"log"

	"github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

func UpdatePerson(db *gorm.DB, person *object.Person) {
	result := db.Save(&person)

	if err := result.Error; err != nil {
	}

	if err := result.Error; err != nil {
		log.Fatalf("Error Updating Person : %v", err)
	}
	fmt.Println("Person Updated")
}
