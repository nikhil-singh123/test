package models

import (
	"time"

	"gorm.io/gorm"
)

type Library struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (l *Library) CheckName() bool {
	var lb Library
	if err := DB.Where("Name = ?", l.Name).First(&lb).Error; err != nil {
		return false
	}
	return true

}

func (l *Library) SaveLibrary() (*Library, error) {
	if err := DB.Create(&l).Error; err != nil {
		return &Library{}, err
	}

	return l, nil

}
