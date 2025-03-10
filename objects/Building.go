package object

import (
	"gorm.io/gorm"
)

// Building Model
type Building struct {
	gorm.Model
	Name     string `gorm:"unique;not null" json:"name"`
	Location string `gorm:"not null" json:"location"`
	Floors   int    `gorm:"not null" json:"floors"`
	Rooms    []Room `gorm:"foreignKey:BuildingID" json:"rooms"`
}
