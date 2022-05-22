package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWTsecret 秘钥
var JWTsecret = []byte("TikTokABAB")

type Claims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken    创建token
func GenerateToken(userId int) (string, error) {
	now := time.Now()
	// 24 小时过期
	expireTime := now.Add(24 * time.Hour)
	claims := Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			// 当前时间 > 过期时间 那token过期了
			ExpiresAt: expireTime.Unix(),
			// 签名 没啥用
			Issuer: "TikTok",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(JWTsecret)
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {

	// 解密
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}
