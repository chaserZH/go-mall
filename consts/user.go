package consts

import "time"

const EncryptMoneyKeyLength = 6

const UserInitMoney = "10000" // 初始金额 1个w

const (
	UserDefaultAvatarOss   = "http://q1.qlogo.cn/g?b=qq&nk=294350394&s=640" // OSS的默认头像
	UserDefaultAvatarLocal = "avatar.JPG"                                   // OSS的默认头像

)

const (
	AccessTokenExpireDuration  = 24 * time.Hour
	RefreshTokenExpireDuration = 10 * 24 * time.Hour
)

const (
	AccessTokenHeader    = "access_token"
	RefreshTokenHeader   = "refresh_token"
	HeaderForwardedProto = "X-Forwarded-Proto"
	MaxAge               = 3600 * 24
)
