package domain

import "errors"

var (
	// ErrAdministratorNotFound not found
	ErrAdministratorNotFound = errors.New("Administrator not found")
)

type (
	// Administrator 管理员
	Administrator struct {
		// 管理员ID
		ID int64 `json:"id"`
		// 管理员用户名
		Username string `json:"username"`
		// 密码
		Password string `json:"password"`
	}
)
