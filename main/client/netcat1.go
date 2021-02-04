package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// runOneClient("localhost:8020")
	// runOneClient("localhost:8030")
	runOneClient("localhost:8040")
}

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
