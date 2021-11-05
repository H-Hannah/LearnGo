package main

import (
	"fmt"
	"time"
)

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

// 将匿名函数作为返回参数
// func main() {
// 	p2 := Add2()
// 	fmt.Println(p2(3))

// 	TwoAdder := Adder(2)
// 	fmt.Println(TwoAdder(3))
// }

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

// 闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。

// 练习 6.9 不使用递归但使用闭包改写斐波那契数列程序
func fibonacci() func() int {
	back1, back2 := 0, 1
	return func() int {
		back1, back2 = back2, back1+back2 // 重新赋值
		return back1
	}
}
func main() {
	// f := fibonacci() //返回一个闭包函数
	// var array [10]int
	// for i := 0; i < 10; i++ {
	// 	array[i] = f()
	// }
	// fmt.Println(array)
	my()
}

// 可以返回其它函数的函数和接受其它函数作为参数的函数均被称之为"高阶函数"，是函数式语言的特点

// 计算函数执行时间
func longCalculation() {
	var a int
	for i := 0; i < 2e10; i++ {
		a += i
	}
}
func my() {
	start := time.Now()
	longCalculation()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("函数执行时间为: %s\n", delta) //4.856767561s
}
