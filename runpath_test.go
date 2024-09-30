package runpath

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPath(t *testing.T) {
	path := Path()
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runpath_test.go"))
}

func TestCurrent(t *testing.T) {
	path := Current()
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runpath_test.go"))
}

func TestCurrentPath(t *testing.T) {
	path := CurrentPath()
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runpath_test.go"))
}

func TestCurrentName(t *testing.T) {
	name := CurrentName()
	t.Log(name)
	require.Equal(t, "runpath_test.go", name)
}

func TestName(t *testing.T) {
	name := Name()
	t.Log(name)
	require.Equal(t, "runpath_test.go", name)
}

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

func TestGetPathChangeExtension(t *testing.T) {
	path := GetPathChangeExtension(".json")
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runpath_test.json"))
}

func TestGetPathRemoveExtension(t *testing.T) {
	path := GetPathRemoveExtension()
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runpath_test"))
}

func TestGetSkipRemoveExtension(t *testing.T) {
	path := GetSkipRemoveExtension(0)
	t.Log(path)
	require.True(t, strings.HasSuffix(path, "runpath/runpath_test"))
}

func TestAbsPath(t *testing.T) {
	//这个方法也是可以的，但是它并不总是有效的，有时候它返回的是项目的根目录，但是在多数情况下也是能拿来用的
	path, err := filepath.Abs(".")
	require.NoError(t, err)
	t.Log(path) //这里是对的，但不表示总是对的
}
