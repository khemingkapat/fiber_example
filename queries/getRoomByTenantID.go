package queries

import (
	"log"

	"github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

func GetRoomByTenantID(db *gorm.DB, id uint) *object.Room {
	var room object.Room

	result := db.Preload("Building").Where("tenant_id = ?", id).First(&room)

	if err := result.Error; err != nil {
		return nil
	}
	log.Println(room)

	return &room
}
