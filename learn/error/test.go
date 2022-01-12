// defer-panic-and-recover 异常处理机制
// 通过在函数和方法中返回错误对象作为它们的唯一或最后一个返回值
// 如果返回 nil，则没有错误发生，并且主调（calling）函数应常去检查收到的错误

package main

import (
	"fmt"
)

// error接口类型
// type error interface {
// 	Error() string
// }

// 创建新的错误类型：使用 errors.New 函数接收合适的错误信息
// var errNotFound error = errors.New("Not found error")

// func Sqrt(f float64) (float64, error) {
// 	if f < 0 {
// 		return 0, errors.New("math - square root of negative number")
// 	}
// }

// func main() {
// 	fmt.Printf("error: %v", errNotFound)
// }

// 运行时异常：有个 RuntimeError() 方法用于区别普通错误
// 在多层嵌套的函数调用中调用 panic，可以马上中止当前函数的执行，所有的 defer 语句都会保证执行并把控制权交还给接收到 panic 的函数调用者
// 这样向上冒泡直到最顶层，并执行（每层的） defer，在栈顶处程序崩溃，并在命令行中用传给 panic 的值报告错误情况：这个终止过程就是 panicking。
// recover 内建函数被用于从 panic 或 错误场景中恢复：让程序可以从 panicking 重新获得控制权，停止终止过程进而恢复正常执行。

// panic_recover.go

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n") // <-- would not reach
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}

// defer-panic-recover 在某种意义上也是一种像 if，for 这样的控制流机制
