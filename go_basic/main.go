package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	//"sync"
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
	//类型推断
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
func basicDataTest() {
	//基本数据类型：int ，string 赋值与初始化
	BasicType()
	//基本数据类型：bool 赋值与初始化
	BoolType()
	//基本数据类型：int 64 与 32 位转化
	IntDemo()
	//基本数据类型：string 拼接
	StringDemo()
	//基本数据类型：string 与 int byte的相互转化
	testConvert()
	//枚举数据类型
	enumDemo()
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

func MustDivece(a, b int) int {
	if b == 0 {
		panic("divice is zore")
	}
	return a / b
}

func runMustDrive(a, b int) (res int, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()
	res = MustDivece(a, b)
	return
}

func errorTest() {
	//defer调用时机 test
	fmt.Println(deferDemo())
	//error 数据类型：自定义error
	runDivice()
	//触发panic 与 恢复
	res, err := runMustDrive(10, 0)
	fmt.Println(res, err)
}

func testArray() {
	//数值类型数组 初始化 赋值
	var arrayInt64 [3]int
	arrayInt64[0], arrayInt64[1] = 1, 2
	fmt.Println(arrayInt64)
	//字符串类型数组 初始化 赋值
	arrayString := [3]string{"zhang", "wang", "li"}
	fmt.Println(arrayString)
	//	省略长度的初始化数组
	arrayFloat := [...]float64{1.1, 2.2, 3.3}
	fmt.Println(arrayFloat)
	//	二维数组初始化
	matrix := [2][2]int64{
		{0, 1},
		{2, 3},
	}
	fmt.Println(matrix)
	//字符串类型数组 初始化 与遍历
	names := [3]string{"zhao", "qian", "sun"}
	fmt.Println(len(names), names[1])
	names[1] = "qian2"
	fmt.Println(len(names), names[1])
	fmt.Println(names)
}

func testSlice() {
	//切片数值类型：赋值，输出
	names := []string{"zhao", "qian", "sun"}
	fmt.Println(names, len(names), cap(names))
	//	切片赋值 另一个变量：验证赋值是一种引用
	names2 := names
	fmt.Println(names2)
	names[0] = "zhao2"
	fmt.Println(names, names2)
	names3 := names2[:2]
	names[0] = "zhao3"
	fmt.Println(names3)
	//	遍历操作
	for index, value := range names {
		fmt.Println(index, value)
	}
	//	slice 操作,验证append 拼接以后生成了新的对象
	vals := make([]int, 0)
	for i := 0; i < 3; i++ {
		vals = append(vals, i)
	}
	fmt.Println(vals)
	vals2 := []int{4, 3, 5}
	newValues := append(vals, vals2...)
	fmt.Println()
	fmt.Println("before is:", newValues)
	sort.Ints(newValues)
	sort.Sort(sort.Reverse(sort.IntSlice(newValues)))
	//fmt.Printf("revert sort value is:",newValues,"\n size is: %d cap is :%d", len(newValues), cap(newValues))
	fmt.Printf("size is: %d cap is :%d \n", len(newValues), cap(newValues))
	newValues = append(newValues, 99)
	newValues = append(newValues, 99)
	newValues = append(newValues, 99)
	newValues = append(newValues, 99)
	fmt.Printf("\n size is: %d cap is :%d\n", len(newValues), cap(newValues))

	vals[0] = 100
	fmt.Println(vals, newValues)
	//slice 的 普通赋值模式与高效模式
	//普通模式
	manyInt := make([]int, 100)
	a := make([]int, 0)
	for _, val := range manyInt {
		a = append(a, val+val)
	}
	fmt.Println("a ====", a)
	//	高效模式
	b := make([]int, len(manyInt))
	for i, val := range manyInt {
		b[i] = val + 1
	}
	fmt.Println("b ====", b)
}

func arrayTest() {
	//	数组类型：初始化与输出
	testArray()
	testSlice()
}

func initMap() map[string]int {
	m := make(map[string]int)
	m["zhang"] = 1
	m["hello"] = 2
	m["world"] = 3
	return m
}

func delMap(m map[string]int) {
	fmt.Println(m)
	delete(m, "hello")
	fmt.Println(m)
	if v, ok := m["hello"]; ok {
		fmt.Printf("%s value is: %d \n", "hello", v)
	} else {
		fmt.Printf("%s value is: %d \n", "hello", -1)
	}
	if v, ok := m["world"]; ok {
		fmt.Printf("%s value is: %d \n", "world", v)
	}
}

func rangeMap(m map[string]int) {
	for k, v := range m {
		fmt.Printf("m[\"%s\"] value is: %d \n", k, v)
	}
}

func setMap() {
	m := make(map[string]bool)
	m["hello"] = true
	m["world"] = true
	m["test_F"] = false
	key := "hello"
	key_f := "test_F"
	if _, ok := m[key]; ok {
		fmt.Printf("%s is exists\n", key)
		fmt.Println(ok)
	}
	if _, ok := m[key_f]; ok {
		fmt.Printf("%s is exists\n", key_f)
		fmt.Println(ok)
	}
}

//计算文本中重复单词的个数
func calNumByMap() {
	FilePath := "./word.txt"
	InputFile, InputError := os.OpenFile(FilePath, os.O_CREATE|os.O_RDWR, 0644)
	if InputError != nil {
		fmt.Println("open file fail ---")
		return
	}
	defer InputFile.Close()
	InputReader := bufio.NewReader(InputFile)
	calMap := make(map[string]int)
	for {
		inputString, readError := InputReader.ReadString('\n')
		//文件结束判断
		if readError == io.EOF {
			fmt.Println("file error ---")
			break
		}
		calMap[inputString] += 1
	}
	fmt.Println(calMap)
	for k, v := range calMap {
		fmt.Println(strings.TrimRight(k, "\n"), "---is---:", v)
	}
}

//map 函数实现 slice 数据的去重
func unique(intVals []int) []int {
	tmpMap := make(map[int]bool)
	for _, val := range intVals {
		if _, ok := tmpMap[val]; ok {
			continue
		}
		tmpMap[val] = true
	}
	fmt.Printf("maps size is:%d \n", len(tmpMap))
	//retSlice := make([]int, 0)
	//for k, _ := range tmpMap {
	//	retSlice = append(retSlice, k)
	//}
	retSlice := make([]int, len(tmpMap))
	i := 0
	for k, _ := range tmpMap{
		retSlice[i] = k
		i += 1
	}
	return retSlice
}

func concurrentMap(){
	m := make(map[int] int)
	//go 写协程
	go func() {
		for i := 0; i < 100; i++{
			m[i] = i
		}
	}()
	time.Sleep(10)
	//go 读协程
	go func(){
		for i := 0; i < 100; i++{
			fmt.Printf("key:%d  value:%d \n", i, m[i])
		}
	}()
}

func mapTest() {
	//	初始化一个map
	m := initMap()
	fmt.Printf("m[\"%s\"] is: %d \n", "hello", m["hello"])
	//删除map 元素
	delMap(m)
	//遍历map 的每一个元素
	rangeMap(m)
	//map 实现一个set
	setMap()
	//	cal words num
	calNumByMap()
	//slice数据类型去重复
	intVals := []int{1, 2, 1, 1, 1, 5, 6, 8, 8, 8, 9, 9, 10}
	retVals := unique(intVals)
	for index, val := range retVals {
		fmt.Printf("index is:%d value is:%d \n", index, val)
	}
//	go 并发
	concurrentMap()
}

func ifSwitchFor(){
	iFlag := 5
	if iFlag > 6{
		fmt.Printf("iFlag is: %d > 6 \n",iFlag)
	}else if iFlag == 6{
		fmt.Printf("iFlag is: %d == 6 \n",iFlag)
	}else {
		fmt.Printf("iFlag is: %d < 6 \n",iFlag)
	}
	switch iFlag {
	case 0, 6:
		fmt.Printf("iFlag is: %d == 0 or 6 \n",iFlag)
	case 1,2,3,4,5:
		fmt.Printf("iFlag is: %d == 1,2,3,4,5 \n",iFlag)
	default:
		fmt.Printf("iFlag is: %d != 1,2,3,4,5,6 \n",iFlag)
	}
}

func sliceFor(){
	sSlice := []string {"key1", "key2", "key3"}
	for index, value := range sSlice{
		fmt.Println(index, value, "++++++")
	}
	//for 遍历 slice 只有一个参数接收返回时，该参数接收的是索引
	for index := range sSlice{
		fmt.Println(index, "=======")
	}
	sMap := map[string] string {"k1":"v1", "k2":"v2"}
	for k, v := range sMap{
		fmt.Println(k,v)
	}
	//for 遍历 map 只有一个参数接收返回时，该参数接收的是key
	for k := range sMap{
		fmt.Println(k, "--- --- ---")
	}
}

func loopTest(){
	ifSwitchFor()
	sliceFor()
}

//类型传递
//内置类型：数值类型、字符串、布尔类型、数组。传递的是副本 (所以一般不用数组啦)，会拷贝原始值，无法修改
//引用类型: 切片、映射、通道、接口和函数类型。通过复制传递应用类型值的副本，本质上就是共享底层数据结构。可以修改
func testAnonymouseFunc(){
	func (s string){
		fmt.Println(s)
	}("hello helllo hello hello hello hello hello hello hello hello hello hello ")
}
//很多语言都有闭包的概念， 所谓闭包就是一个函数“捕获”了和它在同一作用域的其他常量和变量。
//当闭包被调用的时候，不管在程序什么地方调用，闭包能够使用这些常量或者变量，
//并且只要闭包还在使用它，这些变量就不会销毁，一直存在。 上文中提到的匿名函数其实就是闭包。来看一个简单的示例
// 闭包示例
//func testClosure() {
//	suffix := ".go"
//	addSuffix := func(name string) string {
//		return name + suffix // 这里使用到了 suffix 这个变量，所以 addSuffix 就是一个闭包
//	}
//	fmt.Println(addSuffix("hello_world"))
//}
//对于一般不太严重的场景，返回错误值 error 类型 (业务绝大部分场景)
//对于严重的错误需要整个进程退出的场景，使用 panic 来抛异常，及早发现错误
//如果希望捕获 panic 异常，可以使用 recover 函数捕获，并且包装成一个错误返回
//web 框架等会帮你捕获 panic 异常，然后返回客户端一个 http 500 状态码错误

func funcTest(){
	testAnonymouseFunc()
}
func main() {
	//basicDataTest()
	//errorTest()
	//arrayTest()
	//mapTest()
	//loopTest()
	funcTest()
}