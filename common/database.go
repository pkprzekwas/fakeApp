package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkprzekwas/fakeApp/config"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func buildConnString(config *config.DBConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		config.Host,
		config.Username,
		config.DBName,
		config.Password,
	)
}

func Init(config *config.DBConfig) *gorm.DB {
	connStr := buildConnString(config)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic("Couldn't open DB")
	}

	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
