package main

import (
	"fmt"

	"github.com/yylego/runpath"
)

func main() {
	// Get current file path with different extension
	// 获取当前文件的不同扩展名路径
	jsonPath := runpath.GetPathChangeExtension(".json")
	fmt.Println("JSON config path:")
	fmt.Println(jsonPath)

	// Using compact alias GetRex
	// 使用简洁的别名 GetRex
	yamlPath := runpath.GetRex(".yaml")
	fmt.Println("YAML config path:")
	fmt.Println(yamlPath)

	// Get path without extension
	// 获取不带扩展名的路径
	basePath := runpath.GetNox()
	fmt.Println("Base path (no extension):")
	fmt.Println(basePath)

	// Demonstrate use case: reading config file with same name
	// 演示用例：读取同名配置文件
	configPath := runpath.GetRex(".config.json")
	fmt.Println("Config file path:")
	fmt.Println(configPath)
}
