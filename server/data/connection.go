package data

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error) {
	path := os.Getenv("JADE_DB_PATH")
	fmt.Println(path)

	db, err := gorm.Open(sqlite.Open(path))
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
