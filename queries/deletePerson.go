package queries

import (
	"github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

func DeletePerson(db *gorm.DB, id uint) error {
	result := db.Delete(&object.Person{}, id)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}
