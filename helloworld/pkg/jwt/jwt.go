package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

//JwtDecode Jwt 解密
func JWTDecode(jwtSecret string, jwtInfo string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtInfo, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, errors.New("invalid jwt info")
	}
	if token.Valid {
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok {
			return claims, nil
		}
	}
	return nil, errors.New("bad jwt info")
}

//JwtEncode Jwt 加密
func JWTEncode(jwtSecret string, value map[string]interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(value))
	return token.SignedString([]byte(jwtSecret))
}
