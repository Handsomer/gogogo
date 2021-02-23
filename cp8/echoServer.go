package cp8

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func StartEchoServer() {
	Listener, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := Listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleEchoConn(conn)
	}
}

// 获取连接的内容，当输入结束关闭该连接
func handleEchoConn(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		echo(conn, input.Text())
	}
	conn.Close()
}

// 向conn 写入内容
func echo(conn net.Conn, shout string) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(1 * time.Second)
	fmt.Fprintln(conn, "\t", shout)
	time.Sleep(1 * time.Second)
	fmt.Fprintln(conn, "\t", strings.ToLower(shout))
}
