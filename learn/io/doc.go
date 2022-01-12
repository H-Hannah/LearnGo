package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main2() {
	// 文件使用指向 os.File 类型的指针来表示，也叫文件句柄
	inputFile, inputError := os.Open("input.dat")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close() // 关闭文件流

	inputReader := bufio.NewReader(inputFile)
	// 无限循环中按行读取文件
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}

}

// 文件拷贝
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()	// 发生错误仍需关闭文件

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
