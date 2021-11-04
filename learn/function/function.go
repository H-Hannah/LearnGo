package main

import (
	"errors"
	"fmt"
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

func main() {
	n := 0
	reply := &n                 // 复制n变量在内存中的地址
	Multiply(10, 5, reply)      // 传给函数 Multiply
	fmt.Println("Multiply:", n) // Multiply: 50
}
