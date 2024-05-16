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
	path := Path()
	name := Name()
	root := DIR.Path() // perhaps this variable should be named "dir", but "dir" is too ugly, more ugly than "ugly". so I prefer to use "root".
	path2 := filepath.Join(root, name)
	require.Equal(t, path, path2)
}

func Test_parentNamespace_Join2(t *testing.T) {
	path := Path()
	name := Name()
	path2 := DIR.Join(name)
	require.Equal(t, path, path2)
}
