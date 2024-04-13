package services

import (
	"auth/errors"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	key []byte
}

func NewJWTService(key []byte) *JwtService {
	return &JwtService{
		key: key,
	}
}

func (j JwtService) CreateKey(email string, roles []string, pid string) (*string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  email,
		"roles": roles,
		"pid":   pid,
	})
	signed, err := t.SignedString(j.key)
	if err != nil {
		return nil, err
	}
	return &signed, nil
}

func (j JwtService) ValidateToken(tokenString string) (*jwt.MapClaims, *errors.ErrorStruct) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the token's algorithm is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.key, nil
	})
	if err != nil {
		log.Println(err)
		return nil, errors.NewError(err.Error(), 401)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	} else {
		return nil, errors.NewError("Unathorized", 401)
	}
}
