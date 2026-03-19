package utils

import (
	"path/filepath"

	"github.com/pkg/errors"
)

// MustAbsPath panics when the path is not absolute.
// This catches shallow-stack misuse (e.g., package variable init) at once,
// instead of returning "." which could cause catastrophic mistakes.
//
// MustAbsPath 当路径不是绝对路径时 panic。
// 立即捕获浅栈误用（如包级变量初始化），
// 而不是悄悄返回 "." 导致灾难性错误。
func MustAbsPath(path string) string {
	if !filepath.IsAbs(path) {
		panic(errors.Errorf("EXPECTED ABSOLUTE PATH BUT GOT %q", path))
	}
	return path
}
