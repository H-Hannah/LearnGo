package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
)

// 练习 6.1 mult_returnval.go
// 编写一个函数，接收两个整数，然后返回它们的和、积与差。编写两个版本，一个是非命名返回值，一个是命名返回值。
// 命名返回值版本（建议）
func mult_returnval(input1 int, input2 int) (add int, multi int, decide int) {
	add = input1 + input2
	multi = input1 * input2
	decide = input1 - input2
	return
}

// 非命名返回值版本
func mult_returnval1(input1 int, input2 int) (int, int, int) {
	return input1 + input2, input1 * input2, input1 - input2
}

// 练习 6.2 error_returnval.go
// 编写一个名字为 MySqrt 的函数，计算一个 float64 类型浮点数的平方根，如果参数是一个负数的话将返回一个错误。编写两个版本。
func error_returnval(input float64) (res float64, err error) {
	if input < 0 {
		res = float64(math.NaN())
		err = errors.New("当前输入错误！")
	} else {
		res = math.Sqrt(input)
	}
	return
}
func error_returnval1(input float64) (float64, error) {
	if input < 0 {
		return float64(math.NaN()), errors.New("当前输入错误！")
	}
	return math.Sqrt(input), nil
}

func Multiply(a, b int, reply *int) {
	*reply = a * b // 修改reply地址处的变量(即n)的值
}

// func main() {
// 	n := 0
// 	reply := &n                 // 复制n变量在内存中的地址
// 	Multiply(10, 5, reply)      // 传给函数 Multiply
// 	fmt.Println("Multiply:", n) // Multiply: 50
// }

// 变参函数
// func myFunc(a, b, arg ...int) {
// 	// 这个函数接受一个类似某个类型的 slice 的参数
// 	return
// }

func min(s ...int) int {
	if len(s) == 0 { // s为slice切片
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

// func main() {
// 	x := min(1, 3, 2, 0)
// 	fmt.Printf("The minimum is: %d\n", x)
// 	slice := []int{7, 9, 3, 5, 1}
// 	x = min(slice...)
// 	fmt.Printf("The minimum in the slice is: %d\n", x)
// }

// 练习 6.3 varargs.go
// 写一个函数，该函数接受一个变长参数并对每个元素进行换行打印。
func varargs(s ...string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// func main() {
// 	// 变长参数可以作为对应类型的 slice 进行二次传递
// 	s := []string{"a", "b", "c", "h", "e", "l"}
// 	varargs(s...)
// }

// 定义结构体(传递多个不确定参数时)
// type Options struct {
// 	par1 string,
// 	par2 string
// 	...
// }

// func myTest(a int,b int,Options{par1:"hello",par2:"world"}){
// 	fmt.Println(Options)
// }

// 如果变长参数的类型没有被指定，则可用默认空接口 interface{} 来接受任何类型的参数
// func typecheck(..,..,values … interface{}) {
// 	for _, value := range values {
// 		switch v := value.(type) {
// 			case int: …
// 			case float: …
// 			case string: …
// 			case bool: …
// 			default: …
// 		}
// 	}
// }

// 关键字 defer 允许推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数
// 类似于 finally 语句块
func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2() // 推迟到最后执行
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("Function2")
}

// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

// defer常用于函数的收尾工作，eg:关闭连接池，资源解锁，关闭文件流，打印报告

// 使用 defer 语句实现代码追踪
func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

// 示例
func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

// 内置函数
// close	用于管道通信
// len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）
// cap 是容量的意思，用于返回某个类型的最大容量（只能用于数组、切片和管道，不能用于 map）

// new 和 make 均是用于分配内存：
// new 用于值类型和用户定义的类型，如自定义结构
// make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。
// new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针。它也可以被用于基本类型：v := new(int)。
// make(T) 返回类型 T 初始化之后的值，因此它比 new 进行更多的工作

// opy、append	用于复制和连接切片
// panic、recover	两者均用于错误处理机制
// print、println	底层打印函数，在部署环境中建议使用 fmt 包
// complex、real imag	用于创建和操作复数

// 递归函数
// 斐波那契数列
func fabnacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fabnacci(n-1) + fabnacci(n-2)
	}
	return
}

// func main() {
// 	// result := 0
// 	// for i := 0; i <= 10; i++ {
// 	// 	result = fabnacci(i)
// 	// 	fmt.Printf("fibonacci(%d) is: %d\n", i, result)
// 	// }

// 	// myTest1(10)

// 	// multi(10)

// 	// f1()
// }

// 使用递归函数从 10 打印到 1
func myTest1(n int) {
	fmt.Println(n)
	if n > 1 {
		myTest1(n - 1)
	}
}

// 实现一个输出前 30 个整数的阶乘的程序
func multi(n int) (res int) {
	if n == 1 {
		res = 1
	} else {
		res = n * multi(n-1)
	}
	log.Printf("第 %d个数字的阶乘：", n, res)
	return
}

// 将函数可作为参数
