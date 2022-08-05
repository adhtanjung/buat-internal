package auth

import (
	// "github.com/adhtanjung/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

type authRoute struct {
	route   *gin.RouterGroup
	handler AuthHandler
}

type AuthRoute interface {
	Init(route *gin.RouterGroup)
}

func NewAuthRoute(route *gin.RouterGroup, handler AuthHandler) AuthRoute {
	return &authRoute{
		route,
		handler,
	}
}

func (a *authRoute) Init(route *gin.RouterGroup) {

	auth := route.Group("/auth")
	{
		auth.POST("/register", a.handler.Register())
		// auth.POST("/login", a.handler.Login())
	}
	a.handler.Register()

}
