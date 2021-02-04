package main

import (
	"gogogo/cp1"
	"gogogo/cp5"
	"gogogo/cp8"
)

func showAllPackage() {
	cp1.DupFiles()
	cp1.ChFetchAll()
	cp5.AddUse()
	cp1.BuildJoin()
	cp1.BuildPlus()
	cp8.StringUTF8Demo()
	cp8.StartServer()
	cp8.RunOneServer("China", "localhost:8030")
	cp8.MapRangeAndGetUse()
}

func main() {
	// cp8.RunOneServer("China", "localhost:8003")
	// cp8.RunOneServer("America", "localhost:8004")
	cp8.StartServer()
}
