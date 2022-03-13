package admin

import "embed"

//go:embed css js index.html favicon.ico
var Static embed.FS
