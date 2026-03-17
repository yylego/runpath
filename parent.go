// Parent DIR namespace functionality for runtime path operations
// Provides namespace-like organization for parent DIR related functions
// Enables clean API design with PARENT.Path(), PARENT.Join() patterns
//
// 父 DIR 命名空间功能，用于运行时路径操作
// 为父 DIR 相关函数提供类似命名空间的组织方式
// 实现清晰的 API 设计，支持 PARENT.Path()、PARENT.Join() 模式
package runpath

import (
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
)

// parentNamespace provides namespace-like organization for parent DIR operations
// Enables cleaner function names and more explicit usage patterns
// Improves IDE code completion by narrowing selection scope
//
// parentNamespace 为父 DIR 操作提供类似命名空间的组织方式
// 实现更清晰的函数名和更明确的使用模式
// 通过缩小选择范围来改善 IDE 代码完成
type parentNamespace struct{}

// Global namespace instances for parent DIR operations
// Users are trusted not to set these to nil (that would be self-defeating)
// Since Go lacks namespace concept, this provides organized API access
//
// 父 DIR 操作的全局命名空间实例
// 相信用户不会将其设置为 nil（那将是自找麻烦）
// 由于 Go 缺乏命名空间概念，这提供了有组织的 API 访问
var (
	PARENT = &parentNamespace{} // Global instance for parent DIR operations // 父 DIR 操作的全局实例
	DIR    = &parentNamespace{} // Synonym for PARENT, simpler name // PARENT 的同义词，更简单的名称
)

// Path returns the parent DIR path of the calling source file
// Gets the DIR containing the current execution location
//
// WARNING: DO NOT extract runtime.Caller into a shared function.
// runtime.Caller counts stack frames, and each intermediate function (even within
// the same package) can be inlined during compilation, making the frame count
// unpredictable. Each method must invoke runtime.Caller right here with a hardcoded
// skip to guarantee correct and predictable results across Go versions.
//
// Path 返回调用源文件的父 DIR 路径
// 获取包含当前执行位置的 DIR
//
// 警告：不要将 runtime.Caller 提取到共享的辅助函数中。
// runtime.Caller 依赖栈帧计数，任何中间函数（即使在同一个包内）都可能被编译器内联，
// 导致帧数不可预测。每个方法必须直接调用 runtime.Caller 并使用硬编码的 skip 值，
// 以确保在所有 Go 版本下都能获得正确且稳定的结果。
func (T *parentNamespace) Path() string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("wrong"))
	}
	return filepath.Dir(path)
}

// Name returns the parent DIR name of the calling source file
// Gets just the name of the DIR containing current execution location
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
// See Path() comments on inlining and frame count predictableness.
//
// Name 返回调用源文件的父 DIR 名称
// 仅获取包含当前执行位置的 DIR 名称
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
// 关于内联和帧数稳定性的原因，请参见 Path() 的注释。
func (T *parentNamespace) Name() string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("wrong"))
	}
	return filepath.Base(filepath.Dir(path))
}

// Skip returns the parent DIR path with specified call frame skip
// Allows getting parent DIR from different call stack levels
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
// See Path() comments on inlining and frame count predictableness.
//
// Skip 返回指定调用帧跳过的父 DIR 路径
// 允许从不同的调用堆栈级别获取父 DIR
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
// 关于内联和帧数稳定性的原因，请参见 Path() 的注释。
func (T *parentNamespace) Skip(skip int) string {
	_, path, _, ok := runtime.Caller(1 + skip)
	if !ok {
		panic(errors.New("wrong"))
	}
	return filepath.Dir(path)
}

// Join constructs path by joining parent DIR with additional path components
// Dynamically builds paths relative to the calling file's parent DIR
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
// See Path() comments on inlining and frame count predictableness.
//
// Join 通过将父 DIR 与额外的路径组件连接来构建路径
// 动态构建相对于调用文件的父 DIR 的路径
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
// 关于内联和帧数稳定性的原因，请参见 Path() 的注释。
func (T *parentNamespace) Join(names ...string) string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("wrong"))
	}
	return filepath.Join(append([]string{filepath.Dir(path)}, names...)...)
}

// Up navigates up the DIR structure from parent DIR
// Goes up specified number of levels from the calling file's parent DIR
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
// See Path() comments on inlining and frame count predictableness.
//
// Up 从父 DIR 向上导航 DIR 结构
// 从调用文件的父 DIR 向上跳过指定数量的级别
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
// 关于内联和帧数稳定性的原因，请参见 Path() 的注释。
func (T *parentNamespace) Up(skip int) string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("wrong"))
	}
	root := filepath.Dir(path)
	for i := 0; i < skip; i++ {
		root = filepath.Dir(root)
	}
	return root
}

// UpTo navigates up the DIR structure and joins with additional paths
// Combines Up() and Join() operations in a single call
//
// WARNING: DO NOT restructure — invoke runtime.Caller right here, not via an intermediate function.
// See Path() comments on inlining and frame count predictableness.
//
// UpTo 向上导航 DIR 结构并与额外路径连接
// 在单次调用中组合 Up() 和 Join() 操作
//
// 警告：不要重构——runtime.Caller 必须在此处直接调用，不能通过任何包装函数。
// 关于内联和帧数稳定性的原因，请参见 Path() 的注释。
func (T *parentNamespace) UpTo(skip int, names ...string) string {
	_, path, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("wrong"))
	}
	root := filepath.Dir(path)
	for i := 0; i < skip; i++ {
		root = filepath.Dir(root)
	}
	return filepath.Join(append([]string{root}, names...)...)
}
