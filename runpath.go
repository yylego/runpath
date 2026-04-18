// Package runpath: Runtime path and execution location tracking engine
// Provides precise source code location tracking via Go's runtime package
// Supports path manipulation, extension handling, and parent DIR navigation
// Enables dynamic config file path construction based on execution context
//
// runpath: 运行时路径获取和执行位置跟踪引擎
// 通过 Go 的 runtime 包提供精确的源代码位置跟踪
// 支持路径操作、扩展名处理和父 DIR 导航
// 基于执行上下文实现动态配置文件路径构建
package runpath

import (
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
	"github.com/yylego/runpath/internal/utils"
)

// Path returns the runtime source file path at execution location
// Gets the absolute path of the invoking source file
//
// Path 获得运行时的源码文件路径
// 获取调用源文件的绝对路径
func Path() string {
	return utils.MustAbsPath(Skip(1))
}

// Current returns the source file path at the invocation site
// Since the package name is runpath, it means "run path of this code"
// Kept concise and simple to use
//
// Current 获得当前源码文件路径
// 因为包名是 runpath，Current 的含义就是 "current run path"
// 保持简洁便于使用
func Current() string {
	return utils.MustAbsPath(Skip(1))
}

// CurrentPath returns the source file path at the invocation site
// Same as Current, with a more explicit name
//
// CurrentPath 获得当前源码文件路径
// Current 的替代版本，提供给喜欢明确命名的用户
func CurrentPath() string {
	return utils.MustAbsPath(Skip(1))
}

// CurrentName returns the source file name at the invocation site
// Gets just the filename without the DIR path
//
// CurrentName 获得当前源码文件名称
// 仅获取文件名，不包含 DIR 路径
func CurrentName() string {
	return filepath.Base(utils.MustAbsPath(Skip(1)))
}

// Name returns the runtime source file name
// Gets just the filename at execution location
//
// Name 获得运行时的源码文件名称
// 获取执行位置的文件名
func Name() string {
	return filepath.Base(utils.MustAbsPath(Skip(1)))
}

// Skip returns the runtime source location with specified frame skip count
// When skip=0, returns the invoking point's path
// The skip argument represents skips from the invocation position
// Skip(1) equals runtime.Caller(1)'s path at the same position
// Skip(2) equals runtime.Caller(2)'s path at the same position
//
// Skip 获得运行时的源码位置
// 当传0的时候就是调用点的路径
// 核心原则：实参 skip 是调用位置的 skip 次数
// Skip(1) 相当于在相同位置调用 runtime.Caller(1) 的路径
// Skip(2) 相当于在相同位置调用 runtime.Caller(2) 的路径
func Skip(skip int) string {
	_, path, _, ok := runtime.Caller(1 + skip) // +1 to account this function frame // 这里又调用了一层因此这里得补1次
	if !ok {
		panic(errors.Errorf("RUNTIME CALLER FAILED AT SKIP DEPTH %d", 1+skip)) // 因为在99%的场景下都是不会出错的，而且跟获取代码路径相关的逻辑，通常也不会用在线上环境，因此直接 panic
	}
	return utils.MustAbsPath(path)
}

// GetPathChangeExtension changes the source file extension at invocation site
// Removes .go suffix and adds new extension like ".xxx.zzz"
// Common use: in config.go, get config.json path to load configuration
// Can add ".json", "_dev.json", "_uat.json" to match different environments
// This function is the backbone of dynamic config file loading
//
// GetPathChangeExtension 把当前源码的文件路径去除结尾.go，再增加新的结尾
// 可以增加 ".xxx.qqq.zzz" 等任意扩展名
// 常见用途：在 config.go 里获取 config.json 的路径来读取配置
// 可以增加 ".json"、"_dev.json"、"_uat.json" 用于不同环境
// 这个函数对动态配置文件加载非常重要
func GetPathChangeExtension(pointExtension string) string {
	return utils.MustAbsPath(GetSkipRemoveExtension(1) + pointExtension)
}

// GetRex is a concise name of GetPathChangeExtension
// Changes the source file extension with a new one
//
// GetRex 是 GetPathChangeExtension 的简短别名
// 更改当前源文件的扩展名
func GetRex(pointExtension string) string {
	return utils.MustAbsPath(GetSkipRemoveExtension(1) + pointExtension)
}

// GetNox returns the source file path without extension
// Removes the .go suffix from the file path
//
// GetNox 返回不带扩展名的当前源文件路径
// 从当前文件路径中移除 .go 后缀
func GetNox() string {
	return utils.MustAbsPath(GetSkipRemoveExtension(1))
}

// GetPathRemoveExtension removes the .go extension from the source file path
// Less used but kept as a complete API
//
// GetPathRemoveExtension 把当前源码的文件路径去除结尾.go
// 使用频率较低但保留以保持完整性
func GetPathRemoveExtension() string {
	return utils.MustAbsPath(GetSkipRemoveExtension(1))
}

// GetSkipRemoveExtension removes the .go extension with specified frame skip count
//
// GetSkipRemoveExtension 返回指定调用帧跳过的去除 .go 扩展名的源文件路径
func GetSkipRemoveExtension(skip int) string {
	_, path, _, ok := runtime.Caller(1 + skip)
	if !ok {
		panic(errors.Errorf("RUNTIME CALLER FAILED AT SKIP DEPTH %d", 1+skip)) // 因为在99%的场景下都是不会出错的，而且跟获取代码路径相关的逻辑，通常也不会用在线上环境，因此直接 panic
	}
	const extension = ".go"
	if !strings.HasSuffix(strings.ToLower(path), extension) {
		panic(errors.Errorf("EXPECTED %s EXTENSION BUT GOT PATH %s", extension, path))
	}
	return utils.MustAbsPath(path[:len(path)-len(extension)])
}
