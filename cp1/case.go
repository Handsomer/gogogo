package main

import (
	"bufio"
	"fmt"
	"os"
)

// 计算输入数据出现的次数
func cal_nums(){
	// 定义一个key value对集合，
	counts := make(map[string] int)
	// 定义一个扫描器
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		if input.Text() == "end" {break}
		counts[input.Text()] ++
	}
	for line, n := range counts{
		if n > 1{
			fmt.Printf("%d\t%s\n", n , line)
		}
	}
}

func main(){
	cal_nums()
}
