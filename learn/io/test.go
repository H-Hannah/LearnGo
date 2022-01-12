// 从控制台读取输入:
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
// firstName, lastName string
// s                   string
// i                   int
// f                   float32
// input = "56.12 / 5212 / Go"
// format              = "%f / %d / %s"
)

func main1() {
	// 1 扫描来自标准输入的文本
	// fmt.Scanln(&firstName, &lastName)
	// 2 扫描按"%s %s"格式输入的文本
	// fmt.Scanf("%s %s", &firstName, &lastName)
	// fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi hannah,joy !

	// 3 从字符串input中按照format读取数据
	// fmt.Sscanf(input, format, &f, &i, &s)
	// fmt.Println("From the string we read: ", f, i, s)

	// 4 bufio包提供的缓冲读取数据
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Println(input)
	}
}
