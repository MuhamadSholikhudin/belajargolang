package models

import "time"

type Sub_Menu struct {
	Id         int64     `gorm:"primaryKey" json:"id"`
	Menu_id    int       `gorm:"int(11)" json:"menu_id"`
	Title      string    `gorm:"varchar(100)" json:"title"`
	Url        string    `gorm:"varchar(225)" json:"url"`
	Icon       string    `gorm:"varchar(100)" json:"icon"`
	Created_at time.Time `gorm:"timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"timestamp" json:"updated_at"`
}
