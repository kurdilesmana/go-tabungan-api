package logging

import "runtime"

func GetCaller() (file, function string, line int) {
	pc, file, line, _ := runtime.Caller(3)
	function = runtime.FuncForPC(pc).Name()
	return
}
