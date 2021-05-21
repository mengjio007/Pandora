package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

const jwtKey = "0x9527"

// 创建token
/*
 *    HS256 --- HMAC SHA256
 */
// CreateTokenHMAC
func CreateTokenHMAC(m map[string]interface{}, keys ...string) string {
	key := jwtKey
	if len(keys) > 0 {
		key = keys[0]
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	for index, val := range m {
		claims[index] = val
	}
	token.Claims = claims
	signedString, _ := token.SignedString([]byte(key))
	return signedString
}

// 解析token
/*
 *
 */
// ParseToKenHMAC
func ParseTokenHMAC(tokenString string, keys ...string) (map[string]interface{}, bool) {
	key := jwtKey
	if len(keys) > 0 {
		key = keys[0]
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mapData := make(map[string]interface{})
		for index, val := range claims {
			mapData[index] = fmt.Sprintf("%v", val)
		}
		return mapData, true
	} else {
		return nil, false
	}
}
