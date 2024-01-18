package helper

import (
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func GetCallerInfo(skipCaller int) (fileName string, line string, funcName string) {
	pc := make([]uintptr, 1)
	runtime.Callers(skipCaller+1, pc)
	funcInfo := runtime.FuncForPC(pc[0])
	file, lineInt := funcInfo.FileLine(pc[0])
	dir, fileBase := filepath.Split(file)
	dirBase := filepath.Base(dir)
	name := formatFuncName(funcInfo.Name())
	return dirBase + "." + fileBase, strconv.Itoa(lineInt), name
}

func formatFuncName(name string) string {
	name = path.Base(name)
	split := strings.Split(name, ".")
	return split[len(split)-1]
}
