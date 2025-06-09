package static

import (
	"embed"
)

//go:embed web/index.html
var index embed.FS

// GetFileSystem returns an http.FileSystem that serves the embedded static files
func GetIndex() []byte {
	fsys, err := index.ReadFile("web/index.html")
	if err != nil {
		panic(err)
	}
	return fsys
}
