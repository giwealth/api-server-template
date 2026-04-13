package request

// Empty 空请求
type Empty struct{}

// Query 查询
type Query struct {
	Keyword string `form:"keyword" json:"keyword"`                                             // [可选 默认空字符串] 搜索关键字
	Page    int    `form:"page" json:"page" binding:"required" error:"query page required"`    // 页码，从1开始
	Limit   int    `form:"limit" json:"limit" binding:"required" error:"query limit required"` // 每页显示数量
}
