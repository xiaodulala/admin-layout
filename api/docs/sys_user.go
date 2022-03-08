package docs

import (
	"github.com/xiaodulala/admin-layout/internal/application/sys/service/dto"
)

// swagger:route GET /api/v1/sys/user  系统用户管理 sysUserGetPage
//
// 分页查询用户
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysUserGetPage
type sysUserGetPageWrapper struct {
	// in:query
	Username string `json:"username"`
}

// swagger:route GET /api/v1/sys/user/{userId}  系统用户管理 sysUserGet
//
// 查询用户
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysUserGet
type sysUserGetWrapper struct {
	// in:path
	UserId string `json:"userId"`
}

// swagger:route POST /api/v1/sys/user  系统用户管理 sysUserInsert
//
// 创建用户
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysUserInsert
type sysUserInsertWrapper struct {
	// in:body
	Body dto.SysUserInsertReq
}

// swagger:route PUT /api/v1/sys/user  系统用户管理 sysUserUpdate
//
// 更新用户
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysUserUpdate
type sysUserUpdateWrapper struct {
	// in:body
	Body dto.SysUserUpdateReq
}

// swagger:route DELETE /api/v1/sys/user/{userId}  系统用户管理 sysUserDelete
//
// 删除用户
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysUserDelete
type sysUserDeleteWrapper struct {
	// in:path
	UserId string `json:"userId"`
}

// swagger:route POST /api/v1/sys/user/avatar  系统用户管理 sysUserAvatar
//
// 修改用户头像
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysUserAvatar
type sysUserAvatarWrapper struct {
}

// swagger:route PUT /api/v1/sys/user/status  系统用户管理 sysUserStatus
//
// 修改用户状态
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysUserStatus
type sysUserStatusWrapper struct {
	// in:body
	Body dto.UpdateSysUserStatusReq
}

// swagger:route PUT /api/v1/sys/user/pwd/reset  系统用户管理 sysUserPwdReset
//
// 重置密码
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysUserPwdReset
type sysUserPwdResetWrapper struct {
	// in:body
	Body dto.ResetSysUserPwdReq
}

// swagger:route PUT /api/v1/sys/user/pwd/set  系统用户管理 sysUserPwdSet
//
// 修改密码
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysUserPwdSet
type sysUserPwdSetWrapper struct {
	// in:body
	Body dto.PassWord
}

// swagger:route GET /api/v1/sys/user/profile  系统用户管理
//
// 获取用户个人中心
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
