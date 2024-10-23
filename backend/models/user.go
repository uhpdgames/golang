package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	//gorm.Model // Atto added fields: ID, CreatedAt, UpdatedAt, DeletedAt
    ID        uint           `json:"id" gorm:"primaryKey"`

	Name     string `json:"username" gorm:"type:varchar(255);not null"`
	Email    string `json:"email"  gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type RegisterInput struct {
	Name     string `json:"name" binding:"required" validate:"min=2,max=100"`
	Email    string `json:"email" binding:"required,email" validate:"email"`
	Password string `json:"password" binding:"required" validate:"min=6"`
}
