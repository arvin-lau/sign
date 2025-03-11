package client

import (
	"runtime"
	"strings"
)

func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	res := strings.Split(f.Name(), ".")
	return res[len(res)-1]
}
