package persistent

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID       string `gorm:"type:varchar(255);primaryKey"`
	LoginID  string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Name     string `gorm:"type:varchar(255);not null"`
}

type Task struct {
	ID          string `gorm:"type:varchar(255);primaryKey"`
	UserID      string `gorm:"type:varchar(255);not null"`
	User        User
	Subject     string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text;not null"`
	Status      uint      `gorm:"type:smallint;not null"`
	DueDate     time.Time `gorm:"not null"`
}

func GenerateTaskID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), err
}
