package services

import "github.com/golang-jwt/jwt/v5"

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
		"PID":   pid,
	})
	signed, err := t.SignedString(j.key)
	if err != nil {
		return nil, err
	}
	return &signed, nil
}
