package queries

import (
	"log"

	"gorm.io/gorm"

	"github.com/khemingkapat/fiber_example/objects"
)

func GetRooms(db *gorm.DB) []object.Room {
	var rooms []object.Room

	result := db.Preload("Building").Find(&rooms)

	if err := result.Error; err != nil {
		log.Printf("Error Getting Room : %v", err)
		return nil
	}

	return rooms
}
