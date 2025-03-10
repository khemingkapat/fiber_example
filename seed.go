package main

// import (
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"time"
//
// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )
//
// // Models
// type Building struct {
// 	gorm.Model
// 	Name     string `gorm:"unique;not null"`
// 	Location string `gorm:"not null"`
// 	Floors   int    `gorm:"not null"`
// 	Rooms    []Room `gorm:"foreignKey:BuildingID"`
// }
//
// type Room struct {
// 	gorm.Model
// 	BuildingID uint    `gorm:"not null"`
// 	RoomNumber string  `gorm:"not null"`
// 	Size       float64 `gorm:"not null"`
// 	Type       string  `gorm:"not null"` // "single" or "double"
// 	TenantID   *uint   `gorm:"unique"`
// 	Building   Building
// 	Tenant     *User `gorm:"foreignKey:TenantID"`
// }
//
// type User struct {
// 	gorm.Model
// 	Name     string `gorm:"not null"`
// 	Email    string `gorm:"unique;not null"`
// 	Password string `gorm:"not null"`
// 	Role     string `gorm:"not null"` // "tenant" or "manager"
// 	Room     *Room  `gorm:"foreignKey:TenantID"`
// }
//
// func main() {
// 	// Replace with your actual PostgreSQL connection string
// 	dsn := "user=admin password=admin dbname=dorm host=localhost port=5430 sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to database:", err)
// 	}
//
// 	// Auto Migrate
// 	db.AutoMigrate(&Building{}, &Room{}, &User{})
//
// 	// Generate Example Data
// 	createExampleData(db)
//
// 	fmt.Println("Example data successfully generated!")
// }
//
// // Helper function to hash passwords
// func hashPassword(password string) string {
// 	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	return string(hashed)
// }
//
// // Generate Example Data
// func createExampleData(db *gorm.DB) {
// 	rand.Seed(time.Now().UnixNano())
//
// 	// Create Buildings
// 	buildings := []Building{
// 		{Name: "Dormitory A", Location: "Main Street 1", Floors: 5},
// 		{Name: "Dormitory B", Location: "Main Street 2", Floors: 4},
// 		{Name: "Dormitory C", Location: "Main Street 3", Floors: 6},
// 	}
//
// 	db.Create(&buildings)
//
// 	// Create Rooms (Randomly 15-20)
// 	var rooms []Room
// 	roomCount := rand.Intn(6) + 15 // 15 to 20 rooms
//
// 	for i := 1; i <= roomCount; i++ {
// 		building := buildings[rand.Intn(len(buildings))] // Random building
// 		roomType := "single"
// 		if rand.Intn(2) == 1 {
// 			roomType = "double"
// 		}
// 		room := Room{
// 			BuildingID: building.ID,
// 			RoomNumber: fmt.Sprintf("Room-%d", i),
// 			Size:       float64(rand.Intn(20) + 10), // 10-30 sqm
// 			Type:       roomType,
// 		}
// 		rooms = append(rooms, room)
// 	}
//
// 	db.Create(&rooms)
//
// 	// Create Users (Randomly 15-20)
// 	var users []User
// 	userCount := rand.Intn(6) + 15 // 15 to 20 users
// 	tenantCount := 0
//
// 	for i := 1; i <= userCount; i++ {
// 		isManager := i%5 == 0 // Every 5th user is a manager
// 		role := "tenant"
// 		if isManager {
// 			role = "manager"
// 		}
//
// 		password := fmt.Sprintf("user%dpassword", i) // Password format user<ID>password
//
// 		user := User{
// 			Name:     fmt.Sprintf("User %d", i),
// 			Email:    fmt.Sprintf("user%d@example.com", i),
// 			Password: hashPassword(password),
// 			Role:     role,
// 		}
//
// 		if !isManager && tenantCount < len(rooms) { // Assign tenant to a room
// 			user.Room = &rooms[tenantCount]
// 			tenantCount++
// 		}
//
// 		users = append(users, user)
// 	}
//
// 	db.Create(&users)
// }
