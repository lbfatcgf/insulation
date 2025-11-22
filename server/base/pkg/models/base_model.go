package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string         `gorm:"column:id;primaryKey;default:uuidv7();type:uuid" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp(0);not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp(0);not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;type:timestamp(0)"`
}
