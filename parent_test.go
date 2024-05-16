package runpath

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parentNamespace_Path(t *testing.T) {
	t.Log(PARENT.Path())
}

func Test_parentNamespace_Root(t *testing.T) {
	t.Log(PARENT.Root())
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
