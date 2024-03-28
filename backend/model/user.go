package model

import "time"

type User struct {
	ID		uint      `gorm:"primaryKey" json:"id" param:"id"`
	Name   string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}