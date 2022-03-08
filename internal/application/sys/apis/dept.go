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
	"strconv"
)

type SysDept struct {
	api.Api
}

func (e SysDept) GetPage(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	list := make([]models.SysDept, 0)
	list, err = s.SetDeptPage(&req)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, list)
}

func (e SysDept) Get(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	var object models.SysDept

	err = s.Get(&req, &object)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	httpresponse.WriteSuccessResponse(c, object)
}

func (e SysDept) Insert(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptInsertReq{}
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
	httpresponse.WriteSuccessResponse(c, gin.H{"deptId": req.GetId()})
}

func (e SysDept) Update(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
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
	httpresponse.WriteSuccessResponse(c, gin.H{"deptId": req.GetId()})
}

func (e SysDept) Delete(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
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
	httpresponse.WriteSuccessResponse(c, gin.H{"deptId": req.GetId()})
}

// Get2Tree 用户管理 左侧部门树
func (e SysDept) Get2Tree(c *gin.Context) {
	s := service.SysDept{}
	req := dto.SysDeptGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	list := make([]dto.DeptLabel, 0)
	list, err = s.SetDeptTree(&req)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, list)
}

// GetDeptTreeRoleSelect TODO: 此接口需要调整不应该将list和选中放在一起
func (e SysDept) GetDeptTreeRoleSelect(c *gin.Context) {
	s := service.SysDept{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	id, err := strconv.Atoi(c.Param("roleId"))
	result, err := s.SetDeptLabel()
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	menuIds := make([]int, 0)
	if id != 0 {
		menuIds, err = s.GetWithRoleId(id)
		if err != nil {
			httpresponse.WriteErrResponse(c, err, nil)
			return
		}
	}
	httpresponse.WriteSuccessResponse(c, gin.H{
		"depts":       result,
		"checkedKeys": menuIds,
	})
}
