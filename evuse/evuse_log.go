package evuse

import (
	"fmt"
	. "show_module/lib"
	"strconv"
)

// LogTest 测试log : Info, Warnning, Error.
func LogTest() {
	SetLogger()
	Info.Printf("line %v, this log from %s 's info log ", 48, "usr_log")
	Warning.Printf("line %v, this log from %s 's warning log ", 48, "usr_log")
}

func ItoaUse() {
	join_str := ":" + strconv.Itoa(100)
	fmt.Println(join_str)
}
