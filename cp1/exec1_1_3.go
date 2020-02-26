package main

//比较低效程序 和 使用join的程序的耗时对比
import (
	"fmt"
	"os"
	"strings"
	"time"
)

func build_plus(){
	s, sep := "", ""
	for _, arg := range os.Args{
		s += sep + arg
		sep = " "
	}
}

func build_join(){
	b := ""
	b += strings.Join(os.Args, " ")
}

// MacBook-Pro:cp1 zzz$ go run exec1_1_3.go   aaa bbb ccc
// 0.0000072360s 
// MacBook-Pro:cp1 zzz$ go run exec1_1_3.go   aaa bbb ccc
// 0.0000031730s 
// 时间上差一倍

func main(){
	start := time.Now()
	// build_plus()
	build_join()
	fmt.Printf("%.10f s \n", time.Since(start).Seconds())
}