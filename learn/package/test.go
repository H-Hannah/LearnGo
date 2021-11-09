package main

import (
	"bytes"
	"fmt"
	"sync"
)

// container/list包，可实现元素插入、删除、移动等功能

func test1() {
	// 练习 9.1：使用 container/list 包实现一个双向链表，将 101、102 和 103 放入其中并打印出来。
}

type Info struct {
	mu sync.Mutex // sync.Mutex 是一个互斥锁
}

func test2(info *Info) {
	// 锁和sync包
	// 一次只能让一个线程对共享变量进行操作
	info.mu.Lock() // 上锁
	// info.Str = "hello" // 修改资源
	info.mu.Unlock() // 解锁
}

// 上锁的共享缓冲器
type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

// func main() {
// test1()
// test2(&Info{})
// }

// new()一个 map 返回的是一个指向 nil 的指针，它尚未被分配内存，可以使用make
// make() 一个 type 返回的是一个指向 nil 的指针，建议使用 new

// 接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针。
// 一个类型 type 加上它的方法 method 等价于面向对象中的一个类

// T（或 *T）和它的方法集（method set）
// 不允许方法重载
// 基于接收者类型，是有重载的：具有同样名字的方法可以在 2 个或多个不同的接收者类型上存在，比如在同一个包里这么做是允许的

// func (a *denseMatrix) Add(b Matrix) Matrix
// func (a *sparseMatrix) Add(b Matrix) Matrix

// 定义一个方法
// func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
// 使用这个方法
// recv.methodName()

// 示例：结构体类型上的方法
type TwoInts struct {
	a int
	b int
}

func (tn TwoInts) Add() int {
	return tn.a + tn.b
}

func (tn TwoInts) AddThem(param int) int {
	return tn.a + tn.b + param
}

// func main() {
// 	my := new(TwoInts)
// 	my.a = 2
// 	my.b = 3
// 	fmt.Println(my.Add())
// 	fmt.Println(my.AddThem(10))
// }

// 示例：非结构体类型上的方法
type IntVector []int

func (v IntVector) Sum() (result int) {
	for _, x := range v {
		result += x
	}
	return
}

// func main() {
// 	fmt.Println(IntVector{1, 2, 3}.Sum()) //6
// }

// 练习 10.6 employee_salary.go
// 定义结构体 employee，它有一个 salary 字段，给这个结构体定义一个方法 giveRaise 来按照指定的百分比增加薪水。

type Employee struct {
	salary float64
}

// 不能忘记方法名后面的括号
func (em Employee) GiveRaise() (res int) {
	if em.salary < 15 {
		res = 100
	} else {
		res = 200
	}
	return
}

// func main() {
// 	employee := new(Employee)
// 	employee.salary = 20
// 	fmt.Println(employee.GiveRaise())
// }

// 练习 10.8 inheritance_car.go
// 创建一个上面 Car 和 Engine 可运行的例子，并且给 Car 类型一个 wheelCount 字段和一个 numberOfWheels() 方法。
// 创建一个 Mercedes 类型，它内嵌 Car，并新建 Mercedes 的一个实例，然后调用它的方法。
// 然后仅在 Mercedes 类型上创建方法 sayHiToMerkel() 并调用它。

type Car struct {
	wheelCount int
}

func (myCar Car) numberOfWheels() int {
	return myCar.wheelCount
}

type Engine interface {
	Start()
	Stop()
}

type Mercedes struct {
	Car // 内嵌
}

func (mer Mercedes) sayHiToMerkel() {
	fmt.Println(11111111)
	mer.sayHiToMerkel()
}

func main() {
	my := new(Mercedes) // 创建实例
	fmt.Println(my.Car) // 调用方法
}

// 如何在类型中嵌入功能？两种方法
// A：聚合：包含一个所需功能类型的具名字段
// B：内嵌：内嵌（匿名地）所需功能类型

// 多重继承

