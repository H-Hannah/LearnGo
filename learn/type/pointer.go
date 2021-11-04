package main

import "fmt"

func myFunc3() {
	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
	var iP *int      // 一个指针变量通常缩写为 ptr
	fmt.Println(iP)  // nil 一个指针被定义后没有分配到任何变量时
	iP = &i1         // 此时iP指向i1，一个指针变量可以指向任何一个值的内存地址
	fmt.Println(iP)  // 0xc000012080
	fmt.Println(*iP) // 间接引用指针指向的值 //5

}
func main() {
	myFunc3()
}
