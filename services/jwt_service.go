package services

import (
	"fmt"
	"log"
	"time"

	"github.com/a180024/golang-api/config"
	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(name string) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "Your-Organization",
	}
}

func getSecretKey() string {
	c := config.GetConfig()
	secretKey := c.GetString("secretKey")
	if secretKey == "" {
		secretKey = "secret"
	}
	return secretKey
}

func (jwtSrv *jwtService) GenerateToken(username string) string {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    jwtSrv.issuer,
		IssuedAt:  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return ss
}

func (jwtSrv *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSrv.secretKey), nil
	})
}
