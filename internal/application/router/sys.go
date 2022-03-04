package router

import "github.com/xiaodulala/admin-layout/internal/application/sys/router"

func init() {
	appRouters = append(appRouters, router.InitRouter)
}
