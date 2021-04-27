package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github/qm012/nacos-adress/global"
	"time"
)

type JWTService interface {
	GenerateToken(username string, admin bool) string
	ValidateToken(tokenString string) (*jwtCustomClaims, error)
}

type jwtCustomClaims struct {
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
	expire    int
}

func (j *jwtService) GenerateToken(username string, admin bool) string {
	claims := &jwtCustomClaims{
		Username: username,
		Admin:    admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(j.expire) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    j.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(fmt.Sprintf("token.SignedString failed:%v", err.Error()))
	}
	return t
}

func (j *jwtService) ValidateToken(tokenString string) (*jwtCustomClaims, error) {

	var jwtClaims = &jwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, jwtClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwtCustomClaims)
	if !ok && !token.Valid {
		err = errors.New("authorization failed")
		return nil, err
	}
	return claims, nil
}

func NewJwtService() JWTService {
	return &jwtService{
		secretKey: global.Server.Jwt.Secret,
		expire:    global.Server.Jwt.Expire,
		issuer:    "nacos-address",
	}
}
