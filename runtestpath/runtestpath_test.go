package runtestpath

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestSrcPath tests the SrcPath function returns source file path from test file
// TestSrcPath 测试 SrcPath 函数从测试文件返回源文件路径
func TestSrcPath(t *testing.T) {
	path := SrcPath(t)
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath.go"))
}

// TestSrcName tests the SrcName function returns source file name from test file
// TestSrcName 测试 SrcName 函数从测试文件返回源文件名
func TestSrcName(t *testing.T) {
	name := SrcName(t)
	t.Log(name)
	require.Equal(t, "runtestpath.go", name)
}

// TestSrcSkip tests the SrcSkip function with specified skip depth
// TestSrcSkip 测试 SrcSkip 函数使用指定的跳过深度
func TestSrcSkip(t *testing.T) {
	path := SrcSkip(t, 0)
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath.go"))
}

// TestSrcPathChangeExtension tests changing source file extension from test context
// TestSrcPathChangeExtension 测试从测试上下文中更改源文件扩展名
func TestSrcPathChangeExtension(t *testing.T) {
	path := SrcPathChangeExtension(t, ".json")
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath.json"))
}

// TestSrcRex tests the SrcRex name used to change source file extension
// TestSrcRex 测试 SrcRex 别名用于更改源文件扩展名
func TestSrcRex(t *testing.T) {
	path := SrcRex(t, ".json")
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath.json"))
}

// TestSrcNox tests getting source file path without extension from test context
// TestSrcNox 测试从测试上下文中获取不带扩展名的源文件路径
func TestSrcNox(t *testing.T) {
	path := SrcNox(t)
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath"))
}

// TestSrcPathRemoveExtension tests removing extension from source file path in test context
// TestSrcPathRemoveExtension 测试在测试上下文中从源文件路径中移除扩展名
func TestSrcPathRemoveExtension(t *testing.T) {
	path := SrcPathRemoveExtension(t)
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath"))
}

// TestSrcSkipRemoveExtension tests removing extension with specified skip depth
// TestSrcSkipRemoveExtension 测试使用指定跳过深度移除扩展名
func TestSrcSkipRemoveExtension(t *testing.T) {
	path := SrcSkipRemoveExtension(t, 0)
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath"))
}

// TestAbsPath tests filepath.Abs function outcome
// This method works but might give the project root in some cases
//
// TestAbsPath 测试 filepath.Abs 函数行为
// 这个方法可行但在某些情况下可能返回项目根目录
func TestAbsPath(t *testing.T) {
	path, err := filepath.Abs(".")
	require.NoError(t, err)
	t.Log(path) // Might give different results depending on context // 根据上下文可能返回不同结果
}

// TestOsGetWD tests os.Getwd function outcome
// Returns working DIR which might not match the source file location
//
// TestOsGetWD 测试 os.Getwd 函数行为
// 返回工作 DIR 可能与源文件位置不同
func TestOsGetWD(t *testing.T) {
	path, err := os.Getwd()
	require.NoError(t, err)
	t.Log(path) // Working DIR path // 工作 DIR 路径
}
