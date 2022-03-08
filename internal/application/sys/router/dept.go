package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/xiaodulala/admin-layout/internal/application/sys/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysDeptRouter)
}

func registerSysDeptRouter(g *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysDept{}
	r := g.Group("/dept")
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
		r.GET("/deptTree", api.Get2Tree)
	}
}
