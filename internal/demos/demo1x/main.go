package main

import (
	"fmt"

	"github.com/yylego/runpath"
)

func main() {
	// Get current source file path
	// 获取当前源文件路径
	currentPath := runpath.Path()
	fmt.Println("Current source path:")
	fmt.Println(currentPath)

	// Get current source file name
	// 获取当前源文件名称
	currentName := runpath.Name()
	fmt.Println("Current source name:")
	fmt.Println(currentName)

	// Get parent DIR path
	// 获取父 DIR 路径
	parentPath := runpath.PARENT.Path()
	fmt.Println("Parent DIR path:")
	fmt.Println(parentPath)

	// Get parent DIR name
	// 获取父 DIR 名称
	parentName := runpath.PARENT.Name()
	fmt.Println("Parent DIR name:")
	fmt.Println(parentName)
}
