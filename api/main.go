package api

import (
	"singo/model"
	"singo/serializer/handler"
	
	"github.com/gin-gonic/gin"
)

// Ping 状态检查页面
func Ping(c *handler.Context) (handler.ActionResponse, error) {
	return "Pong",nil
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}


