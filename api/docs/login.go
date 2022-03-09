package docs

import "github.com/xiaodulala/admin-layout/internal/pkg/ginauth"

// swagger:route POST /login  认证 login
//
// 用户登录
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
// swagger:parameters login
type loginWrapper struct {
	// in:body
	Body ginauth.Login
}

// swagger:route GET /logout  认证 logout
//
// 用户登录
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//

// swagger:route GET /refresh-token  认证 refresh
//
// 用户登录
//
//     Responses:
//       success: successResponse
//       error: errResponse
//
//
