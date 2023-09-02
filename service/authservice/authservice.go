package authservice

import (
	"Q/A-GameApp/entity"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

type Config struct {
	SignKey               string `json:"sign_key"`
	AccessExpirationTime  time.Duration
	RefreshExpirationTime time.Duration
	AccessSubject         string
	RefreshSubject        string
}
type Service struct {
	config Config
}

func New(config Config) Service {
	return Service{
		config: config,
	}

}
func (s Service) CreateAccessToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.config.AccessSubject, s.config.AccessExpirationTime)
}
func (s Service) CreateRefreshToken(user entity.User) (string, error) {
	return s.createToken(user.ID, s.config.RefreshSubject, s.config.RefreshExpirationTime)
}
func (s Service) VarifyToken(tokenStr string) (*Claims, error) {
	strings.Replace(tokenStr, "Bearer", "", 1)
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	})
	if err != nil {
		return nil, err
	}
	if clamis, ok := token.Claims.(*Claims); ok && token.Valid {
		return clamis, nil
	} else {
		return nil, err
	}
}

func (s Service) createToken(userID uint, subject string, expireDuration time.Duration) (string, error) {
	//create a new token
	// TODO : replace with rsa256 - https://github.com/golang-jwt/jwt/
	//see our claims
	Claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
		UserID: userID,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	tokenString, err := accessToken.SignedString([]byte(s.config.SignKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
