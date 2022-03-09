package docs

import "github.com/xiaodulala/admin-layout/internal/application/sys/service/dto"

// swagger:route GET /api/v1/sys/role  系统角色管理 sysRoleGetPage
//
// 分页查询角色
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysRoleGetPage
type sysRoleGetPageWrapper struct {
	// in:query
	RoleName string `json:"roleName"`
	// in:query
	Status string `json:"status"`
	// in:query
	RoleKey string `json:"roleKey"`
	// in:query
	PageSize int `json:"pageSize"`
	// in:query
	PageIndex int `json:"pageIndex"`
}

// swagger:route GET /api/v1/sys/role/{roleId}  系统角色管理 sysRoleGet
//
// 查询角色信息
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysRoleGet
type sysRoleGetWrapper struct {
	// in:path
	RoleId string `json:"roleId"`
}

// swagger:route POST /api/v1/sys/role  系统角色管理 sysRoleInsert
//
// 创建角色
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysRoleInsert
type sysRoleInsertWrapper struct {
	// in:body
	// example: {"roleName":"财务总监","roleKey":"account","roleSort":1,"status":"2","menuIds":[3,43,44,45,46],"deptIds":[],"sysMenu":[]}
	Body dto.SysRoleInsertReq
}

// swagger:route PUT /api/v1/sys/role/{id}  系统角色管理 sysRoleUpdate
//
// 更新角色
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysRoleUpdate
type sysRoleUpdateWrapper struct {
	// in:path
	Id int `json:"id"`
	// in:body
	Body dto.SysRoleUpdateReq
}

// swagger:route DELETE /api/v1/sys/role  系统角色管理 sysRoleDelete
//
// 删除角色
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysRoleDelete
type sysRoleDeleteWrapper struct {
	// in:body
	Body dto.SysRoleDeleteReq
}

// swagger:route PUT /api/v1/sys/role/role-status  系统角色管理 sysRoleUpdateStatue
//
// 更新角色状态
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysRoleUpdateStatue
type sysRoleUpdateStatueWrapper struct {
	// in:body
	Body dto.UpdateStatusReq
}

// swagger:route PUT /api/v1/sys/role/role-scope  系统角色管理 sysRoleUpdateDataScope
//
// 更新角色数据权限
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysRoleUpdateDataScope
type sysRoleUpdateDataScope struct {
	// in:body
	Body dto.RoleDataScopeReq
}
