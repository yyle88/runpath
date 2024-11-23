# runpath

`runpath` 包提供了获取 Go 代码执行位置的功能，包括当前源文件的绝对路径和行号。

## README

[ENGLISH-DOC](README.md)

`runpath` 是一个用于获取 Go 程序执行位置的包，它利用 Go 的 `runtime` 包，帮助你轻松获得执行路径。相比 Go 内置的 `filepath.Abs()`，`runpath` 能够提供更加精确的路径，特别是在某些特殊场景里。

当然使用这个也是可行的，而且非常简便和标准：` filepath.Abs(".") ` 但它的结果也不总是符合预期的，有的时候不行。

因此我更倾向于使用 `runpath` 获取代码所在的路径。

### 安装

使用以下命令安装该包：

```shell
go get github.com/yyle88/runpath
```

### 示例用法

#### 获取当前源文件的绝对路径（执行位置）：

```go
path := runpath.Path()
```

此方法返回当前执行位置的源文件的绝对路径。

#### 获取当前源文件所在目录的绝对路径（执行目录）：

```go
path := runpath.PARENT.Path()
```

此方法返回当前执行位置的源文件所在目录的绝对路径。

#### 获取测试文件的路径：

```go
path := runtestpath.SrcPath(t)
```

此方法返回测试文件的源文件路径，通常用于定位当前正在运行的测试文件。

---

### 配置文件路径管理

你还可以利用 `runpath` 方便地构建配置文件的路径，特别是在测试中需要根据执行位置加载不同配置时非常有用。例如：

```go
path := runpath.DIR.Join("config.json")
```

该方法通过当前执行位置的路径动态构建配置文件路径。

你也可以使用 `PARENT` 来获取配置文件路径：

```go
path := runpath.PARENT.Join("config.json")
```

如果你需要向上导航多个目录层级，可以使用：

```go
path := runpath.PARENT.UpTo(3, "config.json")
```

这将返回类似 `filepath.Join(runpath.PARENT.Path(), "../../..", "config.json")` 的路径。

---

### 在测试中定位源代码路径

在测试用例中，可能需要动态引用源代码路径，可以使用以下方法：

```go
func TestSrcPath(t *testing.T) {
    path := runpath.SrcPath(t)
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/utils_runtestpath.go"))
}
```

这种方式特别适用于生成的代码需要与原始源文件一起放置的情况。

---

### 修改文件扩展名以适应不同的测试场景

你还可以根据测试场景动态更改文件扩展名，例如从 `.go` 转换为 `.json`：

```go
func TestSrcPathChangeExtension(t *testing.T) {
    path := runpath.SrcPathChangeExtension(t, ".json")
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/utils_runtestpath.json"))
}
```

这样你就可以根据不同的测试需求加载不同格式的文件。

---

## 函数概览

- `Path()`: 获取当前源文件的绝对路径，表示当前执行的位置。
- `Current()`, `CurrentPath()`, `CurrentName()`, `Name()`: `Path()` 的不同变体，用于获取当前文件的路径或文件名。
- `Skip(int)`: 从指定的调用帧获取路径（用于获取调用者的执行位置）。
- `GetPathChangeExtension()`: 返回当前源文件的路径，并可以更改文件扩展名（如将 `.go` 转为 `.json`）。
- `GetPathRemoveExtension()`: 返回当前源文件的路径，但去掉 `.go` 扩展名。
- `Join()`: 将当前目录与其他路径组件连接，动态构建路径。
- `Up()`, `UpTo()`: 向上导航指定层数的目录结构，并获取对应的路径。

---

### 测试文件特定操作：

- `SrcPath(t *testing.T)`: 获取测试文件的源路径。
- `SrcName(t *testing.T)`: 获取测试文件的文件名。
- `SrcPathChangeExtension(t *testing.T, ext string)`: 修改测试文件路径的扩展名（例如将 `.go` 更改为 `.json`）。
- `SrcSkipRemoveExtension(t *testing.T)`: 移除测试文件路径中的 `.go` 扩展名。

此包特别适用于测试文件中，帮助你根据测试文件所在位置引用源代码或配置文件路径。

---

## 许可证

本项目采用 MIT 许可证 - 请查看 [LICENSE](LICENSE) 文件了解详细信息。

---

## 贡献

欢迎大家贡献代码或改善该包！期待您的 stars 和 Pull Requests！

感谢使用 `runpath`！

Give me stars! Thank you!!!
