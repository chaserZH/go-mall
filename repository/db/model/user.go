package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string `GORM:"size:1000"`
	Money          string
	Relations      []User `gorm:"many2many:relation;"`
}
