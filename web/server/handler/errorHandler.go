package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerFunc func(*gin.Context) error

// api错误的结构体
type APIException struct {
	Code      int    `json:"-"`
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
	Request   string `json:"request"`
}

// 实现接口
func (e *APIException) Error() string {
	return e.Msg
}

// 500 错误处理
func ServerError() *APIException {
	return newAPIException(http.StatusInternalServerError, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

// 未知错误
func UnknownError(message string) *APIException {
	return newAPIException(http.StatusForbidden, http.StatusForbidden, message)
}

// 参数错误
func ParameterError(message string) *APIException {
	return newAPIException(http.StatusBadRequest, http.StatusBadRequest, message)
}

func newAPIException(code int, errorCode int, msg string) *APIException {
	return &APIException{
		Code:      code,
		ErrorCode: errorCode,
		Msg:       msg,
	}
}

func Wrapper(handler HandlerFunc) func(*gin.Context) {
	return func(ctx *gin.Context) {
		err := handler(ctx)
		if err != nil {
			var apiException *APIException
			if h, ok := err.(*APIException); ok {
				apiException = h
			} else if e, ok := err.(error); ok {
				if gin.Mode() == "debug" {
					// 错误
					apiException = UnknownError(e.Error())
				} else {
					// 未知错误
					apiException = UnknownError(e.Error())
				}
			} else {
				apiException = ServerError()
			}
			apiException.Request = ctx.Request.Method + " " + ctx.Request.URL.String()
			ctx.JSON(apiException.Code, apiException)
			return
		}
	}
}
