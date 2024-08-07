package model

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := "host=postgres user=postgres password=postgres123 dbname=mydb port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("Success Started DB ")

	return nil
}

type Message struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"size:255"`
	CreatedAt time.Time
}

func Migrate() {
	DB.AutoMigrate(&Message{})
}

func SaveMessage(content string) error {
	message := Message{Content: content}
	return DB.Create(&message).Error
}
