package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"type:uuid;primaryKey" json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `gorm:"unique" json:"email"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}

	if u.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashed)
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		if len(u.Password) < 60 {
			hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
			if err != nil {
				return err
			}
			u.Password = string(hashed)
		}
	}
	return nil
}
