package queries

import (
	"gorm.io/gorm"

	"github.com/khemingkapat/fiber_example/objects"
)

func GetRoomsWithTenant(db *gorm.DB) []object.RoomTenant {
	var rooms []object.RoomTenant

	result := db.Preload("Tenant").Table("rooms").Joins("JOIN users AS tenant ON tenant.id = rooms.tenant_id").
		Select("rooms.id, rooms.room_number, rooms.size, rooms.type, tenant.name AS tenant_name, tenant.email AS tenant_email").
		Scan(&rooms)

	if err := result.Error; err != nil {
		return nil
	}

	return rooms
}
