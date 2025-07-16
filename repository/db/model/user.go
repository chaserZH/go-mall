package model

import (
	"github.com/CocaineCong/secret"
	"go-mall/conf"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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

const (
	PassWordCost        = 12       // 密码加密难度
	Active       string = "active" // 激活用户
)

// SetPassword 设置密码
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

// EncryptMoney 加密金额
func (u *User) EncryptMoney(key string) (money string, err error) {
	aesObj, err := secret.NewAesEncrypt(conf.Config.EncryptSecret.MoneySecret, key, "", secret.AesEncrypt128, secret.AesModeTypeCBC)
	if err != nil {
		return
	}
	money = aesObj.SecretEncrypt(u.Money)

	return
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}

// AvatarURL 头像地址
func (u *User) AvatarURL() string {
	return ""
	//if conf.Config.System.UploadModel == consts.UploadModelOss {
	//	return u.Avatar
	//}
	//pConfig := conf.Config.PhotoPath
	//return pConfig.PhotoHost + conf.Config.System.HttpPort + pConfig.AvatarPath + u.Avatar
}
