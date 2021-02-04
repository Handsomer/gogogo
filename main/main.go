package main

import (
	"gogogo/cp1"
	"gogogo/cp5"
	"gogogo/cp8"
)

func showAllPackage() {
	// go 文件操作
	cp1.DupFiles()
	// fetch 网络请求：获得所有内容
	cp1.ChFetchAll()
	cp5.AddUse()
	// 字符串拼接 join
	cp1.BuildJoin()
	// 字符串拼接 +
	cp1.BuildPlus()
	// Unicode 编码长度验证
	cp8.StringUTF8Demo()
	// cp8.StartServer()
	// cp8.RunOneServer("China", "localhost:8030")
	// map 数据结构的遍历
	cp8.MapRangeAndGetUse()
	// 使用 sync map
	cp8.SyncMapUse()
}

func main() {
	// cp8.RunOneServer("China", "localhost:8003")
	// cp8.RunOneServer("America", "localhost:8004")
	showAllPackage()
}
