package main

import (
	"fmt"
	"unicode/utf8"
)

// 布尔型(以is开头命名)、数字型和字符型
// 任何不同类型之间的转换都必须显式说明
// 值之间的比较必须是类型相同的值

func myFunc() {
	// 定义一个 string 的类型别名 Rope，并声明一个该类型的变量
	type Rope string
	var s Rope = "hello"
	fmt.Println(s)
}

func myFunc1() {
	// 创建一个用于统计字节和字符（rune）的程序，并对字符串 asSASA ddd dsjkdsjs dk 进行分析
	// 然后再分析 asSASA ddd dsjkdsjsこん dk，最后解释两者不同的原因（提示：使用 unicode/utf8 包）
	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The number of bytes in string str1 is %d\n", len(str1))
	fmt.Printf("The number of characters in string str1 is %d\n", utf8.RuneCountInString(str1))
	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The number of bytes in string str2 is %d\n", len(str2))
	fmt.Printf("The number of characters in string str2 is %d", utf8.RuneCountInString(str2))
}
// func main() {
// 	// 类型别名，int可表示为TZ
// 	// 类型别名得到的新类型并非和原类型完全相同，新类型不会拥有原类型所附带的方法
// 	type TZ int
// 	var a TZ = 3
// 	fmt.Println(a)

// 	// 字符串拼接：循环中不用+，使用strings.Join()或字节缓冲 bytes.Buffer
// 	s := "hel" + "lo,"
// 	s += "world!"
// 	fmt.Println(s)

// 	myFunc1()
// }
