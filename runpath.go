// Package runpath: Runtime path retrieval and execution location tracking engine
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
)

// Path returns the runtime source file path at execution location
// Gets the absolute path of the calling source file
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
// See parent.go Path() comments on inlining and frame count predictableness.
//
// Path 获得运行时的源码文件路径
// 获取调用源文件的绝对路径
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
// 关于内联和帧数稳定性的原因，请参见 parent.go 中 Path() 的注释。
func Path() string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("wrong"))
	}
	return path
}

// Current returns the current source file path
// Since the package name is runpath, Current means "current run path"
// Kept concise for ease of use
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
//
// Current 获得当前源码文件路径
// 因为包名是 runpath，Current 的含义就是 "current run path"
// 保持简洁便于使用
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
func Current() string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("wrong"))
	}
	return path
}

// CurrentPath returns the current source file path
// Alternative to Current for those who prefer explicit naming
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
//
// CurrentPath 获得当前源码文件路径
// Current 的替代版本，提供给喜欢明确命名的用户
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
func CurrentPath() string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("wrong"))
	}
	return path
}

// CurrentName returns the current source file name
// Gets just the filename without the DIR path
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
//
// CurrentName 获得当前源码文件名称
// 仅获取文件名，不包含 DIR 路径
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
func CurrentName() string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("wrong"))
	}
	return filepath.Base(path)
}

// Name returns the runtime source file name
// Gets just the filename at execution location
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
//
// Name 获得运行时的源码文件名称
// 获取执行位置的文件名
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
func Name() string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("wrong"))
	}
	return filepath.Base(path)
}

// Skip returns the runtime source location with specified call frame skip
// When skip=0, returns the calling point's path
// Core principle: the skip parameter represents skips from the calling position
// Skip(1) equals runtime.Caller(1)'s path at the same position
// Skip(2) equals runtime.Caller(2)'s path at the same position
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
//
// Skip 获得运行时的源码位置
// 当传0的时候就是调用点的路径
// 核心原则：实参 skip 是调用位置的 skip 次数
// Skip(1) 相当于在相同位置调用 runtime.Caller(1) 的路径
// Skip(2) 相当于在相同位置调用 runtime.Caller(2) 的路径
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
func Skip(skip int) string {
	_, path, _, ok := runtime.Caller(1 + skip) // +1 to account this function frame // 这里又调用了一层因此这里得补1次
	if !ok {
		panic(errors.New("wrong")) // 因为在99%的场景下都是不会出错的，而且跟获取代码路径相关的逻辑，通常也不会用在线上环境，因此直接 panic
	}
	return path
}

// GetPathChangeExtension changes the current source file extension
// Removes .go suffix and adds new extension like ".xxx.yyy.zzz"
// Common use: in config.go, get config.json path to read configuration
// Can add ".json", "_dev.json", "_uat.json" for different environments
// This function is essential for dynamic config file loading
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
//
// GetPathChangeExtension 把当前源码的文件路径去除结尾.go，再增加新的结尾
// 可以增加 ".xxx.yyy.zzz" 等任意扩展名
// 常见用途：在 config.go 里获取 config.json 的路径来读取配置
// 可以增加 ".json"、"_dev.json"、"_uat.json" 用于不同环境
// 这个函数对动态配置文件加载非常重要
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
func GetPathChangeExtension(pointExtension string) string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("wrong"))
	}
	return removeGoExtension(path) + pointExtension
}

// GetRex is a shorter alias for GetPathChangeExtension
// Changes current source file extension with new one
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
//
// GetRex 是 GetPathChangeExtension 的简短别名
// 更改当前源文件的扩展名
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
func GetRex(pointExtension string) string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("wrong"))
	}
	return removeGoExtension(path) + pointExtension
}

// GetNox returns current source file path without extension
// Removes the .go suffix from current file path
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
//
// GetNox 返回不带扩展名的当前源文件路径
// 从当前文件路径中移除 .go 后缀
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
func GetNox() string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("wrong"))
	}
	return removeGoExtension(path)
}

// GetPathRemoveExtension removes the .go extension from current source file path
// Less frequently used but kept for completeness
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
//
// GetPathRemoveExtension 把当前源码的文件路径去除结尾.go
// 使用频率较低但保留以保持完整性
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
func GetPathRemoveExtension() string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("wrong"))
	}
	return removeGoExtension(path)
}

// GetSkipRemoveExtension removes the .go extension with specified frame skip count
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
//
// GetSkipRemoveExtension 返回指定调用帧跳过的去除 .go 扩展名的源文件路径
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
func GetSkipRemoveExtension(skip int) string {
	_, path, _, ok := runtime.Caller(1 + skip)
	if !ok {
		panic(errors.New("wrong"))
	}
	return removeGoExtension(path)
}

// removeGoExtension removes the .go extension from a file path
// This is a pure string operation with no runtime.Caller involvement, reuse is safe.
//
// removeGoExtension 从文件路径中移除 .go 扩展名
// 这是纯字符串操作，不涉及 runtime.Caller，可以安全共享。
func removeGoExtension(path string) string {
	const extension = ".go"
	if !strings.HasSuffix(strings.ToLower(path), extension) {
		panic(errors.Errorf("%s %s", path, extension))
	}
	return path[:len(path)-len(extension)]
}
