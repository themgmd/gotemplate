package jwt

import "github.com/golang-jwt/jwt/v5"

func Generate() {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	_ = claims
}

func Verify() {

}
