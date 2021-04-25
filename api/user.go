package api

import (
	"fmt"
	"singo/serializer/handler"
	"singo/service"
	
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) (handler.ActionResponse, error) {
	var service service.UserRegisterService
	
	if err := c.ShouldBind(&service); err != nil {
		return nil, err
	}
	return service.Register()
	
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) (handler.ActionResponse, error) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err != nil {
		return nil, err
	}
	return service.Login(c)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) (handler.ActionResponse, error) {
	s := sessions.Default(c)
	s.Clear()
	err:= s.Save()
	if err!=nil{
		fmt.Println(err.Error())
	}
	return "登出成功", nil
	
}
