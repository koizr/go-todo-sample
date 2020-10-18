package persistent

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID       string `gorm:"primaryKey"`
	LoginID  string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
}

type ProvisionalUser struct {
	LoginID  string `json:"loginId"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func HashPassword(rawPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func NewUser(user *ProvisionalUser) (*User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	password, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       id.String(),
		LoginID:  user.LoginID,
		Password: password,
		Name:     user.Name,
	}, nil
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
