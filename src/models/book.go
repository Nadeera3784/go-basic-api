package models

import (
    "gorm.io/gorm"
)

type Book struct {
    gorm.Model
    Title   string `gorm:"type:varchar(100); NOT NULL" validate:"required"`
	Description string `gorm:"type:text;NOT NULL" validate:"required"`
}
