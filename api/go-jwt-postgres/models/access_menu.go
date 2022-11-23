package models

import "time"

type Access_menu struct {
	Id         int64     `gorm:"primaryKey" json:"id"`
	Role_id    int       `gorm:"int(11)" json:"role_id"`
	Menu_id    int       `gorm:"int(11)" json:"menu_id"`
	Created_at time.Time `gorm:"timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"timestamp" json:"updated_at"`
}
