package runtestpath

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSrcPath(t *testing.T) {
	path := SrcPath(t)
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath.go"))
}

func TestSrcName(t *testing.T) {
	name := SrcName(t)
	t.Log(name)
	require.Equal(t, "runtestpath.go", name)
}

func TestSrcSkip(t *testing.T) {
	path := SrcSkip(t, 0)
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath.go"))
}

func TestSrcPathChangeExtension(t *testing.T) {
	path := SrcPathChangeExtension(t, ".json")
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath.json"))
}

func TestSrcRex(t *testing.T) {
	path := SrcRex(t, ".json")
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath.json"))
}

func TestSrcNox(t *testing.T) {
	path := SrcNox(t)
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath"))
}

func TestSrcPathRemoveExtension(t *testing.T) {
	path := SrcPathRemoveExtension(t)
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath"))
}

func TestSrcSkipRemoveExtension(t *testing.T) {
	path := SrcSkipRemoveExtension(t, 0)
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath"))
}

func TestAbsPath(t *testing.T) {
	//这个方法也是可以的，但是它并不总是有效的，有时候它返回的是项目的根目录，但是在多数情况下也是能拿来用的
	path, err := filepath.Abs(".")
	require.NoError(t, err)
	t.Log(path) //这里是对的，但不表示总是对的
}

func TestOsGetWD(t *testing.T) {
	//这个是获取当前工作目录，和源码文件所在目录不一定是同一个目录，这和你运行测试的位置有关
	path, err := os.Getwd()
	require.NoError(t, err)
	t.Log(path) //这里是工作区目录
}
