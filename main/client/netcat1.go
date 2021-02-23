package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	runEchoServer()
}

// 向回声服务器发送请求：
//		主进程：stdin 阻塞式写入 conn
//		协程：  输出到 stdout
func runEchoServer() {
	conn, err := net.Dial("tcp", "localhost:9001")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

// runOneClient("localhost:8020")
// runOneClient("localhost:8030")
// runOneClient("localhost:8040")
func runOneClient(host string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
