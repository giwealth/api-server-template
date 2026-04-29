package docs

import (
	"embed" // 导入内嵌静态文件包
)

// StaticFS embeds the docs static files.
//go:embed *
var StaticFS embed.FS