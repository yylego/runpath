package runpath

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestPath tests the Path function returns current source file path
// TestPath 测试 Path 函数返回当前源文件路径
func TestPath(t *testing.T) {
	path := Path()
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runpath_test.go"))
}

// TestCurrent tests the Current function returns current source file path
// TestCurrent 测试 Current 函数返回当前源文件路径
func TestCurrent(t *testing.T) {
	path := Current()
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runpath_test.go"))
}

// TestCurrentPath tests the CurrentPath function returns current source file path
// TestCurrentPath 测试 CurrentPath 函数返回当前源文件路径
func TestCurrentPath(t *testing.T) {
	path := CurrentPath()
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runpath_test.go"))
}

// TestCurrentName tests the CurrentName function returns current source file name
// TestCurrentName 测试 CurrentName 函数返回当前源文件名称
func TestCurrentName(t *testing.T) {
	name := CurrentName()
	t.Log(name)
	require.Equal(t, "runpath_test.go", name)
}

// TestName tests the Name function returns current source file name
// TestName 测试 Name 函数返回当前源文件名称
func TestName(t *testing.T) {
	name := Name()
	t.Log(name)
	require.Equal(t, "runpath_test.go", name)
}

// TestSkip tests the Skip function with different skip depths
// TestSkip 测试 Skip 函数使用不同的跳过深度
func TestSkip(t *testing.T) {
	for skp := 0; skp <= 10; skp++ {
		t.Log("-----------------------")
		pc, path, pos, ok := runtime.Caller(skp)
		if !ok {
			t.Log(skp, pc, path, pos, ok)
			return
		} else {
			t.Log(skp, pc, path, pos, ok)
			require.Equal(t, path, Skip(skp))
		}
	}
}

// TestGetPathChangeExtension tests changing file extension from .go to .json
// TestGetPathChangeExtension 测试将文件扩展名从 .go 改为 .json
func TestGetPathChangeExtension(t *testing.T) {
	path := GetPathChangeExtension(".json")
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runpath_test.json"))
}

// TestGetRex tests the GetRex function alias changes file extension
// TestGetRex 测试 GetRex 函数别名更改文件扩展名
func TestGetRex(t *testing.T) {
	path := GetRex(".json")
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runpath_test.json"))
}

// TestGetNox tests removing file extension from path
// TestGetNox 测试从路径中移除文件扩展名
func TestGetNox(t *testing.T) {
	path := GetNox()
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runpath_test"))
}

// TestGetPathRemoveExtension tests removing .go extension from path
// TestGetPathRemoveExtension 测试从路径中移除 .go 扩展名
func TestGetPathRemoveExtension(t *testing.T) {
	path := GetPathRemoveExtension()
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runpath_test"))
}

// TestGetSkipRemoveExtension tests removing extension with skip depth
// TestGetSkipRemoveExtension 测试使用跳过深度移除扩展名
func TestGetSkipRemoveExtension(t *testing.T) {
	path := GetSkipRemoveExtension(0)
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runpath_test"))
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

// TestSkip_PanicOnDeepSkip tests that Skip panics when skip depth exceeds stack
// TestSkip_PanicOnDeepSkip 测试当跳过深度超出栈帧时 Skip 会 panic
func TestSkip_PanicOnDeepSkip(t *testing.T) {
	require.Panics(t, func() {
		Skip(9999)
	})
}

// TestGetSkipRemoveExtension_PanicOnDeepSkip tests panic on excessive skip depth
// TestGetSkipRemoveExtension_PanicOnDeepSkip 测试跳过深度过大时 panic
func TestGetSkipRemoveExtension_PanicOnDeepSkip(t *testing.T) {
	require.Panics(t, func() {
		GetSkipRemoveExtension(9999)
	})
}
