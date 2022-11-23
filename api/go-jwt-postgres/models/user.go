package models

import "time"

type User struct {
	Id           int64     `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"varchar(300)" json:"name"`
	Username     string    `gorm:"varchar(300)" json:"username"`
	Password     string    `gorm:"varchar(300)" json:"password"`
	Email        string    `gorm:"varchar(300)" json:"email"`
	Phone_number string    `gorm:"varchar(20)" json:"phone_number"`
	Role_id      int8      `gorm:"int(11)" json:"role_id"`
	Access_token string    `gorm:"varchar(300)" json:"access_token"`
	Is_active    int8      `gorm:"int(11)" json:"Is_active"`
	Created_at   time.Time `gorm:"timestamp" json:"created_at"`
	Updated_at   time.Time `gorm:"timestamp" json:"updated_at"`
}
