[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/runpath/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/runpath/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/runpath)](https://pkg.go.dev/github.com/yyle88/runpath)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/runpath/master.svg)](https://coveralls.io/github/yyle88/runpath?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/runpath.svg)](https://github.com/yyle88/runpath/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/runpath)](https://goreportcard.com/report/github.com/yyle88/runpath)

# runpath

`runpath` package provides func to get the execution location of Go code, including the absolute path of the source file.

## README

[中文说明](README.zh.md)

`runpath` is a package designed to retrieve the **execution location** of your Go code, including the exact path where your code is running. It uses Go’s `runtime` package to offer precise location tracking. Unlike the built-in `filepath.Abs()`, `runpath` is especially useful in certain scenarios where the exact execution path is needed.

we can also use the built-in approach: ` filepath.Abs(".") ` while, this doesn’t always provide the expected result in certain situations.

so I like to use `runpath` to get the path where the code is(execution location abs-path).

### Installation

Install the package with:

```shell
go get github.com/yyle88/runpath
```

### Example Usage

#### Get the absolute path of the current source file (execution location abs-path):

```go
path := runpath.Path()
```

This method returns the absolute path of the source file at the point of execution, which is the **execution location** of the current code.

#### Get the absolute path of the directory containing the current source file (execution directory abs-path):

```go
path := runpath.PARENT.Path()
```

This method returns the absolute path of the directory where the current source file resides, which can be useful for identifying the **execution directory abs-path**.

#### Get the source file path of the test being executed in testing:

```go
path := runtestpath.SrcPath(t)
```

This method returns the source path of the test file that is currently being run. It is particularly useful when testing and needing to know the exact location of the test file.

---

### Working with Configuration Files

You can use `runpath` to easily build paths to configuration files based on the **execution location** of the code. This is especially useful in tests where different configurations are loaded depending on where the test is being executed.

```go
path := runpath.DIR.Join("config.json")
```

This dynamically constructs the path to `config.json` relative to the **execution directory abs-path**.

You can also use `PARENT` for similar func:

```go
path := runpath.PARENT.Join("config.json")
```

If you need to navigate up multiple directory levels from the execution location, use:

```go
path := runpath.PARENT.UpTo(3, "config.json")
```

This will return the path to `config.json` located three levels up from the **execution directory abs-path**.

---

### Locating Source Code in Test Cases

When running tests, you may need to generate code or reference the source file path dynamically. Here's how you can locate the test file’s source path:

```go
func TestSrcPath(t *testing.T) {
    path := runpath.SrcPath(t)
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/utils_runtestpath.go"))
}
```

This approach helps when generating code that needs to be placed alongside the original source files based on the **execution location**.

---

### Changing File Extensions Based on Context

You can also change the file extension depending on the test context (e.g., from `.go` to `.json`):

```go
func TestSrcPathChangeExtension(t *testing.T) {
    path := runpath.SrcPathChangeExtension(t, ".json")
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/utils_runtestpath.json"))
}

func TestSrcRex(t *testing.T) {
    path := SrcRex(t, ".json")
    t.Log(path)
    require.True(t, strings.HasSuffix(path, "runpath/runtestpath/runtestpath.json"))
}
```

This allows you to load different types of files (e.g., configuration files) based on the **execution location** and test requirements.

---

## Function Overview

- `Path()`: Returns the absolute path of the source file where the code is executed, representing the **execution location**.
- `Current()`, `CurrentPath()`, `CurrentName()`, `Name()`: Variants of `Path()` that retrieve the file path or name based on the current execution context.
- `Skip(int)`: Retrieves the path from a specified call frame, useful for getting the execution location of the caller.
- `GetPathChangeExtension()`: Returns the current source file path with a new extension (e.g., changing `.go` to `.json`).
- `GetPathRemoveExtension()`: Returns the current source file path without the `.go` extension.
- `Join()`: Joins the current directory abs-path with additional path components, dynamically constructing paths based on the **execution location**.
- `Up()`, `UpTo()`: Navigates up the directory structure a specified number of levels from the **execution location**.

---

### Test-Specific Operations:

- `SrcPath(t *testing.T)`: Retrieves the source path of the file being tested.
- `SrcName(t *testing.T)`: Retrieves the name of the source file being tested.
- `SrcPathChangeExtension(t *testing.T, ext string)`: Changes the extension of the test file path (e.g., from `.go` to `.json`).
- `SrcSkipRemoveExtension(t *testing.T)`: Removes the `.go` extension from the test file path.

This package is particularly useful in test files, where you need to reference source code paths or configuration files based on the **execution location** of the test file.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Contributing

Feel free to contribute or improve the package! Stars and pull requests are always welcome!

Thank you for using `runpath`!

---

## Starring

[![starring](https://starchart.cc/yyle88/runpath.svg?variant=adaptive)](https://starchart.cc/yyle88/runpath)

Give me stars! Thank you!!!

---
