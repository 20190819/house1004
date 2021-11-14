package models

import (
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"primary key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt gorm.DeletedAt `gorm:"index"`
}
