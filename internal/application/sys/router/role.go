package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/xiaodulala/admin-layout/internal/application/sys/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysRoleRouter)
}

func registerSysRoleRouter(g *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysRole{}
	r := g.Group("/role", authMiddleware.MiddlewareFunc())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.PUT("/role-status", api.Update2Status)
		r.PUT("/role-scope", api.Update2DataScope)
	}
}
