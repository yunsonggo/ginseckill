package common

import (
	"2022/ginseckill/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	jc = config.Conf.Jwt
	jwtKey = []byte(jc.JwtKey)
	jwtIssuer = jc.Issuer
	exptime = jc.ExpTime
)

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

// 发放token
func ReleaseToken(userId int64) (string,error) {
	// 有效期
	expirationTime := time.Now().Add(time.Duration(exptime) * 24 * time.Hour)
	claims := &Claims{
		UserId:userId,
		StandardClaims:jwt.StandardClaims{
			// 有效期
			ExpiresAt: expirationTime.Unix(),
			// 开始时间
			IssuedAt: time.Now().Unix(),
			// 发放者
			Issuer: jwtIssuer,
			// 主题
			Subject:"user token",
		},
	}
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString ,err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func ParseToken(tokenString string) (*jwt.Token,*Claims,error) {
	claims := &Claims{}
	token,err := jwt.ParseWithClaims(tokenString,claims,func(token *jwt.Token) (i interface{},err error) {
		return jwtKey,nil
	})
	return token,claims,err
}