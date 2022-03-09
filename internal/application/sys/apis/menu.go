package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/xiaodulala/admin-layout/internal/application/common/api"
	"github.com/xiaodulala/admin-layout/internal/application/sys/models"
	"github.com/xiaodulala/admin-layout/internal/application/sys/service"
	"github.com/xiaodulala/admin-layout/internal/application/sys/service/dto"
	"github.com/xiaodulala/admin-layout/pkg/httpresponse"
	"github.com/xiaodulala/admin-layout/pkg/middleware"
)

type SysMenu struct {
	api.Api
}

func (e SysMenu) GetPage(c *gin.Context) {
	s := service.SysMenu{}
	req := dto.SysMenuGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	var list = make([]models.SysMenu, 0)
	err = s.GetPage(&req, &list)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, list)
}

func (e SysMenu) Get(c *gin.Context) {
	req := dto.SysMenuGetReq{}
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	var object = models.SysMenu{}

	err = s.Get(&req, &object)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, object)
}

func (e SysMenu) Insert(c *gin.Context) {
	req := dto.SysMenuInsertReq{}
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	// 设置创建人
	req.SetCreateBy(middleware.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, req.GetId())
}

func (e SysMenu) Update(c *gin.Context) {
	req := dto.SysMenuUpdateReq{}
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	req.SetUpdateBy(middleware.GetUserId(c))
	err = s.Update(&req)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, req.GetId())
}

func (e SysMenu) Delete(c *gin.Context) {
	control := new(dto.SysMenuDeleteReq)
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(control, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	err = s.Remove(control)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, control.GetId())
}

func (e SysMenu) GetMenuRole(c *gin.Context) {
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	result, err := s.SetMenuRole(middleware.GetRoleKey(c))

	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	httpresponse.WriteSuccessResponse(c, result)
}

func (e SysMenu) GetMenuTreeSelect(c *gin.Context) {
	m := service.SysMenu{}
	r := service.SysRole{}
	req := dto.SelectRole{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&m.Service).
		MakeService(&r.Service).
		Bind(&req, nil).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	result, err := m.SetLabel()
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	menuIds := make([]int, 0)
	if req.RoleId != 0 {
		menuIds, err = r.GetRoleMenuId(req.RoleId)
		if err != nil {
			httpresponse.WriteErrResponse(c, err, nil)
			return
		}
	}

	httpresponse.WriteSuccessResponse(c, gin.H{
		"menus":       result,
		"checkedKeys": menuIds,
	})
}
