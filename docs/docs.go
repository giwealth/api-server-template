package docs

import (
	"embed" // 导入内嵌静态文件包
)

// StaticFS 静态文件
//
//go:embed *
var StaticFS embed.FS
