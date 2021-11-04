package main

import "fmt"

const a = 1 // 常量

var b = 2 // 变量

// 入参名称后紧跟类型
func myFunction(x string, y bool) {
	fmt.Println(a + b)
	fmt.Println("string x:", x)
	fmt.Println("bool y:", y)
}

// 没有main或有参有返回值时会报构建错误
func main() {
	// 原始打印
	// 导入某个包不使用时会报错
	fmt.Println("hello world")

	// 包fmt的别名fm(alias) import fm "fmt"
	// fm.Println("hello world")

	// 调用其他函数
	myFunction("hello", true)

	// 多语言
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")

	// 调用有返回值的函数
	fmt.Println(function1(1, 2))

	fmt.Println(myTest())
} // 正常退出返回0，否则1

/** 类型：
基本类型 -- int、float、bool、string
结构化的 -- struct、array、slice、map、channel （默认值为nil）
描述行为 -- interface
*/

// 有入参与返回值的函数
func function1(m int, n int) (t1 int, t2 int) {
	return m, n
}

// 多个类型定义
type (
	IZ  int
	FZ  float64
	STR string
)

// 含有该init函数的包默认会先执行init
func init() {
	var i = 0
	i = 1
	fmt.Println("i", i)
}

/** go执行顺序
1 按顺序导入所有被 main 包引用的其它包，然后在每个包中执行如下流程：
2 如果该包又导入了其它的包，则从第一步开始递归执行，但是每个包只会被导入一次。
3 然后以相反的顺序在每个包中初始化常量和变量，如果该包含有 init 函数的话，则调用该函数。
4 在完成这一切之后，main 也执行同样的过程，最后调用 main 函数开始执行程序。
*/

func myTest() (h IZ) {
	// 简单定义+类型强转
	// a := 5.0
	// b := int(a)
	// return b

	// 相同底层类型的变量之间可以转换
	var a IZ = 5
	c := int(a)
	d := IZ(c)
	return d
}