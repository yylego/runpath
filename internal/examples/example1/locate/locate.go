// Package locate replicates runpath's parentNamespace and Skip function without mustAbsPath protection.
// The code is copied verbatim from runpath (package name changed to locate).
// This package exists to demonstrate the shallow-stack issue: locate.PARENT.Skip(1) produces
// the same wrong result as runpath.PARENT.Skip(1) did before mustAbsPath was added.
//
// locate 从 runpath 原封不动复刻了 parentNamespace 和 Skip 函数，没有 mustAbsPath 保护。
// 代码从 runpath 逐字复制（仅包名改为 locate）。
// 此包用于演示浅栈问题：locate.PARENT.Skip(1) 产生的错误结果
// 与添加 mustAbsPath 之前的 runpath.PARENT.Skip(1) 完全相同。
package locate

import (
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
)

type parentNamespace struct{}

var (
	PARENT = &parentNamespace{}
)

// Skip returns the parent DIR path with specified call frame skip
// Allows getting parent DIR from different call stack levels
//
// Skip 返回指定调用帧跳过的父 DIR 路径
// 允许从不同的调用堆栈级别获取父 DIR
func (T *parentNamespace) Skip(skip int) string {
	return filepath.Dir(Skip(1 + skip))
}

// Skip returns the runtime source location with specified call frame skip
// When skip=0, returns the calling point's path
// Core principle: the skip parameter represents skips from the calling position
// Skip(1) equals runtime.Caller(1)'s path at the same position
// Skip(2) equals runtime.Caller(2)'s path at the same position
//
// Skip 获得运行时的源码位置
// 当传0的时候就是调用点的路径
// 核心原则：实参 skip 是调用位置的 skip 次数
// Skip(1) 相当于在相同位置调用 runtime.Caller(1) 的路径
// Skip(2) 相当于在相同位置调用 runtime.Caller(2) 的路径
func Skip(skip int) string {
	_, path, _, ok := runtime.Caller(1 + skip) // Add 1 to account for this function call // 这里又调用了一层因此这里得补1次
	if !ok {
		panic(errors.New("wrong")) // Panic since this rarely fails and path retrieval is not used in production // 因为在99%的场景下都是不会出错的，而且跟获取代码路径相关的逻辑，通常也不会用在线上环境，因此直接 panic
	}
	return path
}
