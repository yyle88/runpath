// Package runpath: Runtime path retrieval and execution location tracking engine
// Provides precise source code location tracking via Go's runtime package
// Supports path manipulation, extension handling, and parent DIR navigation
// Enables dynamic config file path construction based on execution context
//
// runpath: 运行时路径获取和执行位置跟踪引擎
// 通过 Go 的 runtime 包提供精确的源代码位置跟踪
// 支持路径操作、扩展名处理和父 DIR 导航
// 基于执行上下文实现动态配置文件路径构建
package runpath

import (
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

// Path returns the runtime source file path at execution location
// Gets the absolute path of the calling source file
//
// Path 获得运行时的源码文件路径
// 获取调用源文件的绝对路径
func Path() string {
	return Skip(1)
}

// Current returns the current source file path
// Since the package name is runpath, Current means "current run path"
// Kept concise for ease of use
//
// Current 获得当前源码文件路径
// 因为包名是 runpath，Current 的含义就是 "current run path"
// 保持简洁便于使用
func Current() string {
	return Skip(1)
}

// CurrentPath returns the current source file path
// Alternative to Current for those who prefer explicit naming
//
// CurrentPath 获得当前源码文件路径
// Current 的替代版本，提供给喜欢明确命名的用户
func CurrentPath() string {
	return Skip(1)
}

// CurrentName returns the current source file name
// Gets just the filename without the DIR path
//
// CurrentName 获得当前源码文件名称
// 仅获取文件名，不包含 DIR 路径
func CurrentName() string {
	return filepath.Base(Skip(1))
}

// Name returns the runtime source file name
// Gets just the filename at execution location
//
// Name 获得运行时的源码文件名称
// 获取执行位置的文件名
func Name() string {
	return filepath.Base(Skip(1))
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

// GetPathChangeExtension changes the current source file extension
// Removes .go suffix and adds new extension like ".xxx.yyy.zzz"
// Common use: in config.go, get config.json path to read configuration
// Can add ".json", "_dev.json", "_uat.json" for different environments
// This function is essential for dynamic config file loading
//
// GetPathChangeExtension 把当前源码的文件路径去除结尾.go，再增加新的结尾
// 可以增加 ".xxx.yyy.zzz" 等任意扩展名
// 常见用途：在 config.go 里获取 config.json 的路径来读取配置
// 可以增加 ".json"、"_dev.json"、"_uat.json" 用于不同环境
// 这个函数对动态配置文件加载非常重要
func GetPathChangeExtension(pointExtension string) string {
	return GetSkipRemoveExtension(1) + pointExtension
}

// GetRex is a shorter alias for GetPathChangeExtension
// Changes current source file extension with new one
//
// GetRex 是 GetPathChangeExtension 的简短别名
// 更改当前源文件的扩展名
func GetRex(pointExtension string) string {
	return GetSkipRemoveExtension(1) + pointExtension
}

// GetNox returns current source file path without extension
// Removes the .go suffix from current file path
//
// GetNox 返回不带扩展名的当前源文件路径
// 从当前文件路径中移除 .go 后缀
func GetNox() string {
	return GetSkipRemoveExtension(1)
}

// GetPathRemoveExtension removes the .go extension from current source file path
// Less frequently used but kept for completeness
//
// GetPathRemoveExtension 把当前源码的文件路径去除结尾.go
// 使用频率较低但保留以保持完整性
func GetPathRemoveExtension() string {
	return GetSkipRemoveExtension(1)
}

func GetSkipRemoveExtension(skip int) string {
	_, path, _, ok := runtime.Caller(1 + skip)
	if !ok {
		panic(errors.New("wrong")) // Panic since this rarely fails and path retrieval is not used in production // 因为在99%的场景下都是不会出错的，而且跟获取代码路径相关的逻辑，通常也不会用在线上环境，因此直接 panic
	}
	const extension = ".go"
	if !strings.HasSuffix(strings.ToLower(path), extension) {
		panic(errors.Errorf("%s %s", path, extension))
	}
	return path[:len(path)-len(extension)]
}
