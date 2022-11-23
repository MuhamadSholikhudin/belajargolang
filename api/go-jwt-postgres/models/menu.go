package models

import "time"

type Menu struct {
	Id         int64     `gorm:"primaryKey" json:"id"`
	Menu       string    `gorm:"varchar(200)" json:"menu"`
	Created_at time.Time `gorm:"timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"timestamp" json:"updated_at"`
}
