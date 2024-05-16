# runpath
获取正在执行的golang代码的位置信息，即 execution location，即源代码go文件在电脑里的绝对路径和行号。

使用 "runtime" 获得，因此包名起名为 "runpath" 即可，而不使用比较长的 executionlocation，但含义就是这样的，我还是喜欢短短的东西。

使用方法举例：
`
path := runpath.Path()
`
得到的就是当前这行代码所在文件的绝对路径啦

其次是读配置：
`
path := runpath.DIR.Join("config.json")
`
很明显它能帮你获取到配置文件路径，特别是在 testcase 测试用例中，不同的用例读不同的配置，也都是很正常的

当然我还贴心的准备了个根据测试文件找源码的操作：
`
func TestSrcPath(t *testing.T) {
    path := SrcPath(t)
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/utils_runtestpath.go"))
}
`
这个的使用场景，比如你的测试用例 testcase 运行一下就会 generate 生成源码，而你生成的源码恰好要写在对应的源文件里，要做源代码生成这个是必不可少的

在测试中读配置
`
func TestSrcPathChangeExtension(t *testing.T) {
    path := SrcPathChangeExtension(t, ".json")
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/utils_runtestpath.json"))
}
`
DDDD，毕竟很多时候我们就是需要在 testcase 里读取当前目录下的配置文件，因此直接把绝对路径算出来有利于使用

Give me stars! Thank you!
