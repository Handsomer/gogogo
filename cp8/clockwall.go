package cp8

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

func StartServer() {
	for _, arg := range os.Args[1:] {
		argSlice := strings.Split(arg, "=")
		language, host := argSlice[0], argSlice[1]
		fmt.Println(language, host)
		go RunOneServer(language, host)
		// break
	}
	for {
		time.Sleep(5 * time.Second)
	}
}

func InitZoneTime() (timeZoneMap map[string]interface{}) {
	// timeZoneMap := make(map[string]interface{})
	timeZoneMap = make(map[string]interface{})
	timeZoneMap["China"] = "Asia/Shanghai"
	timeZoneMap["America"] = "America/Los_Angeles"
	return
}

// map 数据结构的遍历; get（通过返回值确定默认值）;
//  go 类型断言的使用：使用类型断言的数据须是 interface{}
func MapRangeAndGetUse() {
	mapStringInterface := InitZoneTime()
	zoneTime := "China11111"
	for k, v := range mapStringInterface {
		fmt.Println(k, " : ", v)
	}
	zoneStr, ok := interface{}(mapStringInterface).(map[string]interface{})[zoneTime].(string)
	if ok {
		fmt.Println(zoneStr)
	}
	fmt.Println("America/Adak")
}

// 运行一个服务, 实现：
// 根据时区每秒向客户端返回 该时区的时间
func RunOneServer(language string, host string) {
	// 初始化目标时区
	timeZoneMap := InitZoneTime()
	timeZone, ok := timeZoneMap[language].(string)
	if !ok {
		timeZone = "America/Adak"
	}
	local1, err := time.LoadLocation(timeZone)
	if err != nil {
		fmt.Println(local1)
		return
	}
	// 初始化服务器：host
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn, local1)
	}
}

// 每隔1 秒向客户端发送一次当前时间
func handleConn(conn net.Conn, local interface{}) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().In(local.(*time.Location)).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

// TODO： utf-8 编码方式
func StringUTF8Demo() {
	/*
	   13
	   9
	   0        h
	   1        e
	   2        l
	   3        l
	   4        o
	   5        ,
	   6
	   7        世
	   10       界
	*/
	s := "hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d \t %c \n", i, r)
		i += size
	}
}
