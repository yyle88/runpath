package runtestpath

import (
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
