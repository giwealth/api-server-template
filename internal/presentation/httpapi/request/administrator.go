package request

type (
	// Administrator 管理员请求数据
	Administrator struct {
		// 管理员ID
		ID int64 `form:"id" uri:"id" json:"id"`
		// 管理员用户名
		Username string `form:"username" json:"username"`
		// 密码
		Password string `form:"password" json:"password"`
	}
)
