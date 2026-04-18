package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yylego/runpath/internal/utils"
)

func TestMustAbsPath(t *testing.T) {
	result := utils.MustAbsPath("/abc/def/ghi")
	require.Equal(t, "/abc/def/ghi", result)
}

func TestMustAbsPath_PanicOnNonAbsPath(t *testing.T) {
	require.Panics(t, func() {
		utils.MustAbsPath("abc/def")
	})
}

func TestMustAbsPath_PanicOnDot(t *testing.T) {
	require.Panics(t, func() {
		utils.MustAbsPath(".")
	})
}
