package admin

import "embed"

//go:embed css js img index.html favicon.ico
var Static embed.FS
