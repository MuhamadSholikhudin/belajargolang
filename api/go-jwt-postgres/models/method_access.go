package models

import "time"

type Method_access struct {
	Id             int64     `gorm:"primaryKey" json:"id"`
	Access_Menu_id int       `gorm:"int(11)" json:"access_menu_id"`
	Sub_Menu_id    int       `gorm:"int(11)" json:"sub_menu_id"`
	View           string    `gorm:"varchar(10)" json:"view"`
	Edit           string    `gorm:"varchar(10)" json:"edit"`
	Delete         string    `gorm:"varchar(10)" json:"delete"`
	Created_at     time.Time `gorm:"timestamp" json:"created_at"`
	Updated_at     time.Time `gorm:"timestamp" json:"updated_at"`
}
