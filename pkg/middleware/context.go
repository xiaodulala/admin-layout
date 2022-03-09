package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaodulala/admin-layout/pkg/log"
)

const (
	Username  = "username"
	UserId    = "userId"
	RoleKey   = "roleKey"
	DataScope = "dataScope"
	RoleId    = "roleId"
	RoleName  = "roleName"
)

// Context 将上下文的键值设置为log包支持的键值，以通过中间件打印gin日志时带有上下文信息
// gin框架log中间件使用了自定义logger中间件。
// 自定义的logger中间件和gin框架默认的日志中间件完全一样，只是最后多了一步打印。
// 打印的时候调用是log包的log.L(c)方法。此方法会提取上下文中的指定的键值。此方法的作用就是将键值重新设置为log包支持的上下文键值
func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(log.KeyRequestID, c.GetString(XRequestIDKey))
		c.Set(log.KeyUsername, c.GetString(Username))
		c.Next()
	}
}

func GetUserId(c *gin.Context) int {
	return int(c.GetFloat64(UserId))
}

func GetUserName(c *gin.Context) string {
	return c.GetString(Username)
}

func GetRoleId(c *gin.Context) int {
	return int(c.GetFloat64(RoleId))
}

func GetRoleKey(c *gin.Context) string {
	return c.GetString(RoleKey)
}

func GetRoleName(c *gin.Context) string {
	return c.GetString(RoleName)
}

func GetDataScope(c *gin.Context) string {
	return c.GetString(DataScope)
}
