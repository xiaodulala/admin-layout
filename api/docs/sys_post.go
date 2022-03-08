package docs

import "github.com/xiaodulala/admin-layout/internal/application/sys/service/dto"

// swagger:route GET /api/v1/sys/post  系统岗位管理 sysPostGetPage
//
// 分页查询岗位
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysPostGetPage
type sysPostGetPageWrapper struct {
	// in:query
	PostName string `json:"postName"`
	// in:query
	PostCode string `json:"postCode"`
	// in:query
	PostId string `json:"postId"`
	// in:query
	Status string `json:"status"`
}

// swagger:route GET /api/v1/sys/post/{postId}  系统岗位管理 sysPostGet
//
// 查询详情
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysPostGet
type sysPostGetWrapper struct {
	// in:path
	PostId string `json:"postId"`
}

// swagger:route POST /api/v1/sys/post  系统岗位管理 sysPostInsert
//
// 创建岗位
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysPostInsert
type sysPostInsertWrapper struct {
	// in:body
	Body dto.SysPostInsertReq
}

// swagger:route PUT /api/v1/sys/post  系统岗位管理 sysPostUpdate
//
// 创建岗位
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysPostUpdate
type sysPostUpdateWrapper struct {
	// in:body
	Body dto.SysPostUpdateReq
}

// swagger:route DELETE /api/v1/sys/post  系统岗位管理 sysPostDelete
//
// 创建岗位
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:parameters sysPostDelete
type sysPostDeleteWrapper struct {
	// in:body
	Body dto.SysMenuDeleteReq
}
