package queries

import (
	"github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

func UpdateUser(db *gorm.DB, person *object.User) error {
	result := db.Save(&person)

	if err := result.Error; err != nil {
		return err
	}
	return nil
}
