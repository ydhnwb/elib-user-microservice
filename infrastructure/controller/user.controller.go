package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-user-microservice/application/service/authservice"
	"github.com/ydhnwb/elib-user-microservice/application/service/userservice"
	"github.com/ydhnwb/elib-user-microservice/infrastructure/dto"
	"github.com/ydhnwb/elib-user-microservice/infrastructure/helper"
)

//UserController is a contract
type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	jwtService  authservice.JWTService
	userService userservice.UserService
}

//NewUserController creates a new instance of UserController
func NewUserController(jwtService authservice.JWTService, userService userservice.UserService) UserController {
	return &userController{
		jwtService:  jwtService,
		userService: userService,
	}
}

func (c *userController) Update(ctx *gin.Context) {
	userUpdateDTO := dto.UserUpdateDTO{}
	e := ctx.ShouldBind(&userUpdateDTO)
	if e != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, e.Error(), helper.EmptyObj{}, ctx)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, _ := c.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	id, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	userUpdateDTO.ID = id
	res, err := c.userService.UpdateProfile(userUpdateDTO)
	if err != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, e.Error(), helper.EmptyObj{}, ctx)
		return
	}
	helper.BuildResponse(http.StatusOK, res, ctx)
}

func (c *userController) Profile(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, _ := c.jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	user := c.userService.GetOwnProfile(fmt.Sprintf("%v", claims["user_id"]))
	helper.BuildResponse(http.StatusOK, user, ctx)
}
