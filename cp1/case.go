package cp1

import (
	"bufio"
	"fmt"
	"os"
)

// 计算输入数据出现的次数
func cal_nums() {
	// 定义一个key value对集合，
	counts := make(map[string]int)
	// 定义一个扫描器
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func DupFiles() {
	counts := make(map[string]int)
	fileList := []string{"/home/sunwei03/go/src/gogogo/cp1/f_case.txt", "/home/sunwei03/go/src/gogogo/cp1/f2_case.txt"}
	if len(fileList) == 0 {
		countlines(os.Stdin, counts)
		out_counts(counts, os.Args[0])
	} else {
		for _, arg := range fileList {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v /n", err)
			}
			countlines(f, counts)
			f.Close()
			out_counts(counts, arg)
			// 有意思的是，Go语言中并没有为 map 提供任何清空所有元素的函数、方法，
			// 清空 map 的唯一办法就是重新 make 一个新的 map，
			// 不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。
			counts = make(map[string]int)
		}
	}
}

func out_counts(counts map[string]int, f_name string) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d \t %s \t %s \n", n, line, f_name)
		}
	}
}

func countlines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
