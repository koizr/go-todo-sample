package persistent

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID       string `gorm:"type:varchar(255);primaryKey"`
	LoginID  string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Name     string `gorm:"type:varchar(255);not null"`
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
