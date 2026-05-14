package dao

import (
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey;column:id"`
	Name string `gorm:"not null;column:name"`
	Email string `gorm:"not null;column:email"`
	Password string `gorm:"not null;column:password"`
}

func (User) TableName() string {
	return "users"
}