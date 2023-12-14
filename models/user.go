package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	FullName string `gorm:"varchar(300)" json:"full_name"`
	Username string `gorm:"varchar(300);not null" json:"username"`
	Password string `gorm:"varchar(300);not null" json:"password"`
}
