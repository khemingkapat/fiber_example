package queries

import (
	"github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

func DeleteUser(db *gorm.DB, id uint) error {
	result := db.Delete(&object.User{}, id)

	if err := result.Error; err != nil {
		return err
	}

	return nil
}
