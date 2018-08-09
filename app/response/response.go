package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ErrNo

type ErrNo struct {
	Code    int
	Message string
}

func (err ErrNo) Error() string {
	return err.Message
}

// Err

type Err struct {
	Code    int
	Message string
	Err     error
}

func (err *Err) Add(message string) error {

	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

// Response
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Json(c *gin.Context, err error, data interface{}) {
	code, message := DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// Common

func New(errNo *ErrNo, err error) *Err {
	return &Err{
		Code:    errNo.Code,
		Message: errNo.Message,
		Err:     err,
	}
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *ErrNo:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
