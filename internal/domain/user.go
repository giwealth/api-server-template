package domain

import "errors"

var (
	// ErrUserNotFound not found
	ErrUserNotFound = errors.New("User not found")
)

type (
	// User 用户
	User struct {
		// 用户ID
		ID int64 `json:"id"`
		// 用户名称
		Name string `json:"name"`
		// 年龄
		Age int `json:"age"`
		// 地址
		Address string `json:"address"`
	}
)
