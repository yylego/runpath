package example1

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestValidCase verifies that NewValidCase (using filepath.Dir(runpath.Skip(1)))
// returns the correct absolute path of this package at init time.
//
// TestValidCase 验证 NewValidCase（使用 filepath.Dir(runpath.Skip(1))）
// 在包级初始化时返回本包的正确绝对路径。
func TestValidCase(t *testing.T) {
	t.Log("Path:", ValidCase.Path)
	t.Log("Name:", ValidCase.Name)
	require.True(t, filepath.IsAbs(ValidCase.Path))
	require.Equal(t, "example1", filepath.Base(ValidCase.Path))
	require.Equal(t, "example1", ValidCase.Name)
}

// TestWrongCase demonstrates that NewWrongCase (using runpath.PARENT.Skip(1))
// returns the WRONG path at init time, because the method invocation consumes one extra
// stack frame that the shallow init stack cannot accommodate.
//
// TestWrongCase 演示 NewWrongCase（使用 runpath.PARENT.Skip(1)）
// 在包级初始化时返回错误的路径，因为方法调用多消耗一帧，而浅栈承受不起。
func TestWrongCase(t *testing.T) {
	// WrongCase.Path is WRONG here — it returns "." instead of the absolute path,
	// because PARENT.Skip(1) cannot reach the correct frame in shallow init stacks.
	// WrongCase.Path 在这里是错误的——返回 "." 而非绝对路径，
	// 因为 PARENT.Skip(1) 在浅栈初始化中无法到达正确的帧。
	t.Log("Path:", WrongCase.Path)
	t.Log("Name:", WrongCase.Name)
}
