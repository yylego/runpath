package example1

import (
	"path/filepath"

	"github.com/yylego/runpath/internal/examples/example1/locate"
)

// PathInfo simulates a cross-package struct that wraps path info.
// This replicates the scenario where a shared package defines
// a struct and a function, and consuming packages use it via package init variables.
//
// PathInfo 模拟跨包使用的路径信息结构体。
// 复刻实际场景：一个公共包定义结构体和工厂函数，其他包通过包级变量调用。
type PathInfo struct {
	Path string
	Name string
}

// NewValidCase creates PathInfo using filepath.Dir(locate.Skip(1)).
// This is a package function (not a method), and locate.Skip is also a package function.
// The invocation depth is: invocation site -> NewValidCase -> locate.Skip -> runtime.Caller.
// With locate.Skip(1), runtime.Caller(1+1=2) skips exact 2 frames to reach the invocation site. STABLE.
//
// NewValidCase 使用 filepath.Dir(locate.Skip(1)) 创建 PathInfo。
// 这是包级函数（不是方法），locate.Skip 也是包级函数。
// 调用深度：调用方 -> NewValidCase -> locate.Skip -> runtime.Caller。
// locate.Skip(1) 时 runtime.Caller(2) 恰好跳过 2 帧到达调用方。稳定可靠。
func NewValidCase(name string) *PathInfo {
	return &PathInfo{
		Path: filepath.Dir(locate.Skip(1)),
		Name: name,
	}
}

// NewWrongCase creates PathInfo using locate.PARENT.Skip(1).
// locate.PARENT.Skip is a METHOD on parentNamespace, which itself is one stack frame.
// The invocation depth is: invocation site -> NewWrongCase -> PARENT.Skip -> Skip -> runtime.Caller.
// With PARENT.Skip(1), runtime.Caller(1+1+1=3) overshoots in package variable init's shallow stacks.
// This is the exact same issue that runpath.PARENT.Skip(1) had before mustAbsPath was added.
//
// NewWrongCase 使用 locate.PARENT.Skip(1) 创建 PathInfo。
// locate.PARENT.Skip 是 parentNamespace 上的方法，方法调用本身占一帧。
// 调用深度：调用方 -> NewWrongCase -> PARENT.Skip -> Skip -> runtime.Caller。
// PARENT.Skip(1) 时 runtime.Caller(1+1+1=3) 在包级变量初始化的浅栈中跳过头了。
// 这与添加 mustAbsPath 之前的 runpath.PARENT.Skip(1) 问题完全一致。
func NewWrongCase(name string) *PathInfo {
	return &PathInfo{
		Path: locate.PARENT.Skip(1),
		Name: name,
	}
}

// ValidCase uses NewValidCase — the stable approach. Expected: this package's absolute path.
// ValidCase 使用 NewValidCase——稳定的方案。预期：本包的绝对路径。
var ValidCase = NewValidCase("example1")

// WrongCase uses NewWrongCase — demonstrates the shallow-stack issue.
// Expected: NOT this package's path (wrong due to the extra method frame).
// WrongCase 使用 NewWrongCase——演示浅栈问题。
// 预期：不是本包的路径（因为方法调用多占一帧导致结果错误）。
var WrongCase = NewWrongCase("example1")
