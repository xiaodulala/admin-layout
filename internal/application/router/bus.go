package router

import "github.com/xiaodulala/admin-layout/internal/application/bus/router"

func init() {
	appRouters = append(appRouters, router.InitRouter)
}
