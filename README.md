[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/runpath/release.yml?branch=main&label=BUILD)](https://github.com/yylego/runpath/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/runpath)](https://pkg.go.dev/github.com/yylego/runpath)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/runpath/main.svg)](https://coveralls.io/github/yylego/runpath?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23%2C%201.24%2C%201.25-lightgrey.svg)](https://github.com/yylego/runpath)
[![GitHub Release](https://img.shields.io/github/release/yylego/runpath.svg)](https://github.com/yylego/runpath/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/runpath)](https://goreportcard.com/report/github.com/yylego/runpath)

# runpath

`runpath` package provides func to get the execution location of Go code, including the absolute path of the source file.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Installation

Install the package with:

```shell
go get github.com/yylego/runpath
```

## Core Features

`runpath` is a runtime path utils package that obtains the absolute path of Go code execution locations.

**Core Capabilities:**

1. **Runtime Path Access** - Get code execution location's absolute path via `runtime.Caller()`
2. **Parent DIR Operations** - Parent DIR related operations through `PARENT`/`DIR` namespace
3. **Testing Support** - `runtestpath` sub-package designed to support getting source file paths from test files
4. **Path Extension Handling** - Support for changing and removing file extensions

**runpath Advantages:**

Unlike the built-in `filepath.Abs(".")` which doesn't always provide the expected result in certain situations, `runpath` uses Go's `runtime` package to provide precise location tracking, making it valuable when the exact execution path is needed.

## Core API Overview

**Path Operations:**

- `Path()`, `Current()`, `CurrentPath()` - Get current source file absolute path
- `Name()`, `CurrentName()` - Get current source file name
- `Skip(skip int)` - Skip specified levels to get caller path
- `GetPathChangeExtension()`, `GetRex()` - Change file extension
- `GetPathRemoveExtension()`, `GetNox()` - Remove .go extension

**DIR Operations:**

- `PARENT.Path()`, `DIR.Path()` - Get parent DIR path
- `PARENT.Join()`, `DIR.Join()` - Join paths
- `PARENT.Up()`, `DIR.UpTo()` - Navigate up DIR structure

**Test Utilities (runtestpath):**

- `SrcPath(t)` - Get tested source file path (from \_test.go to corresponding .go file)
- `SrcName(t)` - Get tested source file name
- `SrcPathChangeExtension(t, ext)` - Change tested file extension

## Common Use Cases

- **Dynamic Config File Paths** - Read config.json from config.go dynamically
- **Test Code Generation** - Locate source files when generating code in tests
- **Accurate Execution Paths** - More reliable than `filepath.Abs(".")` for execution location

## Example Usage

### Basic Path Operations

This example shows how to get the current source file path and parent DIR information.

```go
package main

import (
	"fmt"

	"github.com/yylego/runpath"
)

func main() {
	// Get current source file path
	currentPath := runpath.Path()
	fmt.Println("Current source path:")
	fmt.Println(currentPath)

	// Get current source file name
	currentName := runpath.Name()
	fmt.Println("Current source name:")
	fmt.Println(currentName)

	// Get parent DIR path
	parentPath := runpath.PARENT.Path()
	fmt.Println("Parent DIR path:")
	fmt.Println(parentPath)

	// Get parent DIR name
	parentName := runpath.PARENT.Name()
	fmt.Println("Parent DIR name:")
	fmt.Println(parentName)
}
```

⬆️ **Source:** [Source](internal/demos/demo1x/main.go)

### DIR Navigation and Path Construction

This example shows how to navigate DIR structures and build paths.

```go
package main

import (
	"fmt"

	"github.com/yylego/runpath"
)

func main() {
	// Join paths with parent DIR
	configPath := runpath.PARENT.Join("config.json")
	fmt.Println("Config path:")
	fmt.Println(configPath)

	dataPath := runpath.DIR.Join("data", "example.txt")
	fmt.Println("Data path:")
	fmt.Println(dataPath)

	// Navigate up DIR structure
	upOne := runpath.PARENT.Up(1)
	fmt.Println("Up 1 DIR:")
	fmt.Println(upOne)

	upTwo := runpath.PARENT.Up(2)
	fmt.Println("Up 2 DIRs:")
	fmt.Println(upTwo)

	// Navigate up and join with path
	moduleConfig := runpath.PARENT.UpTo(1, "config.json")
	fmt.Println("Module config:")
	fmt.Println(moduleConfig)

	projectConfig := runpath.PARENT.UpTo(2, "config", "settings.json")
	fmt.Println("Project config:")
	fmt.Println(projectConfig)
}
```

⬆️ **Source:** [Source](internal/demos/demo2x/main.go)

### File Extension Handling

This example demonstrates changing file extensions and working with config files.

```go
package main

import (
	"fmt"

	"github.com/yylego/runpath"
)

func main() {
	// Get current file path with different extension
	jsonPath := runpath.GetPathChangeExtension(".json")
	fmt.Println("JSON config path:")
	fmt.Println(jsonPath)

	// Using compact alias GetRex
	yamlPath := runpath.GetRex(".yaml")
	fmt.Println("YAML config path:")
	fmt.Println(yamlPath)

	// Get path without extension
	basePath := runpath.GetNox()
	fmt.Println("Base path (no extension):")
	fmt.Println(basePath)

	// Demonstrate use case: reading config file with same name
	configPath := runpath.GetRex(".config.json")
	fmt.Println("Config file path:")
	fmt.Println(configPath)
}
```

⬆️ **Source:** [Source](internal/demos/demo3x/main.go)

---

## Working with Configuration Files

You can use `runpath` to build paths to configuration files based on the **execution location** of the code. This is valuable in tests where different configurations are loaded depending on where the test is being executed.

```go
path := runpath.DIR.Join("config.json")
```

This constructs the path to `config.json` relative to the **execution DIR abs-path**.

You can also use `PARENT` with the same approach:

```go
path := runpath.PARENT.Join("config.json")
```

If you need to navigate up multiple DIR levels from the execution location, use:

```go
path := runpath.PARENT.UpTo(3, "config.json")
```

This will return the path to `config.json` located three levels up from the **execution DIR abs-path**.

---

## Locating Source Code in Test Cases

When running tests, you might need to generate code and reference the source file path. Here's how you can locate the test file's source path:

```go
func TestSrcPath(t *testing.T) {
    path := runpath.SrcPath(t)
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/utils_runtestpath.go"))
}
```

This approach helps when generating code that needs to be placed alongside the source files based on the **execution location**.

---

## Changing File Extensions Based on Context

You can also change the file extension depending on the test context (e.g., from `.go` to `.json`):

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

This allows you to load different types of files (e.g., configuration files) based on the **execution location** and test requirements.

---

## Function Overview

- `Path()`: Returns the absolute path of the source file where the code is executed, representing the **execution location**.
- `Current()`, `CurrentPath()`, `CurrentName()`, `Name()`: Variants of `Path()` that retrieve the file path or name based on the current execution context.
- `Skip(int)`: Retrieves the path from a specified stack frame, helpful when getting the execution location of the calling code.
- `GetPathChangeExtension()`: Returns the current source file path with a new extension (e.g., changing `.go` to `.json`).
- `GetPathRemoveExtension()`: Returns the current source file path without the `.go` extension.
- `Join()`: Joins the current DIR abs-path with extra path components, constructing paths based on the **execution location**.
- `Up()`, `UpTo()`: Navigates up the DIR structure a specified amount of levels from the **execution location**.

---

## Test-Specific Operations

- `SrcPath(t *testing.T)`: Retrieves the source path of the file being tested.
- `SrcName(t *testing.T)`: Retrieves the name of the source file being tested.
- `SrcPathChangeExtension(t *testing.T, ext string)`: Changes the extension of the test file path (e.g., from `.go` to `.json`).
- `SrcSkipRemoveExtension(t *testing.T)`: Removes the `.go` extension from the test file path.

This package is valuable in test files, where you need to reference source code paths and configuration files based on the **execution location** of the test file.

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-20 04:26:32.402216 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![starring](https://starchart.cc/yylego/runpath.svg?variant=adaptive)](https://starchart.cc/yylego/runpath)

Give me stars! Thank you!!!

---
