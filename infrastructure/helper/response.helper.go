package helper

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//Response is used for static shape json return
type Response struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//EmptyObj object is used when data doesnt want to be null on json
type EmptyObj struct{}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponse(statusCode int, data interface{}, ctx *gin.Context) {
	res := Response{
		Message: messageMapper(statusCode),
		Errors:  []string{},
		Data:    data,
	}
	ctx.JSON(statusCode, res)
}

//BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponse(statusCode int, err string, data interface{}, ctx *gin.Context) {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Message: messageMapper(statusCode),
		Errors:  splittedError,
		Data:    data,
	}
	ctx.AbortWithStatusJSON(statusCode, res)
}

//NoTokenResponse return a json with 400 Bad request because user doesnt provide
//Authorization in header
func NoTokenResponse(ctx *gin.Context) {
	res := Response{
		Message: "No token provided",
		Errors:  "No token found in header",
		Data:    nil,
	}
	ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func messageMapper(statusCode int) string {
	switch statusCode {
	case 200:
		return "OK!"
	case 201:
		return "Created!"
	case 400:
		return "Bad request. Cannot process your request"
	case 404:
		return "The given data / url not found!"
	case 401:
		return "Unauthorized."
	case 403:
		return "You dont have permission to do this request!"
	case 409:
		return "Conflict"
	case 500:
		return "Internal Server Error."
	case 503:
		return "Service Unaivalable. Try again later or contact your back-end"
	default:
		return ""
	}
}
