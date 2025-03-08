package queries

import (
	"fmt"
	"log"

	"github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

func DeletePerson(db *gorm.DB, id uint) {
	result := db.Delete(&object.Person{}, id)

	if err := result.Error; err != nil {
		log.Fatalf("Error Deleting Person : %v", err)
	}

	fmt.Println("Person Deleted")
}
