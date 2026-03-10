package main

import (
	"fmt"

	"github.com/yylego/runpath"
)

func main() {
	// Join paths with parent DIR
	// 与父 DIR 连接路径
	configPath := runpath.PARENT.Join("config.json")
	fmt.Println("Config path:")
	fmt.Println(configPath)

	dataPath := runpath.DIR.Join("data", "example.txt")
	fmt.Println("Data path:")
	fmt.Println(dataPath)

	// Navigate up DIR structure
	// 向上导航 DIR 结构
	upOne := runpath.PARENT.Up(1)
	fmt.Println("Up 1 DIR:")
	fmt.Println(upOne)

	upTwo := runpath.PARENT.Up(2)
	fmt.Println("Up 2 DIRs:")
	fmt.Println(upTwo)

	// Navigate up and join with path
	// 向上导航并连接路径
	moduleConfig := runpath.PARENT.UpTo(1, "config.json")
	fmt.Println("Module config:")
	fmt.Println(moduleConfig)

	projectConfig := runpath.PARENT.UpTo(2, "config", "settings.json")
	fmt.Println("Project config:")
	fmt.Println(projectConfig)
}
