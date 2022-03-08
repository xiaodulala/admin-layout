package docs

import "github.com/xiaodulala/admin-layout/internal/application/sys/service/dto"

// swagger:route GET /api/v1/sys/menu  系统菜单管理 sysMenuGetPage
//
// 分页查询菜单列表数据
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysMenuGetPage
type sysMenuGetPageWrapper struct {
	// in:query
	MenuName string `json:"menuName"`
}

// swagger:route GET /api/v1/sys/menu/{id}  系统菜单管理 sysMenuGet
//
// 根据id查询菜单详情
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysMenuGet
type sysMenuGetWrapper struct {
	// in:path
	Id string `json:"id"`
}

// swagger:route POST /api/v1/sys/menu  系统菜单管理 sysMenuInsert
//
// 创建菜单
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysMenuInsert
type sysMenuInsertWrapper struct {
	// in:body
	Body dto.SysMenuInsertReq
}

// swagger:route PUT /api/v1/sys/menu  系统菜单管理 sysMenuUpdate
//
// 更新菜单
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysMenuUpdate
type sysMenuUpdateWrapper struct {
	// in:body
	Body dto.SysMenuUpdateReq
}

// swagger:route DELETE /api/v1/sys/menu  系统菜单管理 sysMenuDelete
//
// 删除菜单
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysMenuDelete
type sysMenuDeleteWrapper struct {
	// in:body
	Body dto.SysMenuDeleteReq
}

// swagger:route GET /api/v1/sys/menu/menu-role  系统菜单管理 sysMenuRole
//
// 根据登录角色名称获取菜单列表数据(左菜单使用)
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysMenuRole

// swagger:route GET /api/v1/sys/menu/menu-tree/{roleId}  系统菜单管理 sysMenuTreeSelect
//
// 根据角色ID查询菜单下拉树结构  角色修改使用的菜单列表
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysMenuTreeSelect
type sysMenuTreeSelectWrapper struct {
	// in:path
	RoleId int `json:"roleId"`
}
