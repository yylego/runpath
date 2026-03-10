package runpath

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_parentNamespace_Path tests the PARENT.Path function returns parent DIR path
// Test_parentNamespace_Path 测试 PARENT.Path 函数返回父 DIR 路径
func Test_parentNamespace_Path(t *testing.T) {
	t.Log(PARENT.Path())
}

// Test_parentNamespace_Name tests the PARENT.Name function returns parent DIR name
// Test_parentNamespace_Name 测试 PARENT.Name 函数返回父 DIR 名称
func Test_parentNamespace_Name(t *testing.T) {
	name := PARENT.Name()
	t.Log(name)
	require.Equal(t, "runpath", name)
}

// Test_parentNamespace_Skip tests the PARENT.Skip function with different skip depths
// Test_parentNamespace_Skip 测试 PARENT.Skip 函数使用不同的跳过深度
func Test_parentNamespace_Skip(t *testing.T) {
	t.Log(PARENT.Skip(0))
}

// Test_parentNamespace_Join tests the PARENT.Join function joins paths with parent DIR
// Test_parentNamespace_Join 测试 PARENT.Join 函数与父 DIR 连接路径
func Test_parentNamespace_Join(t *testing.T) {
	t.Log(PARENT.Join("example.json"))
}

// Test_parentNamespace_Join1 tests constructing paths using DIR.Path and filepath.Join
// Test_parentNamespace_Join1 测试使用 DIR.Path 和 filepath.Join 构建路径
func Test_parentNamespace_Join1(t *testing.T) {
	name := Name()
	t.Log(name)
	root := DIR.Path() // Variable named "root" for clean code aesthetics // 变量名为 "root" 保持代码美学
	t.Log(root)
	path := filepath.Join(root, name)
	t.Log(path)
	want := Path()
	t.Log(want)
	require.Equal(t, want, path)
}

// Test_parentNamespace_Join2 tests DIR.Join produces same result as filepath.Join
// Test_parentNamespace_Join2 测试 DIR.Join 与 filepath.Join 产生相同结果
func Test_parentNamespace_Join2(t *testing.T) {
	name := Name()
	t.Log(name)
	path := DIR.Join(name)
	t.Log(path)
	want := Path()
	t.Log(want)
	require.Equal(t, want, path)
}

// Test_parentNamespace_Up tests DIR.Up navigates up DIR structure
// Test_parentNamespace_Up 测试 DIR.Up 向上导航 DIR 结构
func Test_parentNamespace_Up(t *testing.T) {
	for depth := 0; depth < 10; depth++ {
		t.Log(DIR.Up(depth))
	}
}

// Test_parentNamespace_UpTo tests DIR.UpTo navigates up and joins paths
// Test_parentNamespace_UpTo 测试 DIR.UpTo 向上导航并连接路径
func Test_parentNamespace_UpTo(t *testing.T) {
	name := DIR.Name()
	t.Log(name)
	path := DIR.UpTo(1, name)
	t.Log(path)
	want := DIR.Path()
	t.Log(want)
	require.Equal(t, want, path)
}
