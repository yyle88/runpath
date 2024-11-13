package runpath

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parentNamespace_Path(t *testing.T) {
	t.Log(PARENT.Path())
}

func Test_parentNamespace_Name(t *testing.T) {
	name := PARENT.Name()
	t.Log(name)
	require.Equal(t, "runpath", name)
}

func Test_parentNamespace_Skip(t *testing.T) {
	t.Log(PARENT.Skip(0))
}

func Test_parentNamespace_Join(t *testing.T) {
	t.Log(PARENT.Join("example.json"))
}

func Test_parentNamespace_Join1(t *testing.T) {
	name := Name()
	t.Log(name)
	root := DIR.Path() // perhaps this variable should be named "dir", but "dir" is too ugly, more ugly than "ugly". so I prefer to use "root" to clean code.
	t.Log(root)
	path := filepath.Join(root, name)
	t.Log(path)
	want := Path()
	t.Log(want)
	require.Equal(t, want, path)
}

func Test_parentNamespace_Join2(t *testing.T) {
	name := Name()
	t.Log(name)
	path := DIR.Join(name)
	t.Log(path)
	want := Path()
	t.Log(want)
	require.Equal(t, want, path)
}

func Test_parentNamespace_Up(t *testing.T) {
	for depth := 0; depth < 10; depth++ {
		t.Log(DIR.Up(depth))
	}
}

func Test_parentNamespace_UpTo(t *testing.T) {
	name := DIR.Name()
	t.Log(name)
	path := DIR.UpTo(1, name)
	t.Log(path)
	want := DIR.Path()
	t.Log(want)
	require.Equal(t, want, path)
}
