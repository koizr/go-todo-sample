package persistent

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
	TimeZone string
}

func SetUpDB(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s TimeZone=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := migrate(db); err != nil {
		return db, err
	}

	return db, nil
}

func migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&User{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&Task{}); err != nil {
		return err
	}
	return nil
}
