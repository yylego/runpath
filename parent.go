// Parent DIR namespace features used in runtime path operations
// Provides namespace-like organization as parent DIR functions
// Enables clean API design with PARENT.Path(), PARENT.Join() patterns
//
// 父 DIR 命名空间功能，用于运行时路径操作
// 为父 DIR 相关函数提供类似命名空间的组织方式
// 实现清晰的 API 设计，支持 PARENT.Path()、PARENT.Join() 模式
package runpath

import (
	"path/filepath"

	"github.com/yylego/runpath/internal/utils"
)

// parentNamespace provides namespace-like organization as parent DIR operations
// Enables tidied function names and more explicit usage patterns
// Improves IDE code completion via narrowing selection scope
//
// parentNamespace 为父 DIR 操作提供类似命名空间的组织方式
// 实现更清晰的函数名和更明确的使用模式
// 通过缩小选择范围来改善 IDE 代码完成
type parentNamespace struct{}

// Namespace instances as parent DIR operations
// Since Go has no namespace concept, this provides organized API access
//
// 父 DIR 操作的全局命名空间实例
// 由于 Go 缺乏命名空间概念，这提供了有组织的 API 访问
var (
	PARENT = &parentNamespace{} // Namespace instance as parent DIR operations // 父 DIR 操作的全局实例
	DIR    = &parentNamespace{} // Concise name as PARENT // PARENT 的同义词，更简单的名称
)

// Path returns the parent DIR path of the invoking source file
// Gets the DIR containing the execution location
// Panics when the outcome is not an absolute path (shallow-stack misuse protection)
//
// NOTE: These methods consume one stack frame due to the method invocation.
// In shallow-stack scenarios (e.g., package init), use
// filepath.Dir(runpath.Skip(n)) instead, which is a package function and
// does not consume an extra frame.
//
// Path 返回调用源文件的父 DIR 路径
// 获取包含当前执行位置的 DIR
// 当结果不是绝对路径时 panic（浅栈误用保护）
//
// 注意：这些方法因方法调用本身会消耗一个栈帧。
// 在栈帧极浅的场景（如包级变量初始化）中，请使用
// filepath.Dir(runpath.Skip(n))，它是包级别函数，不会消耗额外栈帧。
func (T *parentNamespace) Path() string {
	return utils.MustAbsPath(filepath.Dir(Skip(1)))
}

// Name returns the parent DIR name of the invoking source file
// Gets just the name of the DIR containing execution location
// Panics when the resolved path is not absolute (shallow-stack misuse protection)
//
// Name 返回调用源文件的父 DIR 名称
// 仅获取包含当前执行位置的 DIR 名称
// 当解析的路径不是绝对路径时 panic（浅栈误用保护）
func (T *parentNamespace) Name() string {
	return filepath.Base(utils.MustAbsPath(filepath.Dir(Skip(1))))
}

// Skip returns the parent DIR path with specified frame skip count
// Gets parent DIR from different stack depths
// Panics when the outcome is not an absolute path (shallow-stack misuse protection)
//
// Skip 返回指定调用帧跳过的父 DIR 路径
// 允许从不同的调用堆栈级别获取父 DIR
// 当结果不是绝对路径时 panic（浅栈误用保护）
func (T *parentNamespace) Skip(skip int) string {
	return utils.MustAbsPath(filepath.Dir(Skip(1 + skip)))
}

// Join constructs path via joining parent DIR with extra path components
// Builds paths based on the invoking file's parent DIR
// Panics when the resolved path is not absolute (shallow-stack misuse protection)
//
// Join 通过将父 DIR 与额外的路径组件连接来构建路径
// 动态构建相对于调用文件的父 DIR 的路径
// 当解析的路径不是绝对路径时 panic（浅栈误用保护）
func (T *parentNamespace) Join(names ...string) string {
	path := utils.MustAbsPath(filepath.Dir(Skip(1)))
	subs := append([]string{path}, names...)
	return utils.MustAbsPath(filepath.Join(subs...))
}

// Up navigates up the DIR structure from parent DIR
// Goes up specified count of depths from the invoking file's parent DIR
// Panics when the resolved path is not absolute (shallow-stack misuse protection)
//
// Up 从父 DIR 向上导航 DIR 结构
// 从调用文件的父 DIR 向上跳过指定数量的级别
// 当解析的路径不是绝对路径时 panic（浅栈误用保护）
func (T *parentNamespace) Up(skip int) string {
	path := utils.MustAbsPath(filepath.Dir(Skip(1)))
	for i := 0; i < skip; i++ {
		path = filepath.Dir(path)
	}
	return utils.MustAbsPath(path)
}

// UpTo navigates up the DIR structure and joins with extra paths
// Combines Up() and Join() operations in a single invocation
// Panics when the resolved path is not absolute (shallow-stack misuse protection)
//
// UpTo 向上导航 DIR 结构并与额外路径连接
// 在单次调用中组合 Up() 和 Join() 操作
// 当解析的路径不是绝对路径时 panic（浅栈误用保护）
func (T *parentNamespace) UpTo(skip int, names ...string) string {
	path := utils.MustAbsPath(filepath.Dir(Skip(1)))
	for i := 0; i < skip; i++ {
		path = filepath.Dir(path)
	}
	subs := append([]string{path}, names...)
	return utils.MustAbsPath(filepath.Join(subs...))
}
