package httputil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrNone           = "0"
	ErrNotFound       = "1"
	ErrUnauthorized   = "2"
	ErrForbidden      = "3"
	ErrBadRequest     = "4"
	ErrValidation     = "5"
	ErrInternalServer = "6"
)

type JsonResponse[T any] struct {
	Code    string         `json:"code,omitempty"`
	Message string         `json:"message,omitempty"`
	Data    T              `json:"data,omitempty"`
	Extra   map[string]any `json:"extra,omitempty"`
}

func Success[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, JsonResponse[T]{
		Code: ErrNone,
		Data: data,
	})
}

func Error[T any](c *gin.Context, statusCode int, code, message string) {
	c.JSON(statusCode, JsonResponse[T]{
		Code:    code,
		Message: message,
	})
}

func SuccessWithExtra[T any](c *gin.Context, data T, extra map[string]any) {
	c.JSON(http.StatusOK, JsonResponse[T]{
		Code:  ErrNone,
		Data:  data,
		Extra: extra,
	})
}

func ValidateError(c *gin.Context, message string) {
	Error[string](c, http.StatusBadRequest, ErrValidation, message)
}

func InternalServerError(c *gin.Context, message string) {
	Error[string](c, http.StatusInternalServerError, ErrInternalServer, message)
}

func BadRequestError(c *gin.Context, message string) {
	Error[string](c, http.StatusBadRequest, ErrBadRequest, message)
}

func NotFoundError(c *gin.Context, message string) {
	Error[string](c, http.StatusNotFound, ErrNotFound, message)
}

func UnauthorizedError(c *gin.Context, message string) {
	Error[string](c, http.StatusUnauthorized, ErrUnauthorized, message)
}

func ForbiddenError(c *gin.Context, message string) {
	Error[string](c, http.StatusForbidden, ErrForbidden, message)
}
