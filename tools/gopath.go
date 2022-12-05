package tools

import (
	"path"
	"path/filepath"
	"runtime"
)

/*
GetRootPath 搜索项目的根目录, 并和 myPath 拼接起来
*/
func GetRootPath(myPath string) string {
	_, fileName, _, ok := runtime.Caller(0)
	if !ok {
		panic("Something wrong with getting root path")
	}
	rootPath := path.Dir(path.Dir(fileName))
	return filepath.Join(rootPath, myPath)
}

/*
GetDeployPath 获取部署时的根目录
*/
func GetDeployPath(myPath string) string {
	absPath, err := filepath.Abs("")
	if err != nil {
		panic(any(err))
	}
	return filepath.Join(absPath, myPath)
}
