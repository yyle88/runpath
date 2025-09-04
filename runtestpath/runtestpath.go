// Package runtestpath: Test-specific runtime path utilities for source file discovery
// Enables finding source files from test files during test execution
// For test file path, use runpath.Path() directly
// This package specializes in finding source files being tested
// All functions require *testing.T parameter to indicate test-only usage
//
// runtestpath: 测试专用运行时路径工具，用于源文件发现
// 在测试执行期间从测试文件中查找源文件
// 获取测试文件路径请直接使用 runpath.Path()
// 这个包专门用于查找被测试的源文件
// 所有函数都需要 *testing.T 参数来表示仅限测试使用
package runtestpath

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// SrcPath returns the source file path corresponding to the test file
// If your test file is abc_test.go, returns /aa/bb/cc/abc.go
// Maps from test file to the source file being tested
//
// SrcPath 返回与测试文件对应的源文件路径
// 如果你的测试文件是 abc_test.go，返回 /aa/bb/cc/abc.go
// 从测试文件映射到被测试的源文件
func SrcPath(t *testing.T) string {
	return SrcSkip(t, 1)
}

// SrcName returns the source file name corresponding to the test file
// If your test file is abc_test.go, returns abc.go
// Gets just the filename of the source file being tested
//
// SrcName 返回与测试文件对应的源文件名
// 如果你的测试文件是 abc_test.go，返回 abc.go
// 仅获取被测试源文件的文件名
func SrcName(t *testing.T) string {
	return filepath.Base(SrcSkip(t, 1))
}

// SrcSkip returns source file path with specified call frame skip
// Currently has limited use cases beyond the above two scenarios
// Exported in case there are other external uses
//
// SrcSkip 返回指定调用帧跳过的源文件路径
// 目前除了上述两个场景外使用场景有限
// 导出以防有其他外部用途
func SrcSkip(t *testing.T, skip int) string {
	_, path, _, ok := runtime.Caller(1 + skip)           //这里又调用了一层因此这里得补1次
	require.True(t, ok)                                  // Ensure runtime.Caller succeeds // 确保 runtime.Caller 成功
	require.True(t, strings.HasSuffix(path, "_test.go")) // Verify this is a test file // 验证这是一个测试文件
	return path[:len(path)-len("_test.go")] + ".go"      // Convert _test.go to .go // 将 _test.go 转换为 .go
}

// SrcPathChangeExtension changes the source file extension from test context
// If your test file is abc_test.go and you pass ".json", returns /aa/bb/cc/abc.json
// Gets config files with same name as the source file being tested
//
// SrcPathChangeExtension 从测试上下文中更改源文件扩展名
// 如果你的测试文件是 abc_test.go 且传递 ".json"，返回 /aa/bb/cc/abc.json
// 获取与被测试源文件同名的配置文件
func SrcPathChangeExtension(t *testing.T, pointExtension string) string {
	return SrcSkipRemoveExtension(t, 1) + pointExtension
}

// SrcRex is a shorter alias for SrcPathChangeExtension
// Changes source file extension with new one from test context
//
// SrcRex 是 SrcPathChangeExtension 的简短别名
// 从测试上下文中更改源文件扩展名
func SrcRex(t *testing.T, pointExtension string) string {
	return SrcSkipRemoveExtension(t, 1) + pointExtension
}

// SrcNox returns source file path without extension from test context
// Removes extension from the source file corresponding to test file
//
// SrcNox 从测试上下文中返回不带扩展名的源文件路径
// 从与测试文件对应的源文件中移除扩展名
func SrcNox(t *testing.T) string {
	return SrcSkipRemoveExtension(t, 1)
}

// SrcPathRemoveExtension removes extension from source file path in test context
// Gets source file path without .go extension
//
// SrcPathRemoveExtension 在测试上下文中从源文件路径中移除扩展名
// 获取不带 .go 扩展名的源文件路径
func SrcPathRemoveExtension(t *testing.T) string {
	return SrcSkipRemoveExtension(t, 1)
}

func SrcSkipRemoveExtension(t *testing.T, skip int) string {
	_, path, _, ok := runtime.Caller(1 + skip)
	require.True(t, ok)                                  // Ensure runtime.Caller succeeds // 确保 runtime.Caller 成功
	require.True(t, strings.HasSuffix(path, "_test.go")) // Verify this is a test file // 验证这是一个测试文件
	return path[:len(path)-len("_test.go")]              // Remove _test.go suffix // 移除 _test.go 后缀
}
