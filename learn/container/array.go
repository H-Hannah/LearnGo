package main

import (
	"fmt"
)

var arr1 [5]int //初始化一个数组
// a := [...]string{"a", "b", "c", "d"} //初始化一个数组

func loop() []int {
	// 数组循环方式1
	for i := 0; i < len(arr1); i++ {
		arr1[i] = 2
	}
	// 数组循环方式2
	for i, _ := range arr1 {
		arr1[i] = 2
	}
	return arr1[:]
}

// go语言中的数组是一种值类型，可以通过new创建
// 以下两种创建方式的区别是类型不同：
var arr2 = new([5]int) //*[5]int
var arr3 [5]int        //[5]int

// arr3 := *arr2 //数组赋值时需要做内存拷贝操作，避免相互影响

// 在函数中数组作为参数传入时，如 func1(arr2)，会产生一次数组拷贝
// 想修改原数组，必须通过&操作符以引用方式传，例如 func1(&arr2）

// 练习7.1：array_value.go: 证明当数组赋值时，发生了数组内存拷贝。
func array_value() {
	var a1 [5]int
	var a2 = new([5]int)
	for i, _ := range a2 {
		a2[i] = 1
	}
	fmt.Println("拷贝前：", a1) // [0 0 0 0 0]
	fmt.Println("拷贝前：", a2) // &[1 1 1 1 1]
	a1 = *a2                // 将a2内存地址传给a1
	fmt.Println("拷贝后：", a1) // a1发生改变：[1 1 1 1 1]
	fmt.Println("拷贝后：", a2) // &[1 1 1 1 1]
}

// 练习7.2：for_array.go: 写一个循环并用下标给数组赋值（从 0 到 15）并且将数组打印在屏幕上。
func for_array() {
	var a1 [16]int
	for i := 0; i < 16; i++ {
		a1[i] = i
	}
	fmt.Println(a1)
}

// 练习7.3：fibonacci_array.go: 通过数组计算出 Fibonacci 数。完成该方法并打印出前 50 个 Fibonacci 数字。
func fibonacci_array() {
	var arr [50]int
	arr[0] = 1
	arr[1] = 1
	for i := 2; i <= len(arr)-1; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	fmt.Println(arr)
}
func myTest() {
	// 3中已知值的数组初始化
	var arrAge = [5]int{18, 20}                       // [18 20 0 0 0]
	var arrLazy = [...]int{5, 6, 7, 8, 22}            // [18 20 0 0 0]
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"} // [   Chris Ron]
	fmt.Println(arrAge)
	fmt.Println(arrLazy)
	fmt.Println(arrKeyValue)
	// 可以取任意数组常量的地址来作为指向新实例的指针
	var a *int
	a = &arrAge[0]
	fmt.Println(*a) // 18
}

func math() {
	// 几何点（或者数学向量）是一个使用数组的经典例子。
	type Vector3D [3]float32 // 使用alias别名
	var vec Vector3D
	fmt.Println(vec)
}

// 多维数组 [3][5]int，[2][2][2]float64
func manyArray() {
	const (
		WIDTH  = 5
		HEIGHT = 10
	)
	type pixel int
	var screen [WIDTH][HEIGHT]pixel
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 1
		}
	}
	fmt.Println(screen)
}

// 把大数组传递给函数会消耗很多内存。两种方法避免：
// 1 传递数组的指针 *
// 2 使用数组的切片
func sum(arr *[10]int) {
	res := 0
	for i, _ := range arr {
		res += arr[i]
	}
	fmt.Println(res)
}

// func main() {
// 	// fmt.Println(loop())
// 	// array_value()
// 	// for_array()
// 	// fibonacci_array()
// 	// myTest()
// 	// manyArray()
// 	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	sum(&arr) // 入参为数组的指针
// }
