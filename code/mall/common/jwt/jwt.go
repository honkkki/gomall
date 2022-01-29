package jwt

import "github.com/golang-jwt/jwt"

func GenerateToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["iat"] = iat
	claims["exp"] = iat + seconds
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
