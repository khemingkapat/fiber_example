package queries

import (
	"github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

func UpdatePerson(db *gorm.DB, person *object.Person) error {
	result := db.Save(&person)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}
