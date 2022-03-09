package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/xiaodulala/admin-layout/internal/application/sys/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysUserRouter)
}

func registerSysUserRouter(g *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysUser{}
	r := g.Group("/user").Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("", api.Update)
		r.DELETE("/:userId", api.Delete)
		r.GET("/profile", api.GetProfile)
		r.POST("/avatar", api.InsetAvatar)
		r.PUT("/pwd/set", api.UpdatePwd)
		r.PUT("/pwd/reset", api.ResetPwd)
		r.PUT("/status", api.UpdateStatus)
	}
}
