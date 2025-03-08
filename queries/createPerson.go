package queries

import (
	"github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

func CreatePerson(db *gorm.DB, person *object.Person) error {
	result := db.Create(person)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}
