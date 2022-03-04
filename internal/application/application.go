package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/xiaodulala/admin-layout/internal/application/config"
	"github.com/xiaodulala/admin-layout/internal/application/router"
	"github.com/xiaodulala/admin-layout/pkg/db/mysql"
	"github.com/xiaodulala/admin-layout/pkg/log"
	"github.com/xiaodulala/admin-layout/pkg/mycasbin"
	"github.com/xiaodulala/admin-layout/pkg/utils"
	"gorm.io/gorm"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

type Application struct {
	prepare bool
	orm     *gorm.DB
	casbin  *casbin.SyncedEnforcer
	engine  *gin.Engine
	config  *config.Config
}

func createApp(cfg *config.Config) (*Application, error) {
	// todo 目前支持一种，后续封装多种
	db, err := mysql.New(cfg.MysqlOptions)
	if err != nil {
		return nil, err
	}

	e, err := mycasbin.NewEnforcer(db)
	if err != nil {
		return nil, err
	}

	// gin配置
	gin.SetMode(cfg.ServerOptions.Mode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Infof("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	return &Application{
		orm:    db,
		casbin: e,
		engine: gin.New(),
		config: cfg,
	}, nil
}

func (app *Application) PrepareRun() *Application {

	//初始化路由
	router.LoadRouter(app.config, app.engine)

	app.prepare = true
	return app
}

func (app *Application) Run() error {

	if !app.prepare {
		return errors.New("application is not prepare,please call Prepare() before Running")
	}

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(app.config.ServerOptions.BindAddress, strconv.Itoa(app.config.ServerOptions.BindPort)),
		Handler: app.engine,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err.Error())
		}
	}()

	// 打印服务信息
	app.printRunInfo()

	//服务优雅关停
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Infof("%s shutdown server...\r\n", time.Now().Format("2006-01-02 15:04:05"))

	if err := httpServer.Shutdown(context.Background()); err != nil {
		log.Fatalf("server shutdown err:%s", err.Error())
	}

	log.Infof("server existing")

	return nil
}

func (app *Application) printRunInfo() {
	fmt.Println(color.GreenString(LogoContent))

	fmt.Println(color.YellowString("server run at:"))
	fmt.Printf("-  Local:   http://localhost:%d/ \n", app.config.ServerOptions.BindPort)
	fmt.Printf("-  Network: http://%s:%d/ \r\n", utils.GetLocalIP(), app.config.ServerOptions.BindPort)
	fmt.Printf("Enter Control + C Shutdown Server \r\n")
}
