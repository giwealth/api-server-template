package request

// Empty 空请求
type Empty struct{}

// Query 查询
type Query struct {
	// [可选 默认空字符串] 搜索关键字
	Keyword string `form:"keyword" json:"keyword"`
	// 页码，从1开始
	Page int `form:"page" json:"page" binding:"required" error:"query page required"`
	// 每页显示数量
	Limit int `form:"limit" json:"limit" binding:"required" error:"query limit required"`
}
