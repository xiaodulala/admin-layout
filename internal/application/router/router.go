package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaodulala/admin-layout/internal/application/config"
	"github.com/xiaodulala/admin-layout/internal/pkg/code"
	"github.com/xiaodulala/admin-layout/pkg/errors"
	"github.com/xiaodulala/admin-layout/pkg/httpresponse"
	"github.com/xiaodulala/admin-layout/pkg/log"
	"github.com/xiaodulala/admin-layout/pkg/middleware"
	"github.com/xiaodulala/admin-layout/pkg/version"
	"mime"
)

var appRouters = make([]func(engine *gin.Engine), 0)

func LoadRouter(cfg *config.Config, engine *gin.Engine) {
	// 必须的插件
	engine.Use(middleware.RequestID())
	engine.Use(middleware.Context())

	// 自定义插件
	for _, m := range cfg.ServerOptions.Middlewares {
		mw, ok := middleware.Middlewares[m]
		if !ok {
			log.Warnf("can not find middleware: %s", m)

			continue
		}

		log.Infof("install middleware:%s", m)
		engine.Use(mw)
	}

	// 404
	engine.NoRoute(func(c *gin.Context) {
		httpresponse.WriteErrResponse(c, errors.WithCode(code.ErrPageNotFound, ""), nil)
	})

	//基础路由
	engine.GET("/healthz", func(c *gin.Context) {
		httpresponse.WriteSuccessResponse(c, map[string]string{"status": "ok"})
	})

	engine.GET("version", func(c *gin.Context) {
		httpresponse.WriteSuccessResponse(c, version.Get())
	})

	// 静态路由文件
	if err := mime.AddExtensionType(".js", "application/javascript"); err != nil {
		return
	}
	engine.Static("/static", "./static")

	if cfg.ServerOptions.Mode != "release" {
		engine.Static("/form-generator", "./static/form-generator")
	}

	// swagger 单独

	// 系统路由和业务路由
	for _, f := range appRouters {
		f(engine)
	}
}
