package cp1

//比较低效程序 和 使用join的程序的耗时对比
import (
	"fmt"
	"strings"
)

func BuildPlus() {
	s, sep := "", ""
	stringList := []string{"hello", "nice to meet you", "~"}
	for _, arg := range stringList {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func BuildJoin() {
	b := ""
	stringList := []string{"hello", "nice to meet you", "~"}
	b += strings.Join(stringList, " ")
	fmt.Println(b)
}

// MacBook-Pro:cp1 zzz$ go run exec1_1_3.go   aaa bbb ccc
// 0.0000072360s
// MacBook-Pro:cp1 zzz$ go run exec1_1_3.go   aaa bbb ccc
// 0.0000031730s
// 时间上差一倍

// func main(){
// 	start := time.Now()
// 	// build_plus()
// 	build_join()
// 	fmt.Printf("%.10f s \n", time.Since(start).Seconds())
// }
