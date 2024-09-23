package pkg

import (
	"fmt"
	"gitdeco-api/internal/models"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	username string
	password string
	host     string
	port     string
	name     string
}

func (db *Database) GetUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.username, db.password, db.host, db.port, db.name)
}

func NewDatabase() (*gorm.DB, error) {
	database := Database{
		username: os.Getenv("DATABASE_USERNAME"),
		password: os.Getenv("DATABASE_PASSWORD"),
		host:     os.Getenv("DATABASE_HOST"),
		port:     os.Getenv("DATABASE_PORT"),
		name:     os.Getenv("DATABASE_NAME"),
	}

	db, err := gorm.Open(mysql.Open(database.GetUrl()), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	if err != nil {
		return nil, err
	}

	DB := db
	err = db.AutoMigrate(new(models.User), new(models.Deco))
	if err != nil {
		return nil, err
	}

	return DB, nil
}
