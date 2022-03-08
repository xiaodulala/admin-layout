package docs

import "github.com/xiaodulala/admin-layout/internal/application/sys/service/dto"

// swagger:route GET /api/v1/sys/dept  部门管理 sysDeptGetPage
//
// 分页查询部门
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysDeptGetPage
type sysDeptGetPageWrapper struct {
	// in:query
	DeptName string `json:"deptName"`
	// in:query
	DeptId string `json:"deptId"`
	// in:query
	Position string `json:"position"`
}

// swagger:route GET /api/v1/sys/dept/{deptId}  部门管理 sysDeptGet
//
// 获取部门数据
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysDeptGet
type sysDeptGetWrapper struct {
	// in:path
	DeptId string `json:"deptId"`
}

// swagger:route POST /api/v1/sys/dept  部门管理 sysDeptInsert
//
// 添加部门
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysDeptInsert
type sysDeptInsertWrapper struct {
	// in:body
	Body dto.SysDeptInsertReq
}

// swagger:route PUT /api/v1/sys/dept  部门管理 sysDeptUpdate
//
// 修改部门
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysDeptUpdate
type sysDeptUpdateWrapper struct {
	// in:path
	Id int `json:"id"`
	// in:body
	Body dto.SysRoleUpdateReq
}

// swagger:route DELETE /api/v1/sys/dept  部门管理 sysDeptDelete
//
// 删除部门
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysDeptDelete
type sysDeptDeleteWrapper struct {
	// in:body
	Body dto.SysRoleDeleteReq
}

// swagger:route GET /api/v1/sys/dept/dept-tree  部门管理 sysDeptTree
//
// 用户管理左侧部门树
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters sysDeptTree
