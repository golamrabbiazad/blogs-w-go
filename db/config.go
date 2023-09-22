package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Blog struct {
	ID      int    `json:"id"`
	Author  Author `gorm:"embedded" json:"author"`
	Upvotes int32  `json:"up_votes"`
}

func Gorm() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Dhaka"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Blog{})

	return db
}
