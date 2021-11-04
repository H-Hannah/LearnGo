package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 字符串的预定义处理函数
func myFunc2() {
	myStr := "hello world!"
	// 1 前缀后缀
	var prefix = strings.HasPrefix(myStr, "h") //true
	var suffix = strings.HasSuffix(myStr, "h") //false
	fmt.Println(prefix)
	fmt.Println(suffix)

	// 2 字符串包含
	var contains = strings.Contains(myStr, "ll")  //true
	var index = strings.Index(myStr, "A")         //-1
	var lastIndex = strings.LastIndex(myStr, "o") //7
	fmt.Println(contains)
	fmt.Println(index)
	fmt.Println(lastIndex)

	// 3 字符串替换
	var reStr = strings.Replace(myStr, "l", "L", 2) //heLLo world!
	fmt.Println(reStr)

	// 4 统计字符串出现次数
	var count = strings.Count(myStr, "l") //3
	fmt.Println(count)

	// 5 重复字符串n次展示
	var repate = strings.Repeat(myStr, 3) //hello world!hello world!hello world!
	fmt.Println(repate)

	// 6 大小写转换
	var upper = strings.ToUpper(myStr) //HELLO WORLD!
	var lower = strings.ToLower(myStr) //hello world!
	fmt.Println(upper)
	fmt.Println(lower)

	// 7 修剪字符串，除开头or结尾
	var trimSpace = strings.TrimSpace(myStr)
	var trimL = strings.Trim(myStr, "l")
	var trimLeft = strings.TrimLeft(myStr, "hel")
	var trimRight = strings.TrimRight(myStr, "ld!")
	fmt.Println(trimSpace)
	fmt.Println(trimL)
	fmt.Println(trimLeft)
	fmt.Println(trimRight)

	// 8 字符串分割
	var slice1 = strings.Fields(myStr)     //按空格切割
	var slice2 = strings.Split(myStr, "l") //按指定字符切割
	fmt.Println(slice1)
	fmt.Println(slice2)

	// 10 拼接slice到字符串
	var str = "2021 10 04"
	var slice = strings.Fields(str)
	var str1 = strings.Join(slice, "-") //2021-10-04
	fmt.Println(str1)

	// 11 从字符串中读取内容
	var reader = strings.NewReader(str) //&{str} 返回指向该reader的指针
	fmt.Println(reader)

	// 12 字符串与其他类型转换
	var s = "666"
	var atoi int
	var iton string
	atoi, _ = strconv.Atoi(s) //相当于parseInt
	atoi += 1000
	iton = strconv.Itoa(atoi) //FormatInt
	fmt.Println("int：", atoi)
	fmt.Println("string：", iton)

	// 13 时间和日期
	fmt.Println(time.Now())
	// time.Sleep（d Duration） 可以实现对某个进程（实质上是 goroutine）时长为 d 的暂停
}

// func main() {
// 	myFunc2()
// }
