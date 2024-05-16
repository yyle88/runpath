package runpath

import (
	"path/filepath"
)

type parentNamespace struct{}

var PARENT = &parentNamespace{} //这里提供个全局变量以供外部使用，因为golang没有namespace的概念，但分包好像也没必要吧，就这样吧

func (T *parentNamespace) Path() string {
	return filepath.Dir(Skip(1))
}

func (T *parentNamespace) Root() string {
	return filepath.Dir(Skip(1))
}

func (T *parentNamespace) Name() string {
	return filepath.Base(filepath.Dir(Skip(1)))
}

func (T *parentNamespace) Skip(skip int) string {
	return filepath.Dir(Skip(1 + skip))
}

func (T *parentNamespace) Join(names ...string) string {
	path := filepath.Dir(Skip(1))
	subs := []string{path}
	subs = append(subs, names...)
	return filepath.Join(subs...)
}
