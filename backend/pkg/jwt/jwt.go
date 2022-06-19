package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// MyClaims 自定义声明结构体并内嵌 jwt 只包含了官方字段
// 我们这里需要额外记录一个 UserID 字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

var mySecret = []byte("夏天夏天悄悄过去留下小秘密~")

const TokenExpiration = 1

func GenToken(userID int64, userName string) (string, error) {
	// 创建一个自己的声明
	c := MyClaims{
		UserID:   userID,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpiration).Unix(),
			Issuer:    "konng",
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// 使用指定的 secretKey 签名并且获得完整的编码之后的字符串 token
	return token.SignedString(mySecret)
}

// ParseToken 解析 JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析 token
	// 如果是自定义 Claim 结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的 Claim 则可以直接使用 Parse 方法
		// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对 token 对象中的 Claim 进行类型断言
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验 token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
