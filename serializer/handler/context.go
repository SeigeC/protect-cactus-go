package handler

import (
	"net/http"
	"singo/model"
	"singo/util"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// pageSize limit value and default value
const (
	DefaultPageSize int64 = 20
	MaxPageSize     int64 = 100
)

// split limit and default value
const (
	MaxParamSplitLength = 100

	DefaultSplitSep = ","
)

// Context request context
type Context struct {
	*gin.Context

	User *model.User
}

// UserAgent returns the "User-Agent" header
func (ctx *Context) UserAgent() string {
	return ctx.GetHeader("User-Agent")
}

// ClientIP returns the client IP
func (ctx *Context) ClientIP() string {
	return ctx.ClientIP()
}

// Request return the request
func (ctx *Context) Request() *http.Request {
	return ctx.Request()
}

// GetDefaultParam gets params from query if successful, otherwise post form, otherwise defaultValue
func (ctx *Context) GetDefaultParam(key string, defaultValue string) string {
	value, ok := ctx.GetQuery(key)
	if ok {
		return value
	}
	return ctx.DefaultPostForm(key, defaultValue)
}

// GetParam gets params from query if successful, otherwise post form
func (ctx *Context) GetParam(key string) (string, bool) {
	value, ok := ctx.GetQuery(key)
	if ok {
		return value, true
	}
	return ctx.GetPostForm(key)
}

// GetDefaultParamString returns string without leading and trailing white space if successful,
// otherwise post form, otherwise defaultValue
func (ctx *Context) GetDefaultParamString(key string, defaultValue string) string {
	value, ok := ctx.GetParamString(key)
	if ok {
		return value
	}
	return defaultValue
}

// GetParamString returns string without leading and trailing white space
func (ctx *Context) GetParamString(key string) (string, bool) {
	raw, ok := ctx.GetParam(key)
	if !ok {
		return "", false
	}
	return strings.TrimSpace(raw), true
}

// GetDefaultParamInt gets params from query if successful, otherwise post form, otherwise defaultValue
func (ctx *Context) GetDefaultParamInt(key string, defaultValue int) (int, error) {
	val, ok := ctx.GetParam(key)
	if !ok {
		return defaultValue, nil
	}
	value, err := strconv.Atoi(val)
	return value, err
}

// GetDefaultParamInt64 gets params from query if successful, otherwise post form, otherwise defaultValue
func (ctx *Context) GetDefaultParamInt64(key string, defaultValue int64) (int64, error) {
	val, ok := ctx.GetParam(key)
	if !ok {
		return defaultValue, nil
	}
	value, err := strconv.ParseInt(val, 10, 64)
	return value, err
}

// GetParamInt64 gets params from query if successful, otherwise post form
func (ctx *Context) GetParamInt64(key string) (int64, error) {
	val, ok := ctx.GetParam(key)
	if !ok {
		return 0, ErrEmptyParam
	}
	value, err := strconv.ParseInt(val, 10, 64)
	return value, err
}

// GetParamInt64ArrayFromString gets params to int64 array from query if successful, otherwise post form
// NOTICE: sep ???????????????????????????
func (ctx *Context) GetParamInt64ArrayFromString(key string, sep ...string) ([]int64, error) {
	val, ok := ctx.GetParam(key)
	if !ok {
		return nil, ErrEmptyParam
	}

	splitSep := DefaultSplitSep
	if len(sep) == 1 {
		splitSep = sep[0]
	} else if len(sep) > 1 {
		panic("invalid sep param")
	}
	value, err := util.SplitToInt64Array(val, splitSep)
	if err != nil {
		return nil, ErrInvalidParam
	}
	length := len(value)
	if length == 0 || length > MaxParamSplitLength {
		return nil, ErrInvalidParam
	}
	return value, nil
}

// PageOption param page option
type PageOption struct {
	DefaultPageSize int64 // pageSize, ?????? 20
	MaxPageSize     int64 // ???????????????????????? pageSize, ?????? 100
}

// GetParamPage gets the p and page_size parameters
// NOTICE: return: p ?????????: 1, pageSize ?????????: 20
func (ctx *Context) GetParamPage(opt ...*PageOption) (p, pageSize int64, err error) {
	p, err = ctx.GetDefaultParamInt64("p", 1)
	if err != nil || p <= 0 {
		return 0, 0, ErrInvalidParam
	}
	option := getPageOption(opt...)
	pageSize, err = getPageSize(ctx, option)
	if err != nil || pageSize <= 0 || pageSize > option.MaxPageSize {
		return 0, 0, ErrInvalidParam
	}
	return
}

// TODO: ??????????????????????????? pagesize ??????????????? page_size ???????????????
// NOTICE: ???????????? size, ?????? pagesize, ????????? page_size
func getPageSize(ctx *Context, option *PageOption) (pageSize int64, err error) {
	size, exists := ctx.GetParam("pagesize")
	if !exists {
		size, exists = ctx.GetParam("page_size")
		if !exists {
			return option.DefaultPageSize, nil
		}
	}
	return strconv.ParseInt(size, 10, 64)
}

func getPageOption(opt ...*PageOption) *PageOption {
	length := len(opt)
	if length == 0 {
		return &PageOption{DefaultPageSize: DefaultPageSize, MaxPageSize: MaxPageSize}
	}
	if length > 1 {
		panic("invalid page option")
	}
	if opt[0].DefaultPageSize == 0 {
		opt[0].DefaultPageSize = DefaultPageSize
	}
	if opt[0].MaxPageSize == 0 {
		opt[0].MaxPageSize = MaxPageSize
	}
	if opt[0].DefaultPageSize < 0 || opt[0].MaxPageSize < opt[0].DefaultPageSize {
		panic("invalid page option")
	}
	return opt[0]
}

// Token ??? cookie ????????? token, ????????????????????????
func (ctx *Context) Token() string {
	token, _ := ctx.Cookie("token")
	return token
}
