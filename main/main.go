package main

import (
	"show_module/cp1"
	"show_module/cp5"
	"show_module/cp8"
	"show_module/evuse"
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
	// 使用 log 模块 实现输出结果
	evuse.LogTest()
	// 使用 itoa
	evuse.ItoaUse()
	// SelectGoroutine  chan 缓存队列验证
	cp8.SelectGoroutine()
}

func main() {
	showAllPackage()
}
