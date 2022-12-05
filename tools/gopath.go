package tools

import (
	"path/filepath"
	"runtime"
)

/*
GetRootPath 搜索项目的文件根目录, 并和 myPath 拼接起来
*/
func GetRootPath(myPath string) string {
	_, fileName, _, ok := runtime.Caller(0)
	if !ok {
		panic("Something wrong with getting root path")
	}
	absPath, err := filepath.Abs(fileName)
	rootPath := filepath.Dir(filepath.Dir(absPath))
	if err != nil {
		panic(any(err))
	}
	return filepath.Join(rootPath, myPath)
}

/*
GetDeployPath 获取启动项目的当前目录, 并和 myPath 拼接起来
*/
func GetDeployPath(myPath string) string {
	absPath, err := filepath.Abs("")
	if err != nil {
		panic(any(err))
	}
	return filepath.Join(absPath, myPath)
}
