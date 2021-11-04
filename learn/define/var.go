package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
)

// 变量命名遵循驼峰法：当全局变量希望被外部包所使用，则需将首个单词的首字母也大写
var a1, b1 *int // nil

var (
	x   int    // 0
	y   bool   // false
	str string // ""
)

// 显式指定变量类型
var n int64 = 2

// 声明包级别的全局变量
var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

func myFunction() {
	// 函数体内声明局部变量，简单写法
	han := "hsq"
	fmt.Println(han)
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}

// 值类型和引用类型
// 引用类型有：指针&、切片splice、通道channel、映射map
func myFunction1() {
	// 在内存中将i的值进行拷贝，通过 &i 可获取变量 i 的内存地址
	i := 1
	j := 2
	j = i // 只有引用（地址）被复制
	fmt.Println(&i, &j)
	// fmt.Sprintf 与 Printf 完全相同，前者将格式化后的字符串以返回值的形式返回
	// fmt.Sprintf(&j)
}

func myFunction2() {
	var a0, b0, c0 int
	a0, b0, c0 = 6, 7, 8
	a0, b0 = b0, a0 // 交换两个变量的值
	_, c0 = a0, b0  // 下划线用来占位忽略某值
	// a0, b0, c0 := 5, 7, "abc"	// 上述简写，并行or同时赋值
	fmt.Println(a0, b0, c0)
}

func init() {
	// go backend()	// 调用后台的goroutine
	Pi := 4 * math.Atan(1)
	// var twoPi = 2 * Pi
	fmt.Println(Pi)
}

// 全局变量：允许声明后不使用；局部变量：不允许
// func main() {
// 	myFunction2()
// }
