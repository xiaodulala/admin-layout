// Code generated by "codegen --type=int"; DO NOT EDIT.

package code

import "github.com/xiaodulala/admin-layout/pkg/errors" // init register error codes defines in this source code to `github.com/xiaodulala/admin-layout/pkg/errors`
func init() {
	errors.RegisterErrCoder(ErrUserNotFound, 404, "User not found")
	errors.RegisterErrCoder(ErrUserAlreadyExist, 400, "User already exist")
	errors.RegisterErrCoder(ErrUserPasswordWrong, 400, "Password wrong")
	errors.RegisterErrCoder(ErrRoleNotFound, 404, "Role not found")
	errors.RegisterErrCoder(ErrRoleCasbin, 500, "Role casbin wrong")
	errors.RegisterErrCoder(ErrSuccess, 200, "OK")
	errors.RegisterErrCoder(ErrUnknown, 500, "Internal server error")
	errors.RegisterErrCoder(ErrBind, 400, "Error occurred while binding the request body to the struct")
	errors.RegisterErrCoder(ErrValidation, 400, "Validation failed")
	errors.RegisterErrCoder(ErrTokenInvalid, 401, "Token invalid")
	errors.RegisterErrCoder(ErrPageNotFound, 404, "Page not found")
	errors.RegisterErrCoder(ErrDatabase, 500, "Database error")
	errors.RegisterErrCoder(ErrEncrypt, 401, "Error occurred while encrypting the user password")
	errors.RegisterErrCoder(ErrSignatureInvalid, 401, "Signature is invalid")
	errors.RegisterErrCoder(ErrExpired, 401, "Token expired")
	errors.RegisterErrCoder(ErrInvalidAuthHeader, 401, "Invalid authorization header")
	errors.RegisterErrCoder(ErrMissingHeader, 401, "The `Authorization` header was empty")
	errors.RegisterErrCoder(ErrorExpired, 401, "Token expired")
	errors.RegisterErrCoder(ErrPasswordIncorrect, 401, "Password was incorrect")
	errors.RegisterErrCoder(ErrPermissionDenied, 403, "Permission denied")
	errors.RegisterErrCoder(ErrDataPermissionDenied, 403, "Data Permission denied")
	errors.RegisterErrCoder(ErrEncodingFailed, 500, "Encoding failed due to an error with the data")
	errors.RegisterErrCoder(ErrDecodingFailed, 500, "Decoding failed due to an error with the data")
	errors.RegisterErrCoder(ErrInvalidJSON, 500, "Data is not valid JSON")
	errors.RegisterErrCoder(ErrEncodingJSON, 500, "JSON data could not be encoded")
	errors.RegisterErrCoder(ErrDecodingJSON, 500, "JSON data could not be decoded")
	errors.RegisterErrCoder(ErrInvalidYaml, 500, "Data is not valid Yaml")
	errors.RegisterErrCoder(ErrEncodingYaml, 500, "Yaml data could not be encoded")
	errors.RegisterErrCoder(ErrDecodingYaml, 500, "Yaml data could not be decoded")
}
