package ginauth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/xiaodulala/admin-layout/internal/application/sys/models"
	"github.com/xiaodulala/admin-layout/internal/pkg/runtime"
	"github.com/xiaodulala/admin-layout/pkg/log"
	"github.com/xiaodulala/admin-layout/pkg/middleware"
	"time"
)

func JWTAuth() (*jwt.GinJWTMiddleware, error) {
	ginJWT, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            viper.GetString("jwt.Realm"),
		SigningAlgorithm: "HS256",
		Key:              []byte(viper.GetString("jwt.Key")),
		Timeout:          viper.GetDuration("jwt.timeout"),
		MaxRefresh:       viper.GetDuration("jwt.max-refresh"),
		PayloadFunc:      payloadFunc,
		IdentityHandler:  identityHandler,
		Authenticator:    authenticator,
		Authorizator:     authorizator,
		Unauthorized:     unauthorized,
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(code, gin.H{"token": token, "expire": expire.Format(time.RFC3339)})
		},
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(code, nil)
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(code, gin.H{"token": token, "expire": expire.Format(time.RFC3339)})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	return ginJWT, err
}

// payloadFunc 登录成功后 token载荷
func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(*models.SysUser)
		r, _ := v["role"].(*models.SysRole)
		return jwt.MapClaims{
			middleware.UserId:    u.UserId,
			middleware.Username:  u.Username,
			middleware.DataScope: r.DataScope,
			middleware.RoleId:    r.RoleId,
			middleware.RoleKey:   r.RoleKey,
			middleware.RoleName:  r.RoleName,
		}
	}
	return jwt.MapClaims{}
}

// identityHandler 解析出token载荷
func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		middleware.UserId:    claims[middleware.UserId],
		middleware.Username:  claims[middleware.Username],
		middleware.DataScope: claims[middleware.DataScope],
		middleware.RoleId:    claims[middleware.RoleId],
		middleware.RoleKey:   claims[middleware.RoleKey],
		middleware.RoleName:  claims[middleware.RoleName],
	}
}

func authenticator(c *gin.Context) (interface{}, error) {
	var loginInfo Login
	defer func() {
		// 记录登录日志
	}()

	if err := c.ShouldBind(&loginInfo); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	// todo 验证码验证

	db := runtime.RunTime.Orm()
	user, role, err := loginInfo.GetUser(db)
	if err != nil {
		log.Warnf("%s login failed!", loginInfo.Username)
		return nil, jwt.ErrFailedAuthentication
	}
	return map[string]interface{}{"user": &user, "role": &role}, nil
}

func authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(map[string]interface{}); ok {
		c.Set(middleware.UserId, v[middleware.UserId])
		c.Set(middleware.Username, v[middleware.Username])
		c.Set(middleware.DataScope, v[middleware.DataScope])
		c.Set(middleware.RoleId, v[middleware.RoleId])
		c.Set(middleware.RoleKey, v[middleware.RoleKey])
		c.Set(middleware.RoleName, v[middleware.RoleName])
		return true
	}
	return false
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"message": message})
}
