package persistent

import "time"

type User struct {
	ID       string `gorm:"primaryKey"`
	LoginID  string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
}

type Task struct {
	ID          string `gorm:"primaryKey"`
	UserID      string `gorm:"not null"`
	User        User
	Subject     string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Status      uint      `gorm:"not null"`
	DueDate     time.Time `gorm:"not null"`
}
