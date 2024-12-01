package runpath

import (
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

// Path 获得运行时的源码文件路径
func Path() string {
	return Skip(1)
}

// Current 获得当前源码文件路径，因为该包名就是 runpath，则使用 Current 的含义就是 "current run path"，因此没必要起太长的函数名
func Current() string {
	return Skip(1)
}

// CurrentPath 获得当前源码文件路径，前面的 Current 或许已经够用，但也许别人不想用呢，毕竟我还是喜欢这个函数名
func CurrentPath() string {
	return Skip(1)
}

// CurrentName 获得当前源码文件名称
func CurrentName() string {
	return filepath.Base(Skip(1))
}

// Name 获得运行时的源码文件名称
func Name() string {
	return filepath.Base(Skip(1))
}

// Skip 获得运行时的源码位置
// 当传0的时候就是调用点的路径
// 这个文件的核心原则是，实参skip是调用位置的skip次数，而外部不需要关心里面的次数
// Skip(1) 相当于在相同位置调用runtime.Caller(1)的path
// Skip(2) 相当于在相同位置调用runtime.Caller(2)的path
func Skip(skip int) string {
	_, path, _, ok := runtime.Caller(1 + skip) //这里又调用了一层因此这里得补1次
	if !ok {
		panic(errors.New("wrong")) //因为在99%的场景下都是不会出错的，而且跟获取代码路径相关的逻辑，通常也不会用在线上环境，因此不要处理异常
	}
	return path
}

// GetPathChangeExtension 把当前源码的文件路径去除结尾.go，再增加新的结尾，比如增加 “.xxx.yyy.zzz” 都是可以的
// 这个函数通常用在，比如你的文件名叫做 config.go 你需要在 config.go 里获取到 config.json 的内容，则你可以通过此函数得到 config.json 的路径，接着你就可以调用你的读文件函数去读它
// 当然你不止可以在后面增加 ".json" 甚至可以在后面增加 "_dev.json" 和 "_uat.json" 以获得 "config_dev.json" 和 "config_uat.json" 的路径
// 因此我认为这个函数是非常重要的
func GetPathChangeExtension(pointExtension string) string {
	return GetSkipRemoveExtension(1) + pointExtension
}

func GetRex(pointExtension string) string {
	return GetSkipRemoveExtension(1) + pointExtension
}

func GetNox() string {
	return GetSkipRemoveExtension(1)
}

// GetPathRemoveExtension 把当前源码的文件路径去除结尾.go，这个函数好像使用的频率不会很高吧，当然也不是很确定呢，就这样留着吧
func GetPathRemoveExtension() string {
	return GetSkipRemoveExtension(1)
}

func GetSkipRemoveExtension(skip int) string {
	const suffixGo = ".go"
	_, path, _, ok := runtime.Caller(1 + skip)
	if !ok {
		panic(errors.New("wrong")) //因为在99%的场景下都是不会出错的，而且跟获取代码路径相关的逻辑，通常也不会用在线上环境，因此不要处理异常
	}
	if !strings.HasSuffix(strings.ToLower(path), suffixGo) {
		panic(errors.Errorf("%s %s", path, suffixGo))
	}
	return path[:len(path)-len(suffixGo)]
}
