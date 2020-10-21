package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"fmt"
	"strings"
)

func fetchFile(params_list []string){
	for _, url := range params_list[1:]{
		if !strings.HasPrefix(url, "http://"){
			url = strings.Join([]string{"http://", url}, "")
		}
		//fmt.Println(url, "----")
		resp, err := http.Get(url)
		if err != nil{
			fmt.Printf("%s \n", err)
		}
		fmt.Println(resp.Status)
		out, err := os.Create("./tmp.log")
		if err != nil{
			fmt.Printf("%s \n", err)
		}
		n, err := io.Copy(out, resp.Body)
		if err != nil{
			fmt.Printf("%s \n", err)
		}
		resp.Body.Close()
		fmt.Printf("----- %s ---%d\n", out.Name(), n)
	}
}

func fetchDemo(paramList []string){
	for _, url := range paramList[1:]{
		if !strings.HasPrefix(url, "http://"){
			url = strings.Join([]string{"http://", url}, "")
		}
		resp, err := http.Get(url)
		retStatus := resp.Status
		if err != nil{
			fmt.Printf("%s \n", err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil{
			fmt.Println("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s \n status: %s \n", b, retStatus)
	}
}
func main(){
	fetchDemo(os.Args)
}