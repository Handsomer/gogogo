package main

import (
	"show_module/cp8"
)

func main() {
	//go run main/server/clock1.go China=localhost:8020 America=localhost:8030 unknown=localhost:8040
	cp8.StartEchoServer()
}
