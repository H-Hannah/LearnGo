package main

import (
	"fmt"
)

// 切片
// 切片（slice）是对数组一个连续片段的引用，又称为匿名的“相关数组”。-- 引用类型
// 这个片段可以是整个数组，或者是由起始和终止索引标识的一些项的子集。（终止索引处为开区间）
// 切片提供了一个相关数组的动态窗口。是一个长度可变的动态数组。
// s为切片，cap(s)为s[0]到数组末尾的数组长度。0 <= len(s) <= cap(s)

// 多个切片如果表示同一个数组的片段，它们可以共享数据；因此一个切片和相关数组的其他切片是共享存储的。
// 相反，不同的数组总是代表不同的存储。数组实际上是切片的构建块。

// 声明切片的格式是： var identifier []type 默认为 nil，长度为 0
// 初始化一个切片： var slice1 []type = arr1[start:end]
// 切片扩展到他的长度上限 s = s[:cap(s)]

// 切片在内存中是：一个有 3 个域的结构体 -- 指向相关数组的指针ptr，切片长度len 以及切片容量cap
// 切片后移一位 s = s[1:]

// 问题 7.2： 给定切片 b:= []byte{'g', 'o', 'l', 'a', 'n', 'g'}，那么 b[1:4]、b[:2]、b[2:] 和 b[:] 分别是什么？
func test1() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b[1:4]) // [111 108 97]
	fmt.Println(b[:2])  // ['g', 'o']
	fmt.Println(b[2:])  // ['a', 'n', 'g']
	fmt.Println(b[:])   // ['g', 'o', 'l', 'a', 'n', 'g']
}

// 使用make创建切片
// var slice1 []type = make([]type, len)

// 两种方式生成相同的切片
// make([]int, 50, 100)	参数分别为：指针指向的类型，len，cap
// new([100]int)[0:50]
// 区别：new 函数分配内存，make 函数初始化

// 定义一个字节型buffer：var buffer bytes.Buffer
// 通过 buffer 串联字符串，比 += 更省空间，类似于 stringbuilder
// func myBuffer() {
// 	var buffer bytes.Buffer
// 	for {
// 		if s, ok := getNextString(); ok {
// 			buffer.WriteString(s) // 将字符串s追加到后面
// 		} else {
// 			break
// 		}
// 	}
// 	fmt.Print(buffer.String(), "\n") // 转为string字符串
// }

// 练习 7.5 给定切片 sl，将一个 []byte 数组追加到 sl 后面。
// 写一个函数 Append(slice, data []byte) []byte，该函数在 sl 不能存储更多数据的时候自动扩容。

// func Append(slice, data []byte) []byte {
// 	var buffer bytes.Buffer
// 	for temp := range data {
// 		buffer.WriteString(string(temp))
// 	}
// 	fmt.Println(buffer.String())
// 	return data
// }

// 数组和切片中的for循环

// 第一个返回值 ix 是数组或者切片的索引，第二个是在该索引位置的值；
// 他们都是仅在 for 循环内部可见的局部变量。
// value 只是 slice1 某个索引位置的值的一个拷贝，不能用来修改 slice1 该索引位置的值。
// for ix, value := range slice1 {
// 	...
// }

// 示例如下
func season() {
	seasons := []string{"spring", "summer", "autumn", "winter"}
	for ix, season := range seasons {
		fmt.Printf("第%d个季节是：%s\n", ix, season)
		season = "hello" // 修改值失败
	}
	fmt.Println(seasons) // [spring summer autumn winter]
}

// 切片重组
// 改变切片长度的过程称之为切片重组 reslicing -- slice1 = slice1[0:end]
// 将切片扩展1位 slice1 = slice1[0:len(slice1+1)]
// 切片可以反复扩展直到占据整个相关数组。
func reslicing() {
	sl1 := make([]int, 0, 10)       //长度为0，容量为10的切片
	for i := 0; i < cap(sl1); i++ { // 扩展直到最大的相关分组
		sl1 := sl1[0 : i+1]
		sl1[i] = i
		fmt.Println(sl1) // [0 1 2 3 4 5 6 7 8 9]
	}
	fmt.Println(len(sl1)) // 0
}

// 如果 a 是一个切片，那么 a[n:n] 的长度是多少？ -- 1
// a[n:n+1] 的长度又是多少？ -- 1

// 切片的复制与追加
// 若想要增加切片的容量，1 创建更大的切片  2 复制原有  3 切片追加新元素
func copy_append_slice() {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)
	copy(slTo, slFrom) // 切片复制过去
	fmt.Println(slTo)  // [1 2 3 0 0 0 0 0 0 0]
	slTemp := []int{4, 5, 6}
	slTo = append(slTo, slTemp...) // [1 2 3 0 0 0 0 0 0 0 4 5 6]
	fmt.Println(slTo)
}
func main() {
	// test1()
	// season()
	// reslicing()
	copy_append_slice()
}
