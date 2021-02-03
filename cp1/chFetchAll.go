package cp1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s : %v", url, err)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s ", secs, nbytes, url)
}

func ChFetchAll() {
	start := time.Now()
	ch := make(chan string)
	url_list := []string{"http://pypi.i.brainpp.cn/liveness/dev/%2Bf/167/7dcc9d07f31e1/lovelive_sdk-1.76-py3-none-any.whl", "http://pypi.i.brainpp.cn/liveness/dev/%2Bf/167/7dcc9d07f31e1/lovelive_sdk-1.76-py3-none-any.whl"}
	for _, url := range url_list {
		go fetch(url, ch)
	}
	for range url_list {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
