package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const ToKenExpireDuration = time.Hour * 2

type MyClaims struct {
	UserID   uint64 `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

var MySecret = []byte("github.com/fengwei2002")

// GenToken generate a jwt token
// create a myclaims
// 		to specify diy field and expiration time and signer
// then use method to create a token object
// then return token.SignedString as jwt token ans
func GenToken(userID uint64, username string) (Token string, err error) {
	c := MyClaims{
		userID,
		"username",
		jwt.StandardClaims{ // jwt seven field
			ExpiresAt: time.Now().Add(ToKenExpireDuration).Unix(), // time
			Issuer:    "bluebell",                                 // signer
		},
	}
	Token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(MySecret)
	return
}

// ParseToken parse a jwt token to get myclaims object and a erorr
func ParseToken(tokenString string) (*MyClaims, error) {
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(t *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invaild jwt token")
}
