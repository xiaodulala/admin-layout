// Package docs admin-layout  Server API.
//
//     Schemes: http, https
//     Host: localhost:9099
//     BasePath: /
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//    SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: Authorization
//          in: header
// swagger:meta
package docs

import "github.com/xiaodulala/admin-layout/pkg/httpresponse"

// swagger:response errResponse
type errResponseWrapper struct {
	// in:body
	Body httpresponse.ErrResponse
}

// swagger:response successResponse
type successResponseWrapper struct {
	// in:body
	Body httpresponse.Response
}
