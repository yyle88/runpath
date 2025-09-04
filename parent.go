// Parent DIR namespace functionality for runtime path operations
// Provides namespace-like organization for parent DIR related functions
// Enables clean API design with PARENT.Path(), PARENT.Join() patterns
//
// 父 DIR 命名空间功能，用于运行时路径操作
// 为父 DIR 相关函数提供类似命名空间的组织方式
// 实现清晰的 API 设计，支持 PARENT.Path()、PARENT.Join() 模式
package runpath

import (
	"path/filepath"
)

// parentNamespace provides namespace-like organization for parent DIR operations
// Enables cleaner function names and more explicit usage patterns
// Improves IDE code completion by narrowing selection scope
//
// parentNamespace 为父 DIR 操作提供类似命名空间的组织方式
// 实现更清晰的函数名和更明确的使用模式
// 通过缩小选择范围来改善 IDE 代码完成
type parentNamespace struct{}

// Global namespace instances for parent DIR operations
// Users are trusted not to set these to nil (that would be self-defeating)
// Since Go lacks namespace concept, this provides organized API access
//
// 父 DIR 操作的全局命名空间实例
// 相信用户不会将其设置为 nil（那将是自找麻烦）
// 由于 Go 缺乏命名空间概念，这提供了有组织的 API 访问
var (
	PARENT = &parentNamespace{} // Global instance for parent DIR operations // 父 DIR 操作的全局实例
	DIR    = &parentNamespace{} // Synonym for PARENT, simpler name // PARENT 的同义词，更简单的名称
)

// Path returns the parent DIR path of the calling source file
// Gets the DIR containing the current execution location
//
// Path 返回调用源文件的父 DIR 路径
// 获取包含当前执行位置的 DIR
func (T *parentNamespace) Path() string {
	return filepath.Dir(Skip(1))
}

// Name returns the parent DIR name of the calling source file
// Gets just the name of the DIR containing current execution location
//
// Name 返回调用源文件的父 DIR 名称
// 仅获取包含当前执行位置的 DIR 名称
func (T *parentNamespace) Name() string {
	return filepath.Base(filepath.Dir(Skip(1)))
}

// Skip returns the parent DIR path with specified call frame skip
// Allows getting parent DIR from different call stack levels
//
// Skip 返回指定调用帧跳过的父 DIR 路径
// 允许从不同的调用堆栈级别获取父 DIR
func (T *parentNamespace) Skip(skip int) string {
	return filepath.Dir(Skip(1 + skip))
}

// Join constructs path by joining parent DIR with additional path components
// Dynamically builds paths relative to the calling file's parent DIR
//
// Join 通过将父 DIR 与额外的路径组件连接来构建路径
// 动态构建相对于调用文件的父 DIR 的路径
func (T *parentNamespace) Join(names ...string) string {
	path := filepath.Dir(Skip(1))
	return filepath.Join(append([]string{path}, names...)...)
}

// Up navigates up the DIR structure from parent DIR
// Goes up specified number of levels from the calling file's parent DIR
//
// Up 从父 DIR 向上导航 DIR 结构
// 从调用文件的父 DIR 向上跳过指定数量的级别
func (T *parentNamespace) Up(skip int) string {
	path := filepath.Dir(Skip(1))
	for i := 0; i < skip; i++ {
		path = filepath.Dir(path)
	}
	return path
}

// UpTo navigates up the DIR structure and joins with additional paths
// Combines Up() and Join() operations in a single call
//
// UpTo 向上导航 DIR 结构并与额外路径连接
// 在单次调用中组合 Up() 和 Join() 操作
func (T *parentNamespace) UpTo(skip int, names ...string) string {
	path := filepath.Dir(Skip(1))
	for i := 0; i < skip; i++ {
		path = filepath.Dir(path)
	}
	return filepath.Join(append([]string{path}, names...)...)
}
