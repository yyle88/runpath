package runpath

import (
	"path/filepath"
)

// 提供个类似于namespace的效果，这样能让函数名更简洁，使用起来意义也更明确，而且IDE代码提示的时候也能缩小选择范围
type parentNamespace struct{}

// 我觉得应该不会有使用者把这些全局变量设置为"nil"吧，啊哈哈哈，相信使用者不会给自己挖坑吧
var (
	PARENT = &parentNamespace{} //这里提供个全局变量以供外部使用，因为golang没有namespace的概念，但分包好像也没必要吧，就这样吧
	DIR    = &parentNamespace{} //这是同义词嘛，因为在我的理解中可能 parent 表示父路径，但实际 DIR 更简单些
)

func (T *parentNamespace) Path() string {
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
	return filepath.Join(append([]string{path}, names...)...)
}

func (T *parentNamespace) Up(skip int) string {
	path := filepath.Dir(Skip(1))
	for i := 0; i < skip; i++ {
		path = filepath.Dir(path)
	}
	return path
}

func (T *parentNamespace) UpTo(skip int, names ...string) string {
	path := filepath.Dir(Skip(1))
	for i := 0; i < skip; i++ {
		path = filepath.Dir(path)
	}
	return filepath.Join(append([]string{path}, names...)...)
}
