package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	//http.HandleFunc("/favicon.ico", doNothing)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func doNothing(w http.ResponseWriter, r *http.Request){}

func handler(w http.ResponseWriter, r *http.Request){
	if r.URL.RequestURI() == "/favicon.ico"{
		fmt.Printf("handler view favicon ico %d \n", count)
	}else{
		fmt.Printf("handler not view favicon ico %d \n", count)
	}
	mu.Lock()
	count += 1
	fmt.Fprintf(w, "URL.Path = %q %d \n", r.URL.Path, count)
	mu.Unlock()
}

func counter(w http.ResponseWriter, r *http.Request){
	if r.URL.RequestURI() == "/favicon.ico"{
		fmt.Printf("counter view favicon ico %d \n", count)
	}else{
		fmt.Printf("counter not view favicon ico %d \n", count)
	}
	mu.Lock()
	count += 1
	fmt.Fprintf(w, "Count %d \n", count)
	mu.Unlock()
}