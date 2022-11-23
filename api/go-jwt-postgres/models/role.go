package models

import "time"

type Role struct {
	Id         int64     `gorm:"primaryKey" json:"id"`
	Role       string    `gorm:"varchar(200)" json:"role"`
	Created_at time.Time `gorm:"timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"timestamp" json:"updated_at"`
}
