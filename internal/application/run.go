package application

import "github.com/xiaodulala/admin-layout/internal/application/config"

func Run(cfg *config.Config) error {
	appServer, err := createApp(cfg)
	if err != nil {
		return err
	}

	return appServer.PrepareRun().Run()
}
