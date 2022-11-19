package configs

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"
)

func InitLogger() (err error) {
	// 不打进文件，打进控制台用 grafana 来自动管理最佳
	logrus.SetFormatter(formatter(true))
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugln("[Init] init logger done")
	return err
}

// 自定义日志格式化，将日志打进 console
func formatter(isConsole bool) *nested.Formatter {
	fmtter := &nested.Formatter{
		FieldsOrder:           nil,
		TimestampFormat:       "2006-01-02 15:04:05",
		HideKeys:              false,
		NoColors:              false,
		NoFieldsColors:        false,
		NoFieldsSpace:         false,
		ShowFullLevel:         true,
		NoUppercaseLevel:      false,
		TrimMessages:          false,
		CallerFirst:           true,
		CustomCallerFormatter: func(frame *runtime.Frame) string {
			funcInfo := runtime.FuncForPC(frame.PC)
			if funcInfo == nil {
				return "error of runtime.FuncForPC"
			}
			fullPath, line := funcInfo.FileLine(frame.PC)
			funcSlice := strings.Split(funcInfo.Name(), ".")
			funcName := funcSlice[len(funcSlice) - 1]
			return fmt.Sprintf(" [%v]-[%v]-[%v]", filepath.Base(fullPath), funcName, line)
		},
	}
	// 打进控制台需要颜色
	fmtter.NoColors = !isConsole
	return fmtter
}