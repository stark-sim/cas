package tools

import (
	"path"
	"path/filepath"
	"runtime"
)

/*
	搜索包含项目的 go.mod 文件的根目录, 并和 myPath 拼接起来
*/
func GetRootPath(myPath string) string {
	_, fileName, _, ok := runtime.Caller(0)
	if !ok {
		panic("Something wrong with getting root path")
	}
	rootPath := path.Dir(path.Dir(fileName))
	return filepath.Join(rootPath, myPath)
}
