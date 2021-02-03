package cp1

import (
	"bufio"
	"fmt"
	"os"
)

func dupFiles(f *os.File, counts map[string][]string) {
	input := bufio.NewScanner(f)
	//f.Name()
	for input.Scan() {
		bFlag := false
		// !!! 这里 listCount := counts[input.Text()] ：slice 传值是引用，但只对 slice 原始长度有效
		for _, fileName := range counts[input.Text()] {
			if fileName == f.Name() {
				bFlag = true
			}
		}
		if !bFlag {
			counts[input.Text()] = append(counts[input.Text()], f.Name())
			fmt.Printf("hello")
		}
	}
}
