package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/elib-user-microservice/application/service/authservice"
	"github.com/ydhnwb/elib-user-microservice/domain/entity"
	"github.com/ydhnwb/elib-user-microservice/infrastructure/dto"
	"github.com/ydhnwb/elib-user-microservice/infrastructure/helper"
)

//AuthController is a contract
type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type authController struct {
	jwtService  authservice.JWTService
	authService authservice.AuthService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(jwtService authservice.JWTService, authService authservice.AuthService) AuthController {
	return &authController{
		jwtService:  jwtService,
		authService: authService,
	}
}

func (ctl *authController) Login(ctx *gin.Context) {
	loginDTO := dto.UserLoginDTO{}
	e := ctx.ShouldBind(&loginDTO)
	if e != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, e.Error(), helper.EmptyObj{}, ctx)
		return
	}

	res := ctl.authService.VerifyCredential(loginDTO.Email, loginDTO.Email)
	if v, ok := res.(entity.User); ok {
		token := ctl.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = token
		helper.BuildResponse(http.StatusOK, v, ctx)
		return
	}
	helper.BuildErrorResponse(http.StatusUnauthorized, "Invalid credential", helper.EmptyObj{}, ctx)
}

func (ctl *authController) Register(ctx *gin.Context) {
	registerDTO := dto.UserRegisterDTO{}
	e := ctx.ShouldBind(&registerDTO)
	if e != nil {
		helper.BuildErrorResponse(http.StatusBadRequest, e.Error(), helper.EmptyObj{}, ctx)
		return
	}
	if ctl.authService.IsEmailDuplicate(registerDTO.Email) {
		helper.BuildErrorResponse(http.StatusConflict, "Duplicate email", helper.EmptyObj{}, ctx)
	} else {
		created, err := ctl.authService.RegisterUser(registerDTO)
		if err != nil {
			helper.BuildErrorResponse(http.StatusBadRequest, err.Error(), helper.EmptyObj{}, ctx)
			return
		}
		token := ctl.jwtService.GenerateToken(strconv.FormatUint(created.ID, 10))
		created.Token = token
		helper.BuildResponse(http.StatusCreated, created, ctx)
	}
}
