package main

import "fmt"

// 匿名函数
func f1() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) } // 匿名函数赋值给变量
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g) // 变量 g 代表 func(int)，变量的值是一个内存地址
	}
}

// 练习 6.8 在 main 函数中写一个用于打印 Hello World 字符串的匿名函数并赋值给变量 fv，然后调用该函数并打印变量 fv 的类型。
// func main() {
// 	fv := func() {
// 		fmt.Println("Hello World")
// 	}
// 	fmt.Printf("类型为%T", fv, fv)
// }

// 匿名函数同样被称之为闭包，闭包为函数式编程的术语，它们被允许调用定义在其它环境下的变量。
// 一个闭包继承了函数所声明时的作用域。


