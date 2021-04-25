package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess       = 0
	CodeLoginRequired = 100010006
	CodeUnknownError  = 100010007

	CodeEmptyParam   = 201010001
	CodeInvalidParam = 201010002
)

// errors
var (
	ErrLoginRequired = NewActionError(http.StatusForbidden, CodeLoginRequired, "请先登录")
	ErrBadRequest    = NewActionError(http.StatusBadRequest, CodeUnknownError, "错误的请求")
	ErrEmptyParam    = NewActionError(http.StatusBadRequest, CodeEmptyParam, "参数不可为空")
	ErrInvalidParam  = NewActionError(http.StatusBadRequest, CodeInvalidParam, "参数不合法")

	ErrEmptyValue = errors.New("empty value")
)

// ResponseError response error interface
type ResponseError interface {
	// StatusCode 响应状态码
	StatusCode() int
	// ErrorCode 错误代码
	ErrorCode() int
	// ErrorInfo 响应错误信息
	ErrorInfo() interface{}
	// ErrorMessage 响应错误报错
	ErrorMessage() string

	error
}

// ActionError for action error
type ActionError struct {
	Status  int
	Code    int
	Message string
	Info    string

	err error
}

// NewActionError returns a new ActionError
func NewActionError(status int, code int, msg string) *ActionError {
	return &ActionError{
		Status:  status,
		Code:    code,
		Message: msg,
	}
}

// New set error
func (e *ActionError) New(err error) *ActionError {
	e.err = err
	return e
}

// StatusCode status code
func (e *ActionError) StatusCode() int {
	return e.Status
}

// ErrorCode error code
func (e *ActionError) ErrorCode() int {
	return e.Code
}

// ErrorInfo response info
// returns {"msg":"test err", e.Info...}
func (e *ActionError) ErrorInfo() interface{} {
	if e.Info != "" {
		return e.Info
	}
	return e.Error()
}

// ErrorInfo response info
// returns {"msg":"test err", e.Info...}
func (e *ActionError) ErrorMessage() string {
	if e.err != nil && gin.Mode() != gin.ReleaseMode {
		return e.err.Error()
	}
	return ""
}

func (e *ActionError) Error() string {
	return e.Message
}

// Action structure
type Action struct {
	Method        Method
	Action        ActionFunc
	LoginRequired bool

	handler gin.HandlerFunc
}

// GetHandler for request handler
func (a *Action) GetHandler() gin.HandlerFunc {
	return a.handler
}

// NewAction creates new action
func NewAction(method Method, handler ActionFunc, loginRequired bool) *Action {
	a := Action{Method: method, Action: handler, LoginRequired: loginRequired}
	a.handler = func(c *gin.Context) {
		// init context
		ctx := Context{Context: c}

		if a.LoginRequired {
			abortError(c, ErrLoginRequired)
			return
		}

		r, err := a.Action(&ctx)
		if err != nil {
			abortError(c, err)
			return

		}
		c.JSON(http.StatusOK, Response{Code: CodeSuccess, Data: r})
	}
	return &a
}

func abortError(c *gin.Context, err error) {
	switch v := err.(type) {
	case ResponseError:
		c.Status(v.StatusCode())
		abortWithError(c, v.StatusCode(), v.ErrorCode(), v.ErrorInfo(), v.Error())
	default:
		c.Status(http.StatusInternalServerError)
		abortWithError(c, http.StatusInternalServerError, 500, err.Error(), "")
	}
}

func abortWithError(c *gin.Context, status, code int, info interface{}, err string) {
	c.AbortWithStatusJSON(status, Response{Code: code, Data: info, Error: err})
}
