package api

import (
	"fmt"
	vd "github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/xiaodulala/admin-layout/internal/application/common/service"
	"github.com/xiaodulala/admin-layout/internal/pkg/code"
	"github.com/xiaodulala/admin-layout/internal/pkg/runtime"
	"github.com/xiaodulala/admin-layout/pkg/errors"
	"github.com/xiaodulala/admin-layout/pkg/log"
	"gorm.io/gorm"
)

type Api struct {
	Context *gin.Context
	Orm     *gorm.DB
	Errors  error
}

func (api *Api) AddError(err error) {
	if api.Errors == nil {
		api.Errors = err
	} else if err != nil {
		api.Errors = fmt.Errorf("%v; %w", api.Errors, err)
	}
}

// MakeContext 设置http上下文
func (api *Api) MakeContext(c *gin.Context) *Api {
	api.Context = c
	return api
}

// MakeOrm 设置orm
func (api *Api) MakeOrm() *Api {
	api.Orm = runtime.RunTime.Orm()
	return api
}

func (api *Api) MakeService(service *service.Service) *Api {
	service.Orm = api.Orm
	return api
}

// Bind 参数校验
func (api *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		bindings = constructor.GetBindingForGin(d)
	}
	for i := range bindings {
		if bindings[i] == nil {
			err = api.Context.ShouldBindUri(d)
		} else {
			err = api.Context.ShouldBindWith(d, bindings[i])
		}
		if err != nil && err.Error() == "EOF" {
			log.Warn("request body is not present anymore. ")
			err = nil
			continue
		}
		if err != nil {
			api.AddError(errors.WithCode(code.ErrBind, err.Error()))
			break
		}
	}
	//vd.SetErrorFactory(func(failPath, msg string) error {
	//	return fmt.Errorf(`"validation failed: %s %s"`, failPath, msg)
	//})
	if err1 := vd.Validate(d); err1 != nil {
		api.AddError(errors.WithCode(code.ErrValidation, err1.Error()))
	}
	return api
}
