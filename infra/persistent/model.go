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

func NewUser(user *ProvisionalUser) (*User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       id.String(),
		LoginID:  user.LoginID,
		Password: string(hash),
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
