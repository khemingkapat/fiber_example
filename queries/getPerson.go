package queries

import (
	"github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

func GetPerson(db *gorm.DB, id uint) (*object.Person, error) {
	var person object.Person
	result := db.First(&person, id)

	if err := result.Error; err != nil {
		return nil, err
	}

	return &person, nil
}
