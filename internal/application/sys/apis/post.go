package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/xiaodulala/admin-layout/internal/application/common/api"
	common "github.com/xiaodulala/admin-layout/internal/application/common/models"
	"github.com/xiaodulala/admin-layout/internal/application/sys/models"
	"github.com/xiaodulala/admin-layout/internal/application/sys/service"
	"github.com/xiaodulala/admin-layout/internal/application/sys/service/dto"
	"github.com/xiaodulala/admin-layout/pkg/httpresponse"
	"github.com/xiaodulala/admin-layout/pkg/middleware"
)

type SysPost struct {
	api.Api
}

func (e SysPost) GetPage(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	list := make([]models.SysPost, 0)
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

func (e SysPost) Get(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	var object models.SysPost

	err = s.Get(&req, &object)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, object)
}

func (e SysPost) Insert(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostInsertReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	req.SetCreateBy(middleware.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, req.GetId())
}

func (e SysPost) Update(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostUpdateReq{}
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

func (e SysPost) Delete(c *gin.Context) {
	s := service.SysPost{}
	req := dto.SysPostDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	req.SetUpdateBy(middleware.GetUserId(c))
	err = s.Remove(&req)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, req.GetId())
}
