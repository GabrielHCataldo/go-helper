package helper

import (
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func GetCallerInfo(skip int) (fileName string, line string, funcName string) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		pc, file, lineNo, _ = runtime.Caller(1)
	}
	funcInfo := runtime.FuncForPC(pc).Name()
	dir, fileBase := filepath.Split(file)
	dirBase := filepath.Base(dir)
	name := formatFuncName(funcInfo)
	return dirBase + "/" + fileBase, strconv.Itoa(MinInt(lineNo, 1)), name
}

func formatFuncName(name string) string {
	name = path.Base(name)
	split := strings.Split(name, ".")
	return split[len(split)-1]
}
