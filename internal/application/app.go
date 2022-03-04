package application

import (
	"github.com/xiaodulala/admin-layout/internal/application/config"
	"github.com/xiaodulala/admin-layout/internal/application/options"
	"github.com/xiaodulala/admin-layout/pkg/app"
	"github.com/xiaodulala/admin-layout/pkg/log"
)

const (
	baseName    = "application"
	appName     = "classified carrier manager"
	commandDesc = "command help user run a  server"
)

func NewApp() *app.App {

	opts := options.NewOptions()
	return app.NewApp(appName, baseName,
		app.WithDescription(commandDesc),
		app.WithOptions(opts),
		app.WithRunFunc(run(opts)))
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.LogOptions)
		defer log.Flush()

		//生成配置
		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}
		return Run(cfg)
	}
}
