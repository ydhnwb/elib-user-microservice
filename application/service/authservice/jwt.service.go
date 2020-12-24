package authservice

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTService interface is a contract
type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

//extending jwt standar claims.. will move this up soomewhere
type jwtCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

//NewJWTService creates a new instance JWTService
func NewJWTService() JWTService {
	return &jwtService{
		secretKey: "hardcoded,willmovethistodotenv",
		issuer:    "ydhnwb",
	}
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(tokenForFunc *jwt.Token) (interface{}, error) {
		if _, ok := tokenForFunc.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", tokenForFunc.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}

//Will change this with refreshToken someday...
func (s *jwtService) GenerateToken(userID string) string {
	claims := &jwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}
