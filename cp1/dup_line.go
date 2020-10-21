package main

import(
	"bufio"
	"fmt"
	"os"
)

func dupFiles(f *os.File, counts map[string][] string){
	input := bufio.NewScanner(f)
	//f.Name()
	for input.Scan(){
		bFlag := false
		for _, fileName := range counts[input.Text()]{
			if fileName == f.Name(){
				bFlag = true
			}
		}
		if !bFlag{
			counts[input.Text()] = append(counts[input.Text()], f.Name())
			fmt.Printf("hello")
		}
	}
}

func main() {
	counts := make(map[string][]string)
	for _, file := range os.Args[1:]{
		f, _ := os.Open(file)
		dupFiles(f, counts)
		f.Close()
	}
	for line, fileList := range counts{
		if len(fileList) > 0{
			//fmt.Println(line, int32(n))
			fmt.Printf("%s--- %s \n", line, fileList)
		}
	}
}