package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-user-microservice/application/service/authservice"
	"github.com/ydhnwb/elib-user-microservice/infrastructure/helper"
)

//AuthorizeJWT is a middleware to check the given token is valid or not
func AuthorizeJWT(jwtService authservice.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			helper.NoTokenResponse(c)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			helper.BuildErrorResponse(http.StatusBadRequest, err.Error(), nil, c)
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[Name]: ", claims["user_id"])
		} else {
			log.Println(err)
			helper.BuildErrorResponse(http.StatusUnauthorized, err.Error(), nil, c)
		}
	}
}
