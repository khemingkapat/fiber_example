package object

import (
	"gorm.io/gorm"
)

// Room Model
type Room struct {
	gorm.Model
	BuildingID uint     `gorm:"not null" json:"building_id"`
	RoomNumber string   `gorm:"not null" json:"room_number"`
	Size       float64  `gorm:"not null" json:"size"`
	Type       string   `gorm:"not null" json:"type"` // "single" or "double"
	TenantID   *uint    `gorm:"unique" json:"tenant_id,omitempty"`
	Building   Building `gorm:"foreignKey:BuildingID" json:"building,omitempty"`
	Tenant     *User    `gorm:"foreignKey:TenantID" json:"tenant,omitempty"`
}

type RoomTenant struct {
	ID          uint    `json:"id"`
	RoomNumber  string  `json:"room_number"`
	Size        float64 `json:"size"`
	Type        string  `json:"type"`
	TenantName  string  `json:"tenant_name"`  // Tenant's name
	TenantEmail string  `json:"tenant_email"` // Tenant's email
}
