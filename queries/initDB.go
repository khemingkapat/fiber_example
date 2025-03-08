package queries

import (
	"log"
	"os"
	"time"

	object "github.com/khemingkapat/fiber_example/objects"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB(connStr string) (*gorm.DB, error) { // Corrected function signature
	// Open a connection to the database
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: newLogger, // add Logger
	})
	if err != nil {
		panic("DB Initialize Fail")
	}

	db.AutoMigrate(&object.Person{})

	return db, nil
}
