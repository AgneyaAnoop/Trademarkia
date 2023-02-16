package Storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Host, config.User, config.Password, config.DBName, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}