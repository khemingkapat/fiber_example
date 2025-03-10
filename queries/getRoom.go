package queries

import (
	"github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

func GetRoom(db *gorm.DB, id uint) (*object.Room, error) {
	var room object.Room
	result := db.Preload("Building").First(&room, id)

	if err := result.Error; err != nil {
		return nil, err
	}

	return &room, nil
}
