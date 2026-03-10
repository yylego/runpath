[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/runpath/release.yml?branch=main&label=BUILD)](https://github.com/yylego/runpath/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/runpath)](https://pkg.go.dev/github.com/yylego/runpath)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/runpath/main.svg)](https://coveralls.io/github/yylego/runpath?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23%2C%201.24%2C%201.25-lightgrey.svg)](https://github.com/yylego/runpath)
[![GitHub Release](https://img.shields.io/github/release/yylego/runpath.svg)](https://github.com/yylego/runpath/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/runpath)](https://goreportcard.com/report/github.com/yylego/runpath)

# runpath

`runpath` 包提供了获取 Go 代码执行位置的功能，包括当前源文件的绝对路径。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 安装

使用以下命令安装该包：

```shell
go get github.com/yylego/runpath
```

## 核心特性

`runpath` 是一个运行时路径获取工具包，用于获取 Go 代码执行位置的绝对路径。

**核心功能：**

1. **运行时路径访问** - 通过 `runtime.Caller()` 获取代码执行位置的绝对路径
2. **父 DIR 操作** - 通过 `PARENT`/`DIR` 命名空间提供父 DIR 相关操作
3. **测试支持** - `runtestpath` 子包专为支持测试文件中获取被测试源文件路径
4. **路径扩展处理** - 支持文件扩展名的更改和移除操作

**runpath 优势：**

不同于内置的 `filepath.Abs(".")` 在某些情况下无法提供预期结果，`runpath` 利用 Go 的 `runtime` 包提供精确的位置跟踪，在需要准确执行路径时特别有用。

## 核心 API 概览

**路径操作：**
- `Path()`, `Current()`, `CurrentPath()` - 获取当前源文件绝对路径
- `Name()`, `CurrentName()` - 获取当前源文件名称  
- `Skip(skip int)` - 跳过指定层级获取调用者路径
- `GetPathChangeExtension()`, `GetRex()` - 更改文件扩展名
- `GetPathRemoveExtension()`, `GetNox()` - 移除 .go 扩展名

**DIR 操作：**
- `PARENT.Path()`, `DIR.Path()` - 获取父 DIR 路径
- `PARENT.Join()`, `DIR.Join()` - 拼接路径
- `PARENT.Up()`, `DIR.UpTo()` - 向上导航 DIR 结构

**测试工具 (runtestpath)：**
- `SrcPath(t)` - 获取被测试源文件路径（从 _test.go 得到对应 .go 文件）
- `SrcName(t)` - 获取被测试源文件名称
- `SrcPathChangeExtension(t, ext)` - 更改被测试文件扩展名

## 常见使用场景

- **动态配置文件路径** - 在 config.go 中动态读取 config.json
- **测试代码生成** - 在测试中生成代码时定位源文件位置
- **准确的执行路径** - 比 `filepath.Abs(".")` 更可靠的执行位置获取

## 示例用法

### 基本路径操作

此示例展示如何获取当前源文件路径和父 DIR 信息。

```go
package main

import (
	"fmt"

	"github.com/yylego/runpath"
)

func main() {
	// 获取当前源文件路径
	currentPath := runpath.Path()
	fmt.Println("当前源文件路径:")
	fmt.Println(currentPath)

	// 获取当前源文件名称
	currentName := runpath.Name()
	fmt.Println("当前源文件名称:")
	fmt.Println(currentName)

	// 获取父 DIR 路径
	parentPath := runpath.PARENT.Path()
	fmt.Println("父 DIR 路径:")
	fmt.Println(parentPath)

	// 获取父 DIR 名称
	parentName := runpath.PARENT.Name()
	fmt.Println("父 DIR 名称:")
	fmt.Println(parentName)
}
```

⬆️ **源码:** [源码](internal/demos/demo1x/main.go)

### DIR 导航和路径构建

此示例展示如何导航 DIR 结构并构建路径。

```go
package main

import (
	"fmt"

	"github.com/yylego/runpath"
)

func main() {
	// 与父 DIR 连接路径
	configPath := runpath.PARENT.Join("config.json")
	fmt.Println("配置路径:")
	fmt.Println(configPath)

	dataPath := runpath.DIR.Join("data", "example.txt")
	fmt.Println("数据路径:")
	fmt.Println(dataPath)

	// 向上导航 DIR 结构
	upOne := runpath.PARENT.Up(1)
	fmt.Println("向上1层:")
	fmt.Println(upOne)

	upTwo := runpath.PARENT.Up(2)
	fmt.Println("向上2层:")
	fmt.Println(upTwo)

	// 向上导航并连接路径
	moduleConfig := runpath.PARENT.UpTo(1, "config.json")
	fmt.Println("模块配置:")
	fmt.Println(moduleConfig)

	projectConfig := runpath.PARENT.UpTo(2, "config", "settings.json")
	fmt.Println("项目配置:")
	fmt.Println(projectConfig)
}
```

⬆️ **源码:** [源码](internal/demos/demo2x/main.go)

### 文件扩展名处理

此示例演示如何更改文件扩展名并处理配置文件。

```go
package main

import (
	"fmt"

	"github.com/yylego/runpath"
)

func main() {
	// 获取当前文件的不同扩展名路径
	jsonPath := runpath.GetPathChangeExtension(".json")
	fmt.Println("JSON 配置路径:")
	fmt.Println(jsonPath)

	// 使用简洁的别名 GetRex
	yamlPath := runpath.GetRex(".yaml")
	fmt.Println("YAML 配置路径:")
	fmt.Println(yamlPath)

	// 获取不带扩展名的路径
	basePath := runpath.GetNox()
	fmt.Println("基础路径（无扩展名）:")
	fmt.Println(basePath)

	// 演示用例：读取同名配置文件
	configPath := runpath.GetRex(".config.json")
	fmt.Println("配置文件路径:")
	fmt.Println(configPath)
}
```

⬆️ **源码:** [源码](internal/demos/demo3x/main.go)

---

## 配置文件路径管理

你可以利用 `runpath` 构建配置文件的路径，在测试中需要根据执行位置加载不同配置时很有价值。例如：

```go
path := runpath.DIR.Join("config.json")
```

该方法通过当前执行位置的 DIR 构建配置文件路径。

你也可以使用 `PARENT` 采用相同的方式：

```go
path := runpath.PARENT.Join("config.json")
```

如果你需要向上导航多个目录层级，可以使用：

```go
path := runpath.PARENT.UpTo(3, "config.json")
```

这将返回类似 `filepath.Join(runpath.PARENT.Path(), "../../..", "config.json")` 的路径。

---

## 在测试中定位源代码路径

在测试用例中，可能需要生成代码并引用源代码路径，可以使用以下方法：

```go
func TestSrcPath(t *testing.T) {
    path := runpath.SrcPath(t)
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/utils_runtestpath.go"))
}
```

这种方式特别适用于生成的代码需要与源文件一起放置的情况。

---

## 修改文件扩展名以适应不同的测试场景

你还可以根据测试场景动态更改文件扩展名，例如从 `.go` 转换为 `.json`：

```go
func TestSrcPathChangeExtension(t *testing.T) {
    path := runpath.SrcPathChangeExtension(t, ".json")
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/utils_runtestpath.json"))
}

func TestSrcRex(t *testing.T) {
    path := SrcRex(t, ".json")
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath.json"))
}
```

这样你就可以根据不同的测试需求加载不同格式的文件。

---

## 函数概览

- `Path()`: 获取当前源文件的绝对路径，表示当前执行的位置。
- `Current()`, `CurrentPath()`, `CurrentName()`, `Name()`: `Path()` 的不同变体，用于获取当前文件的路径或文件名。
- `Skip(int)`: 从指定的堆栈帧获取路径（用于获取调用代码的执行位置）。
- `GetPathChangeExtension()`: 返回当前源文件的路径，并可以更改文件扩展名（如将 `.go` 转为 `.json`）。
- `GetPathRemoveExtension()`: 返回当前源文件的路径，但去掉 `.go` 扩展名。
- `Join()`: 将当前 DIR 与其他路径组件连接，构建路径。
- `Up()`, `UpTo()`: 向上导航指定数量的 DIR 结构，并获取对应的路径。

---

## 测试文件特定操作

- `SrcPath(t *testing.T)`: 获取测试文件的源路径。
- `SrcName(t *testing.T)`: 获取测试文件的文件名。
- `SrcPathChangeExtension(t *testing.T, ext string)`: 修改测试文件路径的扩展名（例如将 `.go` 更改为 `.json`）。
- `SrcSkipRemoveExtension(t *testing.T)`: 移除测试文件路径中的 `.go` 扩展名。

此包特别适用于测试文件中，帮助你根据测试文件所在位置引用源代码和配置文件路径。

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-20 04:26:32.402216 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/yylego/runpath.svg?variant=adaptive)](https://starchart.cc/yylego/runpath)

Give me stars! Thank you!!!

---
