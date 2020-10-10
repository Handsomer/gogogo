package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func BasicType() {
	fmt.Println("hello world!---")
	//初始化数字与字符串的初始化
	var i int64
	var b string
	fmt.Println("i = ", i)
	fmt.Println("b = ", b)
	//
	var floatNum float64 = 10.1
	var price1, price2 float64 = 8.8, 9.6
	fmt.Println("floatNum = ", floatNum)
	fmt.Println(price1, "----", price2)
	//
	ii := 1
	ss := "Hello gogogo!"
	fmt.Printf("ii = %v"+
		" ---\n", ii)
	fmt.Printf("ss = %v  ---\n", ss)
}

func BoolType() {
	//默认是false
	var a bool
	b := true
	c := false
	if a {
		fmt.Println("a", a)
	}
	if b {
		fmt.Println("b", b)
	}
	if c {
		fmt.Println("c", c)
	}
}

func IntDemo() {
	//初始化 int64 类型
	var a int64
	a = 10
	fmt.Println(a + 10)
	//初始化 int32 类型
	i32 := int32(10)
	fmt.Println(i32 + 10)
	//输出int64 最大值
	fmt.Println(math.MaxInt64)
	fmt.Println(10 / 10)
}

func StringDemo() {
	s1 := "\"Hello world \""
	s2 := `"Hello world"`
	fmt.Println(s1)
	fmt.Println(s2)
	s3 := `
你好
`
	s4 := "Golang !"
	fmt.Println(s3 + s4)
}

func testConvert() {
	//转换操作 1
	//int 转 string
	sint := strconv.Itoa(97)
	fmt.Println(sint, sint == "97")
	//	byte 转 string
	bytea := byte(1)
	sbyte := strconv.Itoa(int(bytea))
	fmt.Println(sbyte)
	//	int64 转10进制 string
	sint64 := strconv.FormatInt(int64(97), 10)
	fmt.Println(sint64, sint64 == "97")
	// int64 转16进制string
	sint64_16 := strconv.FormatInt(int64(97), 16)
	fmt.Println(sint64_16, sint64_16 == "61")
	// 转换操作 2
	//	string 转 int
	_int, _ := strconv.Atoi("64")
	fmt.Println(_int, _int == 64)
	//string 转int64
	_int64, _ := strconv.ParseInt("97", 10, 64)
	fmt.Println(_int64, _int64 == int64(97))
	//string 转 int32
	_int32, _ := strconv.ParseInt("97", 10, 32)
	fmt.Println(_int32, int32(_int32) == int32(97))
	//int 转string 的三种方式
	i := 42
	fmt.Println(strconv.FormatInt(int64(i), 10) == "42")
	fmt.Println(strconv.Itoa(int(i)) == "42")
	fmt.Println(fmt.Sprint(i) == "42")
}

func enumDemo() {
	const (
		Sunday    = 0
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
		Thursday  = 4
		Friday    = 5
		Saturday  = 6
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	const (
		Sunday_1 = iota
		Mondat_1
		Tuesday_1
		Wednesday_1
		Thursday_1
		Friday_1
		Saturday_1
	)
	fmt.Println(Sunday_1, Mondat_1, Tuesday_1, Wednesday_1, Thursday_1, Friday_1, Saturday_1)
}

func deferDemo() string {
	//defer 的调用顺序
	defer fmt.Println("defer place one")
	defer fmt.Println("defer place two")
	fmt.Println("print in after defer place")
	return "return place"

}

func Divice(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divice by zone")
	}
	return a / b, nil
}

func runDivice() {
	a := 10
	b := 2
	res, err := Divice(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
func main() {
	//基本数据类型：int ，string 赋值与初始化
	//BasicType()
	//基本数据类型：bool 赋值与初始化
	//BoolType()
	//基本数据类型：int 64 与 32 位转化
	//IntDemo()
	//基本数据类型：string 拼接
	//StringDemo()
	//基本数据类型：string 与 int byte的相互转化
	//testConvert()
	//枚举数据类型
	//enumDemo()
	//defer调用时机 test
	//fmt.Println(deferDemo())
	//error 数据类型：自定义error
	runDivice()
}
