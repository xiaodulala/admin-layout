package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/xiaodulala/admin-layout/internal/application/common/api"
	common "github.com/xiaodulala/admin-layout/internal/application/common/models"
	"github.com/xiaodulala/admin-layout/internal/application/sys/models"
	"github.com/xiaodulala/admin-layout/internal/application/sys/service"
	"github.com/xiaodulala/admin-layout/internal/application/sys/service/dto"
	"github.com/xiaodulala/admin-layout/internal/pkg/runtime"
	"github.com/xiaodulala/admin-layout/pkg/httpresponse"
	"github.com/xiaodulala/admin-layout/pkg/middleware"
)

type SysRole struct {
	api.Api
}

func (e SysRole) GetPage(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	list := make([]models.SysRole, 0)
	var count int64

	err = s.GetPage(&req, &list, &count)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	pageRes := common.Page{
		Count:     int(count),
		PageIndex: req.GetPageIndex(),
		PageSize:  req.GetPageSize(),
		List:      list,
	}
	httpresponse.WriteSuccessResponse(c, pageRes)
}

func (e SysRole) Get(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	var object models.SysRole

	err = s.Get(&req, &object)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	httpresponse.WriteSuccessResponse(c, object)
}

func (e SysRole) Insert(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleInsertReq{}
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
	req.CreateBy = middleware.GetUserId(c)
	if req.Status == "" {
		req.Status = "2"
	}
	cb := runtime.RunTime.Casbin()
	err = s.Insert(&req, cb)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, req.GetId())
}

func (e SysRole) Update(c *gin.Context) {
	s := service.SysRole{}
	req := dto.SysRoleUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	cb := runtime.RunTime.Casbin()

	req.SetUpdateBy(middleware.GetUserId(c))

	err = s.Update(&req, cb)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	httpresponse.WriteSuccessResponse(c, req.GetId())
}

func (e SysRole) Delete(c *gin.Context) {
	s := new(service.SysRole)
	req := dto.SysRoleDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	err = s.Remove(&req)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	httpresponse.WriteSuccessResponse(c, req.GetId())
}

func (e SysRole) Update2Status(c *gin.Context) {
	s := service.SysRole{}
	req := dto.UpdateStatusReq{}
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
	err = s.UpdateStatus(&req)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, req.GetId())
}

func (e SysRole) Update2DataScope(c *gin.Context) {
	s := service.SysRole{}
	req := dto.RoleDataScopeReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	data := &models.SysRole{
		RoleId:    req.RoleId,
		DataScope: req.DataScope,
		DeptIds:   req.DeptIds,
	}
	data.UpdateBy = middleware.GetUserId(c)
	err = s.UpdateDataScope(&req)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, nil)
}
