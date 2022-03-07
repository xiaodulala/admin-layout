package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

func InitRouter(engine *gin.Engine) {
	// 注册系统业务路由
	g := engine.Group("/api/v1/sys")

	// 无需认证的路由
	for _, f := range routerNoCheckRole {
		f(g)
	}

	// 需要认证的路由
	for _, f := range routerCheckRole {
		f(g, nil)
	}
}
