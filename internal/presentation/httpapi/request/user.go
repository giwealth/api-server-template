package request

import "fmt"

type (
	// User 用户请求数据
	User struct {
		ID      int64  `form:"id" json:"id"`           // 用户ID
		Name    string `form:"name" json:"name"`       // 用户名称
		Age     int    `form:"age" json:"age"`         // 年龄
		Address string `form:"address" json:"address"` // 地址
	}
)

// Parse 复杂请求参数验证
func (r *User) Parse() error {
	if r.Age < 10 {
		return fmt.Errorf("用户年龄不能小于10")
	}
	return nil
}
