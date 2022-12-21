package content

import (
	"blog/fs"
	"embed"
)

var (
	//go:embed *.html
	htmlFiles embed.FS

	All = fs.NewFS(htmlFiles, "content")
)
