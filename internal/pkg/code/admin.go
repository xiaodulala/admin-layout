package code

const (
	// ErrUserNotFound - 404: User not found.
	ErrUserNotFound int = iota + 110001

	// ErrUserAlreadyExist - 400: User already exist.
	ErrUserAlreadyExist

	// ErrUserPasswordWrong - 400: Password wrong.
	ErrUserPasswordWrong
)

// 角色
const (
	// ErrRoleNotFound - 404: Role not found.
	ErrRoleNotFound int = iota + 110101

	// ErrRoleCasbin - 500: Role casbin wrong.
	ErrRoleCasbin
)
