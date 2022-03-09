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
	r := g.Group("/dept").Use(authMiddleware.MiddlewareFunc())
	{
		r.GET("", api.GetPage)
		r.GET("/:deptId", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:deptId", api.Update)
		r.DELETE("", api.Delete)
		// fixme 待测试用处
		// r.GET("/deptTree", api.Get2Tree)
	}
}
