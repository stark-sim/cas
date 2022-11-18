package tools

import (
	"path/filepath"
)

/*
	搜索包含项目的 go.mod 文件的目录, 并和 myPath 拼接起来
*/
func GetDeployPath(myPath string) string {
	absPath, err := filepath.Abs("")
	if err != nil {
		panic(any(err))
	}
	return filepath.Join(absPath, myPath)
}

