package runtestpath

// 在某些时候我们还需要根据当前运行的测试文件的路径，获取到相应的源代码文件的路径
// 假如你需要在测试文件中得到测试代码的文件路径，就直接用 runpath.Path() 就行
// 这个包就是专攻在运行测试时找被测试的源码文件
// 因此很明显的，该包的函数的第一个参数都是 t *testing.T 就是告诉使用者不应该在测试文件以外的文件里使用该包中的函数

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// SrcPath 假如你的测试文件是 abc_test.go 则在测试文件中调用该函数将返回 /aa/bb/cc/abc.go 即被测试文件的路径
func SrcPath(t *testing.T) string {
	return SrcSkip(t, 1)
}

// SrcName 假如你的测试文件是 abc_test.go 则在测试文件中调用该函数将返回 abc.go 即被测试文件的文件名
func SrcName(t *testing.T) string {
	return filepath.Base(SrcSkip(t, 1))
}

// SrcSkip 该函数在除以上两个场景以外暂时没有其他使用场景，但依然设置为 Export 的，或许在外面有别的用呢
func SrcSkip(t *testing.T, skip int) string {
	_, path, _, ok := runtime.Caller(1 + skip) //这里又调用了一层因此这里得补1次
	require.True(t, ok)
	require.True(t, strings.HasSuffix(path, "_test.go"))
	return path[:len(path)-len("_test.go")] + ".go"
}

// SrcPathChangeExtension 假如你的测试文件是 abc_test.go 则在测试文件中调用该函数且传递 ".json" 将返回 /aa/bb/cc/abc.json 即得到和被测试文件同名的配置文件名
func SrcPathChangeExtension(t *testing.T, pointExtension string) string {
	return SrcSkipRemoveExtension(t, 1) + pointExtension
}

func SrcRex(t *testing.T, pointExtension string) string {
	return SrcSkipRemoveExtension(t, 1) + pointExtension
}

func SrcNox(t *testing.T) string {
	return SrcSkipRemoveExtension(t, 1)
}

func SrcPathRemoveExtension(t *testing.T) string {
	return SrcSkipRemoveExtension(t, 1)
}

func SrcSkipRemoveExtension(t *testing.T, skip int) string {
	_, path, _, ok := runtime.Caller(1 + skip)
	require.True(t, ok)
	require.True(t, strings.HasSuffix(path, "_test.go"))
	return path[:len(path)-len("_test.go")]
}
