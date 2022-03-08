package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/xiaodulala/admin-layout/internal/application/common/actions"
	"github.com/xiaodulala/admin-layout/internal/application/common/api"
	common "github.com/xiaodulala/admin-layout/internal/application/common/models"
	"github.com/xiaodulala/admin-layout/internal/application/sys/models"
	"github.com/xiaodulala/admin-layout/internal/application/sys/service"
	"github.com/xiaodulala/admin-layout/internal/application/sys/service/dto"
	"github.com/xiaodulala/admin-layout/pkg/httpresponse"
	"github.com/xiaodulala/admin-layout/pkg/log"
	"github.com/xiaodulala/admin-layout/pkg/middleware"
)

type SysUser struct {
	api.Api
}

func (su SysUser) GetPage(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserGetPageReq{}
	err := su.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	list := make([]models.SysUser, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
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

func (su SysUser) Get(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserById{}
	err := su.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	var object models.SysUser
	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	httpresponse.WriteSuccessResponse(c, object)
}

func (su SysUser) Insert(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserInsertReq{}
	err := su.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	// 设置创建人
	req.SetCreateBy(c.GetInt("userId"))
	err = s.Insert(&req)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	httpresponse.WriteSuccessResponse(c, req.GetId())
}

func (su SysUser) Update(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserUpdateReq{}
	err := su.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	req.SetUpdateBy(middleware.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, req.GetId())
}

func (su SysUser) Delete(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserById{}
	err := su.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	// 设置编辑人
	req.SetUpdateBy(middleware.GetUserId(c))

	// 数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, req.GetId())
}

func (su SysUser) InsetAvatar(c *gin.Context) {
	s := service.SysUser{}
	req := dto.UpdateSysUserAvatarReq{}
	err := su.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	// 数据权限检查
	p := actions.GetPermissionFromContext(c)
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	guid := uuid.New().String()
	filPath := "static/uploadfile/" + guid + ".jpg"
	for _, file := range files {
		log.Debugf("upload avatar file: %s", file.Filename)
		// 上传文件至指定目录
		err = c.SaveUploadedFile(file, filPath)
		if err != nil {
			httpresponse.WriteErrResponse(c, err, nil)
		}
	}
	req.UserId = p.UserId
	req.Avatar = "/" + filPath

	err = s.UpdateAvatar(&req, p)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, req.GetId())
}

func (su SysUser) UpdateStatus(c *gin.Context) {
	s := service.SysUser{}
	req := dto.UpdateSysUserStatusReq{}
	err := su.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	req.SetUpdateBy(middleware.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.UpdateStatus(&req, p)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, req.GetId())
}

func (su SysUser) ResetPwd(c *gin.Context) {
	s := service.SysUser{}
	req := dto.ResetSysUserPwdReq{}
	err := su.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	req.SetUpdateBy(middleware.GetUserId(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.ResetPwd(&req, p)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, req.GetId())
}

func (su SysUser) UpdatePwd(c *gin.Context) {
	s := service.SysUser{}
	req := dto.PassWord{}
	err := su.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	// 数据权限检查
	p := actions.GetPermissionFromContext(c)

	err = s.UpdatePwd(middleware.GetUserId(c), req.OldPassword, req.NewPassword, p)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}
	httpresponse.WriteSuccessResponse(c, nil)
}

func (su SysUser) GetProfile(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserById{}
	err := su.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	req.Id = middleware.GetUserId(c)

	sysUser := models.SysUser{}
	roles := make([]models.SysRole, 0)
	posts := make([]models.SysPost, 0)
	err = s.GetProfile(&req, &sysUser, &roles, &posts)
	if err != nil {
		httpresponse.WriteErrResponse(c, err, nil)
		return
	}

	httpresponse.WriteSuccessResponse(c, gin.H{
		"user":  sysUser,
		"roles": roles,
		"posts": posts,
	})
}
